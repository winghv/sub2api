package service

import (
	"crypto/pbkdf2"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"strconv"
	"strings"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

// encodeSuperWHVPBKDF2 生成与 SuperWHV Cloud 相同格式的 PBKDF2-SHA256 哈希，
// 仅用于测试（线上格式由 SuperWHV 的 internal/controlplane/auth/password.go 产出）。
func encodeSuperWHVPBKDF2(t *testing.T, password string, iter int) string {
	t.Helper()
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		t.Fatalf("rand: %v", err)
	}
	key, err := pbkdf2.Key(sha256.New, password, salt, iter, 32)
	if err != nil {
		t.Fatalf("pbkdf2: %v", err)
	}
	return strings.Join([]string{
		"pbkdf2-sha256",
		strconv.Itoa(iter),
		base64.RawURLEncoding.EncodeToString(salt),
		base64.RawURLEncoding.EncodeToString(key),
	}, "$")
}

func TestVerifyPasswordHashSupportsSuperWHVPBKDF2(t *testing.T) {
	hash := encodeSuperWHVPBKDF2(t, "s3cret-pass", 210000)
	if !VerifyPasswordHash(hash, "s3cret-pass") {
		t.Fatal("correct password rejected")
	}
	if VerifyPasswordHash(hash, "wrong") {
		t.Fatal("wrong password accepted")
	}
	if !IsSupportedPasswordHash(hash) {
		t.Fatal("pbkdf2 hash not recognized as supported")
	}
}

func TestVerifyPasswordHashStillSupportsBcrypt(t *testing.T) {
	hash, err := bcrypt.GenerateFromPassword([]byte("plain-pass"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("bcrypt: %v", err)
	}
	if !VerifyPasswordHash(string(hash), "plain-pass") {
		t.Fatal("correct password rejected")
	}
	if VerifyPasswordHash(string(hash), "wrong") {
		t.Fatal("wrong password accepted")
	}
}

func TestIsSupportedPasswordHashRejectsGarbage(t *testing.T) {
	for _, bad := range []string{"", "plaintext", "md5$1$abc", "pbkdf2-sha256$abc$xx$yy"} {
		if IsSupportedPasswordHash(bad) {
			t.Fatalf("accepted garbage %q", bad)
		}
	}
}

func TestSetUserPasswordPrefersPrehashed(t *testing.T) {
	user := &User{}
	pb := encodeSuperWHVPBKDF2(t, "x", 210000)
	if err := setUserPassword(user, "ignored-plaintext", pb); err != nil {
		t.Fatalf("setUserPassword: %v", err)
	}
	if user.PasswordHash != pb {
		t.Fatal("pre-hashed value not stored verbatim")
	}
	if err := setUserPassword(user, "whatever", "not-a-hash"); err == nil {
		t.Fatal("unsupported hash accepted")
	}
	bc, err := bcrypt.GenerateFromPassword([]byte("plain"), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("bcrypt: %v", err)
	}
	user2 := &User{}
	if err := setUserPassword(user2, "plain", ""); err != nil {
		t.Fatalf("plaintext path: %v", err)
	}
	if user2.PasswordHash == "" || !VerifyPasswordHash(user2.PasswordHash, "plain") {
		t.Fatal("plaintext path did not produce verifiable hash")
	}
	_ = bc
}
