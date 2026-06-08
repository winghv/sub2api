# Superwhv 使用教程

> 支持 Claude Code、Codex、Gemini CLI 国内直连与配置，官方原版体验

---

## 目录

- [1. 重要说明](#1-重要说明)
- [2. 常遇问题](#3-常遇问题)
- [3. 安装教程](#4-安装教程)

---

## 1. 重要说明

官网兑换地址 ：
- Claude Code / Gemini CLI 都有一定上手门槛
- 本教程只负责让你在官方客户端能用上，第三方包装客户端需要自行研究

  1. 进入网站创建账号并完成基础信息填写。
  2. 使用兑换码完成核销（可在左侧“兑换”菜单操作）。
  3. 创建秘钥（名称可自定义，分组选择每天90或者每天180刀分组，其他分组不能用）。
  4. 密钥使用，可以根据页面引导手动改配置，也可以通过CCS直接一键导入
  5. 使用记录查询（左侧“使用记录”查看明细，右上角可看每日消耗）。

---

## 2. 安装教程

### 选择你的客户端

- [Claude Code](#claude-code)
- [Codex](#codex)
- [Gemini](#gemini)
- [VSCode & Cursor](#vscode--cursor)

### 选择你的系统

- [Windows](#windows)
- [macOS](#macos)
- [Linux](#linux)

---

## Claude Code

### Windows

#### 前置要求：安装 Git

Claude Code 需要 Git 才能运行。如果启动时遇到以下错误：

```
Claude Code on Windows requires git-bash (https://git-scm.com/downloads/win).
If installed but not in PATH, set environment variable pointing to your bash.exe,
similar to: CLAUDE_CODE_GIT_BASH_PATH=C:\Program Files\Git\bin\bash.exe"
```

请先安装 Git：

- [官方下载](https://git-scm.com/downloads/win)
- [国内加速下载](https://static.yoouu.cn/Git-2.52.0-64-bit.exe)

> ⚠️ 必须用 PowerShell 或 Windows Terminal！CMD 是上古遗物，用它必出问题！

**如果已安装但不在 PATH 中，设置环境变量（路径改为你实际的 Git 安装路径）：**

```powershell
# PowerShell
CLAUDE_CODE_GIT_BASH_PATH=C:\Program Files\Git\bin\bash.exe
```

**验证安装：**

```powershell
git --version
```

#### 安装方式

##### 方式一：原生安装（推荐）

无需 Node.js，自动后台更新，官方推荐。

```powershell
# PowerShell (管理员)
irm https://claude.ai/install.ps1 | iex
```

##### 方式二：WinGet 安装

```powershell
# 安装
winget install Anthropic.ClaudeCode

# 更新到最新版本
winget upgrade Anthropic.ClaudeCode
```

##### 方式三（推荐）：npm 安装（需要 Node.js 18+）

如果网络不好，可以使用国内镜像：

```powershell
# 直接安装
npm install -g @anthropic-ai/claude-code

# 使用国内镜像
npm install -g @anthropic-ai/claude-code --registry=https://registry.npmmirror.com
```

#### 验证安装

```powershell
claude --version
```

显示版本号说明安装成功！

#### 设置环境变量

配置连接到中转服务。

**永久设置（系统级别，需管理员权限）：**

```powershell
# PowerShell (管理员)
[System.Environment]::SetEnvironmentVariable("ANTHROPIC_BASE_URL", "https://api.superwhv.me", [System.EnvironmentVariableTarget]::Machine)
[System.Environment]::SetEnvironmentVariable("ANTHROPIC_AUTH_TOKEN", "sk-你的APIKey", [System.EnvironmentVariableTarget]::Machine)
[System.Environment]::SetEnvironmentVariable("API_TIMEOUT_MS", "3000000", [System.EnvironmentVariableTarget]::Machine)
[System.Environment]::SetEnvironmentVariable("CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC", "1", [System.EnvironmentVariableTarget]::Machine)
```

**永久设置（用户级别）：**

```powershell
# PowerShell
[System.Environment]::SetEnvironmentVariable("ANTHROPIC_BASE_URL", "https://api.superwhv.me", [System.EnvironmentVariableTarget]::User)
[System.Environment]::SetEnvironmentVariable("ANTHROPIC_AUTH_TOKEN", "sk-你的APIKey", [System.EnvironmentVariableTarget]::User)
[System.Environment]::SetEnvironmentVariable("API_TIMEOUT_MS", "3000000", [System.EnvironmentVariableTarget]::User)
[System.Environment]::SetEnvironmentVariable("CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC", "1", [System.EnvironmentVariableTarget]::User)
```

> 💡 `sk-你的APIKey` 请替换为你在后台申请的 API Key！

**临时设置（仅当前窗口有效）：**

```powershell
$env:ANTHROPIC_BASE_URL = "https://api.superwhv.me"
$env:ANTHROPIC_AUTH_TOKEN = "sk-你的APIKey"
$env:API_TIMEOUT_MS = "3000000"
$env:CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC = "1"
```

#### 验证环境变量

```powershell
Get-ChildItem Env:ANTHROPIC_*
```

#### 不再使用？删除环境变量

```powershell
# PowerShell (管理员)
[System.Environment]::SetEnvironmentVariable("ANTHROPIC_BASE_URL", $null, [System.EnvironmentVariableTarget]::Machine)
[System.Environment]::SetEnvironmentVariable("ANTHROPIC_AUTH_TOKEN", $null, [System.EnvironmentVariableTarget]::Machine)
[System.Environment]::SetEnvironmentVariable("API_TIMEOUT_MS", $null, [System.EnvironmentVariableTarget]::Machine)
[System.Environment]::SetEnvironmentVariable("CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC", $null, [System.EnvironmentVariableTarget]::Machine)
```

#### 开始使用

```powershell
# 进入项目目录
cd C:\path\to\your\project

# 启动 Claude Code
claude
```

---

### macOS

#### 前置要求：安装 Git

```bash
# 使用 Homebrew 安装
brew install git

# 验证安装
git --version
```

#### 安装方式

##### 方式一：原生安装（推荐）

```bash
curl -fsSL https://claude.ai/install.sh | sh
```

##### 方式二：Homebrew 安装

```bash
brew install claude
```

##### 方式三：npm 安装

```bash
npm install -g @anthropic-ai/claude-code
```

#### 设置环境变量

```bash
# 永久设置（添加到 ~/.zshrc 或 ~/.bash_profile）
echo 'export ANTHROPIC_BASE_URL="https://api.superwhv.me"' >> ~/.zshrc
echo 'export ANTHROPIC_AUTH_TOKEN="sk-你的APIKey"' >> ~/.zshrc
echo 'export API_TIMEOUT_MS="3000000"' >> ~/.zshrc
echo 'export CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC="1"' >> ~/.zshrc

# 使配置生效
source ~/.zshrc
```

> 💡 `sk-你的APIKey` 请替换为你在后台申请的 API Key！

#### 开始使用

```bash
# 进入项目目录
cd /path/to/your/project

# 启动 Claude Code
claude
```

---

### Linux

#### 前置要求：安装 Git

```bash
sudo apt update && sudo apt install git
git --version
```

#### 安装方式

##### 方式一：原生安装（推荐）

```bash
curl -fsSL https://claude.ai/install.sh | sh
```

##### 方式二：npm 安装

```bash
sudo npm install -g @anthropic-ai/claude-code
```

#### 设置环境变量

```bash
# 永久设置（添加到 ~/.bashrc 或 ~/.zshrc）
echo 'export ANTHROPIC_BASE_URL="https://api.superwhv.me"' >> ~/.bashrc
echo 'export ANTHROPIC_AUTH_TOKEN="sk-你的APIKey"' >> ~/.bashrc
echo 'export API_TIMEOUT_MS="3000000"' >> ~/.bashrc
echo 'export CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC="1"' >> ~/.bashrc

# 使配置生效
source ~/.bashrc
```

> 💡 `sk-你的APIKey` 请替换为你在后台申请的 API Key！

#### 开始使用

```bash
# 进入项目目录
cd /path/to/your/project

# 启动 Claude Code
claude
```

---

## Codex

### Windows

#### 前置要求：安装 Node.js

检查是否已安装：

```powershell
node --version
npm --version
```

如果命令不存在，访问 [Node.js 官网](https://nodejs.org/) 下载 Windows 对应的 LTS 安装包。

#### 安装 codex

```powershell
# PowerShell（管理员）
npm i -g @openai/codex --registry=https://registry.npmmirror.com
codex --version
```

#### 配置 codex

打开文件资源管理器，找到 `C:\Users\你的用户名\.codex` 文件夹（不存在则创建）。

**创建 `config.toml` 文件：**

```toml
model_provider = "sub2api"
model = "gpt-5.3-codex"
model_reasoning_effort = "high"
network_access = "enabled"
disable_response_storage = true
windows_wsl_setup_acknowledged = true
model_verbosity = "high"

[model_providers.sub2api]
name = "sub2api"
base_url = "https://api.superwhv.me"
wire_api = "responses"
requires_openai_auth = true
```

**创建 `auth.json` 文件：**

```json
{
  "OPENAI_API_KEY": "sk-你的APIKey"
}
```

> ⚠️ 将 `sk-你的APIKey` 替换为你的 API Key！

#### 启动 codex

```powershell
cd C:\path\to\your\project
codex
```

首次启动时，codex 会进行初始化配置。

#### 编辑器插件使用

在插件市场搜索并安装 **Codex - OpenAI's coding agent**。

VS Code、Cursor、Trae、Kiro 操作基本一样，只要前面的 Codex 配置已经完成，插件打开后就能直接使用。

> 第一次使用时，记得在输入框右上角把 **Sandbox（沙盒）** 打开。

#### Codex 常见报错

| 错误码 | 可能原因 | 解决方案 |
|--------|----------|----------|
| `401` | 分组选错（如选成了 Codex 分组） | 检查后台分组设置，选择正确的每天额度分组 |
| `403` | 把兑换码当成密钥配置 | 应填写后台创建出来的 API 密钥 |
| `429` | 当前额度已用完 | 去后台查看剩余额度，额度恢复后再试 |

---

### macOS

#### 前置要求：安装 Node.js

```bash
# 使用 Homebrew 安装
brew install node

# 验证安装
node --version
npm --version
```

#### 安装 codex

```bash
npm i -g @openai/codex --registry=https://registry.npmmirror.com
codex --version
```

#### 配置 codex

```bash
mkdir -p ~/.codex
```

**创建 `~/.codex/config.toml` 文件：**

```toml
model_provider = "sub2api"
model = "gpt-5.3-codex"
model_reasoning_effort = "high"
network_access = "enabled"
disable_response_storage = true
model_verbosity = "high"

[model_providers.sub2api]
name = "sub2api"
base_url = "https://api.superwhv.me"
wire_api = "responses"
requires_openai_auth = true
```

**创建 `~/.codex/auth.json` 文件：**

```json
{
  "OPENAI_API_KEY": "sk-你的APIKey"
}
```

#### 启动 codex

```bash
cd /path/to/your/project
codex
```

---

### Linux

#### 前置要求：安装 Node.js

```bash
sudo apt update && sudo apt install nodejs npm
node --version
npm --version
```

#### 安装 codex

```bash
sudo npm i -g @openai/codex --registry=https://registry.npmmirror.com
codex --version
```

#### 配置 codex

```bash
mkdir -p ~/.codex
```

**创建 `~/.codex/config.toml` 文件：**

```toml
model_provider = "sub2api"
model = "gpt-5.3-codex"
model_reasoning_effort = "high"
network_access = "enabled"
disable_response_storage = true
model_verbosity = "high"

[model_providers.sub2api]
name = "sub2api"
base_url = "https://api.superwhv.me"
wire_api = "responses"
requires_openai_auth = true
```

**创建 `~/.codex/auth.json` 文件：**

```json
{
  "OPENAI_API_KEY": "sk-你的APIKey"
}
```

#### 启动 codex

```bash
cd /path/to/your/project
codex
```

---

## Gemini

### Windows

#### 安装 Gemini CLI

```powershell
# PowerShell（管理员）
npm i -g @google/gemini-cli --registry=https://registry.npmmirror.com
```

#### 配置环境变量

```powershell
# PowerShell (管理员)
[System.Environment]::SetEnvironmentVariable("GEMINI_BASE_URL", "https://api.superwhv.me/v1beta", [System.EnvironmentVariableTarget]::Machine)
[System.Environment]::SetEnvironmentVariable("GEMINI_API_KEY", "sk-你的APIKey", [System.EnvironmentVariableTarget]::Machine)
```

#### 开始使用

```powershell
gemini
```

---

### macOS

#### 安装 Gemini CLI

```bash
npm i -g @google/gemini-cli --registry=https://registry.npmmirror.com
```

#### 配置环境变量

```bash
echo 'export GEMINI_BASE_URL="https://api.superwhv.me/v1beta"' >> ~/.zshrc
echo 'export GEMINI_API_KEY="sk-你的APIKey"' >> ~/.zshrc
source ~/.zshrc
```

#### 开始使用

```bash
gemini
```

---

### Linux

#### 安装 Gemini CLI

```bash
sudo npm i -g @google/gemini-cli --registry=https://registry.npmmirror.com
```

#### 配置环境变量

```bash
echo 'export GEMINI_BASE_URL="https://api.superwhv.me/v1beta"' >> ~/.bashrc
echo 'export GEMINI_API_KEY="sk-你的APIKey"' >> ~/.bashrc
source ~/.bashrc
```

#### 开始使用

```bash
gemini
```

---

## VSCode & Cursor

### 安装插件

在 VSCode 或 Cursor 的插件市场搜索并安装：

- **Claude Code**: 搜索 `claude` 或 `Anthropic`
- **Codex**: 搜索 `Codex - OpenAI's coding agent`

### 配置

插件会自动使用系统环境变量中配置的 `ANTHROPIC_BASE_URL` 和 `ANTHROPIC_AUTH_TOKEN`。

确保已正确设置环境变量后，重启编辑器即可使用。

---

## 错误码参考

| 错误码 | 说明 |
|--------|------|
| `401` | API Key 无效或已过期 |
| `403` | 余额不足或权限不足 |
| `429` | 请求过于频繁，触发限流 |
| `500` | 服务器内部错误 |
| `502` | 上游服务异常 |
| `503` | 服务暂时不可用 |

---

## 技术支持

如遇问题，请联系管理员或查看后台使用记录。

---
