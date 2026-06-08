# Sub2API 部署记录

> **敏感信息已脱敏** - 本文档可安全提交到 Git
> 真实凭证请查看 `.env` 文件（已加入 .gitignore）

---

## 部署信息

| 项目 | 值 |
|------|-----|
| 部署日期 | 2026-04-08 |
| 服务器 | RackNerd VPS (Dallas, TX) |
| 服务器 IP | 192.227.249.159 |
| 域名 | api.superwhv.me |
| 版本 | v0.1.109 |

---

## 架构

```
用户请求 https://api.superwhv.me
         ↓
    Cloudflare Edge (全球节点)
         ↓
    Cloudflare Tunnel (加密 QUIC/UDP)
         ↓
    cloudflared (Docker容器内)
         ↓
    Sub2API :8080 (Docker容器)
         ↓
    PostgreSQL :5432
    Redis :6379
```

---

## 服务状态

```bash
# VPS 上查看服务状态
ssh root@<SERVER_IP> "cd /opt/sub2api/deploy && docker compose ps"
```

| 服务 | 状态 |
|------|------|
| sub2api | healthy |
| postgres | healthy |
| redis | healthy |
| cloudflared | running |

---

## 访问地址

| 地址 | 说明 |
|------|------|
| https://api.superwhv.me | 主站 (登录/注册/Dashboard) |
| https://api.superwhv.me/setup | 设置向导 (首次访问) |

---

## 默认账号

| 账号 | 密码 |
|------|------|
| admin@superwhv.me | **[已修改]** |

---

## Docker Compose 配置

使用的配置文件: `docker-compose.yml`

主要服务:
- `sub2api` - 主应用 (端口 8080)
- `postgres` - PostgreSQL 16 (端口 5432)
- `redis` - Redis 7 (端口 6379)
- `cloudflared` - Cloudflare Tunnel

---

## 常用命令

```bash
# 登录 VPS
ssh root@<SERVER_IP>

# 查看服务状态
cd /opt/sub2api/deploy && docker compose ps

# 查看日志
docker compose logs -f sub2api
docker compose logs -f cloudflared

# 重启服务
docker compose restart

# 完全重建 (慎用! 删除所有数据)
docker compose down -v && docker compose up -d

# 更新镜像
docker compose pull && docker compose up -d
```

---

## .env 配置说明

`.env` 文件包含以下敏感信息 (**不提交到 Git**):

| 变量 | 说明 |
|------|------|
| `POSTGRES_PASSWORD` | PostgreSQL 数据库密码 |
| `JWT_SECRET` | JWT 签名密钥 |
| `TOTP_ENCRYPTION_KEY` | TOTP 加密密钥 |
| `CLOUDFLARE_TUNNEL_TOKEN` | Cloudflare Tunnel 连接令牌 |

生成强密码:
```bash
openssl rand -base64 32
```

---

## 修改密码

### 1. 修改 PostgreSQL 密码

```bash
# 1. 编辑 .env 文件
vim /opt/sub2api/deploy/.env
# 修改 POSTGRES_PASSWORD=your_new_password

# 2. 重启服务 (会重置数据库)
cd /opt/sub2api/deploy && docker compose down -v && docker compose up -d

# 3. 重新完成设置向导
# 访问 https://api.superwhv.me/setup
```

### 2. 修改 Admin 密码

通过 Web 界面:
1. 登录 https://api.superwhv.me
2. 进入 **Profile** → 修改密码

---

## 迁移到新服务器

### 数据导出

```bash
# 在原服务器执行
docker exec sub2api-postgres-1 pg_dump -U sub2api sub2api > postgres_backup.sql
docker cp sub2api-postgres-1:/var/lib/postgresql/data/postgres_backup.sql ./
docker compose down
```

### 数据导入

```bash
# 在新服务器执行
# 1. 安装 Docker + Docker Compose
# 2. 复制 .env 和 docker-compose.yml
# 3. 启动基础服务
docker compose up -d postgres redis

# 4. 导入数据
docker exec -i sub2api-postgres-1 psql -U sub2api sub2api < postgres_backup.sql

# 5. 启动全部服务
docker compose up -d
```

### 更新 Cloudflare Tunnel

1. 在 Cloudflare Dashboard 删除旧 Tunnel
2. 创建新 Tunnel，获取新 Token
3. 更新 `.env` 中的 `CLOUDFLARE_TUNNEL_TOKEN`
4. 重启服务

---

## 故障排查

### 服务不健康

```bash
# 查看详细日志
docker compose logs sub2api

# 检查容器网络
docker network ls
docker network inspect deploy_sub2api-net
```

### 数据库连接失败

```
错误: pq: password authentication failed
解决: 重置 .env 中的 POSTGRES_PASSWORD，然后重建数据库
```

### Tunnel 连接失败

```bash
# 检查 Token 是否正确
docker compose logs cloudflared | grep -i token

# 重建 Tunnel
docker compose down cloudflared
docker compose up -d cloudflared
```

---

## 安全建议

1. **定期更新镜像**
   ```bash
   docker compose pull && docker compose up -d
   ```

2. **启用 Cloudflare WAF**
   - Dashboard → Security → WAF

3. **设置 Rate Limiting**
   - Dashboard → Security → Tools → Rate Limiting
   - 建议: 100 requests/minute per IP

4. **不要提交 .env 到 Git**
   - 已加入 `.gitignore`

---

## 备份策略

```bash
# 创建备份脚本
#!/bin/bash
DATE=$(date +%Y%m%d_%H%M%S)
BACKUP_DIR=/opt/sub2api/backups
mkdir -p $BACKUP_DIR

# 备份数据库
docker exec sub2api-postgres-1 pg_dump -U sub2api sub2api > $BACKUP_DIR/postgres_$DATE.sql

# 保留最近 7 天
find $BACKUP_DIR -mtime +7 -delete
```

添加到 crontab:
```bash
0 3 * * * /opt/sub2api/backup.sh
```
