# Sub2API + Cloudflare Tunnel 部署指南

## 概述

本方案让你在本地 Docker 运行 Sub2API，通过 Cloudflare Tunnel 实现全球访问，**无需公网 IP、无需开放端口**。

```
用户请求 https://api.yourdomain.com
         ↓
    Cloudflare Edge (全球节点)
         ↓
    Cloudflare Tunnel (加密 WebSocket)
         ↓
    cloudflared (你的本地 Docker)
         ↓
    Sub2API 服务
```

---

## 完整部署步骤

### 第一步：准备域名 (如有，跳过)

1. 购买一个域名 (如 `example.com`)
2. 将域名 NS 改为 Cloudflare 的 NS:
   ```
   ns1.cloudflare.com
   ns2.cloudflare.com
   ```
3. 等待生效 (可能需要几分钟到 48 小时)

### 第二步：创建 Cloudflare Tunnel

1. 登录 https://dash.cloudflare.com/
2. 进入 **Zero Trust** → **Networks** → **Tunnels**
3. 点击 **Create a tunnel**
4. 选择 **Cloudflared** 作为 connector
5. 隧道名称填: `sub2api`
6. 点击 **Save tunnel**
7. 复制 **Token** (格式: `eyJhIjoi...`)，后面要用

### 第三步：配置 DNS (自动)

创建隧道后，Cloudflare 会提示你添加 DNS 记录。

在 Cloudflare DNS 设置中添加:
```
Type: CNAME
Name: api (或你想用的子域名)
Target: <tunnel-id>.cfdtunnel.com
Proxy status: DNS only (或 Proxied)
```

### 第四步：下载并配置

```bash
# 进入部署目录
cd ~/sub2api

# 下载 docker-compose 和配置文件
# (如果你是复制粘贴，手动创建以下文件)

# 复制环境变量文件
cp .env.cloudflare .env

# 编辑 .env，填入你的值:
# - CLOUDFLARE_TUNNEL_TOKEN (必填)
# - POSTGRES_PASSWORD (必填)
vim .env
```

### 第五步：启动服务

```bash
# 启动所有服务
docker compose -f docker-compose.cloudflare.yml up -d

# 查看状态
docker compose -f docker-compose.cloudflare.yml ps

# 查看日志
docker compose -f docker-compose.cloudflare.yml logs -f
```

### 第六步：验证

```bash
# 检查服务健康
curl https://api.yourdomain.com/health

# 或本地检查
curl http://localhost:8080/health

# 查看 cloudflared 状态
docker compose -f docker-compose.cloudflare.yml logs cloudflared
```

---

## 目录结构

```
~/sub2api/
├── docker-compose.cloudflare.yml   # Docker 编排配置
├── config.cloudflare.yaml          # Sub2API 配置
├── .env                            # 环境变量 (包含密钥)
└── .env.cloudflare                 # 环境变量模板
```

---

## 常用命令

```bash
# 启动
docker compose -f docker-compose.cloudflare.yml up -d

# 停止
docker compose -f docker-compose.cloudflare.yml down

# 重启
docker compose -f docker-compose.cloudflare.yml restart

# 查看日志
docker compose -f docker-compose.cloudflare.yml logs -f [service]

# 只看 sub2api 日志
docker compose -f docker-compose.cloudflare.yml logs -f sub2api

# 只看 cloudflared 日志
docker compose -f docker-compose.cloudflare.yml logs -f cloudflared

# 重建 (代码更新后)
docker compose -f docker-compose.cloudflare.yml up -d --force-recreate

# 完全清理 (慎用! 会删除数据)
docker compose -f docker-compose.cloudflare.yml down -v
```

---

## 暴露多个服务

如果需要同时暴露管理后台或 Prometheus:

1. 在 Cloudflare Dashboard 添加更多 DNS 记录:
   ```
   admin.yourdomain.com → <tunnel-id>.cfdtunnel.com
   metrics.yourdomain.com → <tunnel-id>.cfdtunnel.com
   ```

2. 修改 docker-compose.cloudflare.yml，添加更多 expose 端口:
   ```yaml
   sub2api:
     expose:
       - "8080"      # API
       - "3030"      # 管理后台
       - "9090"      # Prometheus
   ```

3. 重启服务:
   ```bash
   docker compose -f docker-compose.cloudflare.yml up -d
   ```

---

## 文件描述符优化

Linux 系统可能需要调整:

```bash
# 临时生效
ulimit -n 100000

# 永久生效 (编辑 /etc/security/limits.conf)
* soft nofile 100000
* hard nofile 100000
```

---

## 常见问题

### Q: Tunnel 连不上

```bash
# 1. 检查 Token 是否正确
docker compose -f docker-compose.cloudflare.yml logs cloudflared | grep -i token

# 2. 检查网络
docker compose -f docker-compose.cloudflare.yml exec cloudflared curl -v https://www.cloudflare.com/

# 3. 重建 Tunnel
docker compose -f docker-compose.cloudflare.yml down
docker compose -f docker-compose.cloudflare.yml up -d
```

### Q: 访问很慢

1. 检查 Cloudflare 节点位置: https://speed.cloudflare.com/
2. 启用 Argo Smart Routing (Dashboard → Networks → Argo)
3. 检查本地网络

### Q: 如何迁移服务器?

1. 在新服务器安装 Docker
2. 复制整个 `~/sub2api` 目录
3. 修改 `.env` 中的密码
4. 在 Cloudflare Dashboard 删除旧隧道
5. 创建新隧道，获取新 Token
6. 启动服务

### Q: 支持 WebSocket 吗?

**支持!** Cloudflare Tunnel 原生支持 WebSocket，Sub2API 的 WS 功能完全可用。

---

## 安全建议

1. **不要把 .env 文件提交到 Git**
   ```bash
   echo ".env" >> .gitignore
   ```

2. **定期更新镜像**
   ```bash
   docker compose -f docker-compose.cloudflare.yml pull
   docker compose -f docker-compose.cloudflare.yml up -d
   ```

3. **启用 Cloudflare WAF**
   - Dashboard → Security → WAF
   - 建议开启 "Cloudflare Specials" 规则

4. **设置访问频率限制**
   - Dashboard → Security → Tools
   - 建议: 100 requests/minute per IP
