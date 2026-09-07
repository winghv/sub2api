package service

import (
	"crypto/pbkdf2"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// 跨系统用户同步的密码哈希兼容层。
//
// SuperWHV Cloud 使用 PBKDF2-SHA256（格式 `pbkdf2-sha256$<iter>$<salt>$<key>`，
// salt/key 为 RawURL base64），sub2api 使用 bcrypt。用户同步会把来源系统生成
// 的哈希原样存入本系统，因此验证必须同时支持两种格式；明文密码永远不跨系统
// 传输，只有不可逆哈希。

const superwhvPasswordHashPrefix = "pbkdf2-sha256$"

// ErrUnsupportedPasswordHash 表示密码哈希不是可识别的格式。
var ErrUnsupportedPasswordHash = errors.New("unsupported password hash format")

// IsSupportedPasswordHash 报告 hash 是否为可验证的 bcrypt 或 PBKDF2-SHA256 哈希。
func IsSupportedPasswordHash(hash string) bool {
	if strings.HasPrefix(hash, superwhvPasswordHashPrefix) {
		parts := strings.Split(hash, "$")
		if len(parts) != 4 {
			return false
		}
		iter, err := strconv.Atoi(parts[1])
		return err == nil && iter > 0
	}
	return len(hash) == 60 && strings.HasPrefix(hash, "$2")
}

// VerifyPasswordHash 按哈希格式分流验证：PBKDF2-SHA256 走本实现，其余按 bcrypt。
func VerifyPasswordHash(hashedPassword, password string) bool {
	if strings.HasPrefix(hashedPassword, superwhvPasswordHashPrefix) {
		return verifyPBKDF2SHA256(hashedPassword, password)
	}
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

func verifyPBKDF2SHA256(encoded, password string) bool {
	parts := strings.Split(encoded, "$")
	if len(parts) != 4 {
		return false
	}
	iter, err := strconv.Atoi(parts[1])
	if err != nil || iter <= 0 {
		return false
	}
	salt, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil || len(salt) == 0 {
		return false
	}
	want, err := base64.RawURLEncoding.DecodeString(parts[3])
	if err != nil || len(want) == 0 {
		return false
	}
	got, err := pbkdf2.Key(sha256.New, password, salt, iter, len(want))
	if err != nil {
		return false
	}
	return subtle.ConstantTimeCompare(got, want) == 1
}

// setUserPassword 按"预哈希优先"设置密码：同步传入的哈希直存（需为受支持格式），
// 否则按明文走 bcrypt。明文密码不会落库或出现在日志。
func setUserPassword(user *User, password, passwordHash string) error {
	if passwordHash != "" {
		if !IsSupportedPasswordHash(passwordHash) {
			return ErrUnsupportedPasswordHash
		}
		user.PasswordHash = passwordHash
		return nil
	}
	return user.SetPassword(password)
}
