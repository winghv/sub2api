<template>
  <AppLayout>
    <div class="min-h-screen min-w-0 bg-gray-50 dark:bg-dark-950">
      <!-- Header -->
      <div class="border-b border-gray-200 bg-white dark:border-dark-700 dark:bg-dark-900">
        <div class="mx-auto max-w-7xl px-4 py-4 sm:px-6 lg:px-8">
          <div class="flex items-center justify-between">
            <h1 class="text-2xl font-bold text-gray-900 dark:text-white">
              {{ t('guide.title') }}
            </h1>
            <router-link to="/home" class="btn btn-secondary">
              <Icon name="home" size="sm" class="mr-2" />
              {{ t('common.backToHome') }}
            </router-link>
          </div>
        </div>
      </div>

      <!-- Content -->
      <div class="mx-auto max-w-7xl px-4 py-8 sm:px-6 lg:px-8">
        <div class="grid min-w-0 gap-8 lg:grid-cols-[280px_minmax(0,1fr)]">
          <!-- Navigation Sidebar -->
          <aside class="hidden lg:block">
            <nav class="sticky top-8 space-y-1 rounded-lg border border-gray-200 bg-white p-4 dark:border-dark-700 dark:bg-dark-900">
              <h2 class="mb-3 text-sm font-semibold text-gray-900 dark:text-white">
                {{ t('guide.tableOfContents') }}
              </h2>
              <a
                v-for="section in sections"
                :key="section.id"
                :href="`#${section.id}`"
                class="block rounded-md px-3 py-2 text-sm transition-colors hover:bg-gray-50 dark:hover:bg-dark-800"
                :class="activeSection === section.id ? 'bg-primary-50 text-primary-600 dark:bg-primary-900/20 dark:text-primary-400' : 'text-gray-700 dark:text-gray-300'"
                @click.prevent="scrollToSection(section.id)"
              >
                {{ section.title }}
              </a>
            </nav>
          </aside>

          <!-- Main Content -->
          <main class="min-w-0 space-y-8">
            <!-- Purchase Channels -->
            <section id="purchase" class="rounded-lg border border-gray-200 bg-white p-6 dark:border-dark-700 dark:bg-dark-900">
              <h2 class="mb-6 text-xl font-bold text-gray-900 dark:text-white">
                {{ t('guide.purchaseChannels') }}
              </h2>

              <div class="space-y-6">
                <div>
                  <h3 class="mb-3 text-lg font-semibold text-gray-800 dark:text-gray-100">
                    {{ t('guide.shop1') }}
                  </h3>
                  <img
                    src="/docs/images/ddfdf1c88c71087a5b3ca819813b5a29.jpg"
                    alt="一号店铺"
                    class="h-auto w-64 rounded-lg shadow-md"
                  />
                </div>

                <div>
                  <h3 class="mb-3 text-lg font-semibold text-gray-800 dark:text-gray-100">
                    {{ t('guide.shop2') }}
                  </h3>
                  <img
                    src="/docs/images/857368a89cf4a257a74d48563e4c4412.jpg"
                    alt="二号店铺"
                    class="h-auto w-64 rounded-lg shadow-md"
                  />
                </div>

                <div>
                  <h3 class="mb-3 text-lg font-semibold text-gray-800 dark:text-gray-100">
                    {{ t('guide.afterSalesGroup') }}
                  </h3>
                  <div v-if="loadingSettings" class="h-64 w-64 animate-pulse rounded-lg bg-gray-200 dark:bg-dark-800"></div>
                  <img
                    v-else-if="afterSalesQRCode"
                    :src="afterSalesQRCode"
                    alt="售后群"
                    class="h-auto w-64 rounded-lg shadow-md"
                  />
                  <img
                    v-else
                    src="/docs/images/9c39d05c1ce5b1076dd4d8a37efe9c20.jpg"
                    alt="售后群（默认）"
                    class="h-auto w-64 rounded-lg shadow-md"
                  />
                </div>
              </div>
            </section>

            <!-- Important Notice -->
            <section id="notice" class="rounded-lg border border-gray-200 bg-white p-6 dark:border-dark-700 dark:bg-dark-900">
              <h2 class="mb-6 text-xl font-bold text-gray-900 dark:text-white">
                {{ t('guide.importantNotice') }}
              </h2>

              <div class="space-y-4">
                <div class="rounded-lg border border-amber-200 bg-amber-50 p-4 dark:border-amber-800 dark:bg-amber-950/30">
                  <ul class="list-inside list-disc space-y-2 text-sm text-amber-900 dark:text-amber-100">
                    <li>{{ t('guide.notice1') }}</li>
                    <li>{{ t('guide.notice2') }}</li>
                  </ul>
                </div>

                <div class="space-y-3 text-sm text-gray-700 dark:text-gray-300">
                  <p class="font-semibold text-gray-900 dark:text-white">{{ t('guide.usageSteps') }}</p>
                  <ol class="list-inside list-decimal space-y-2 pl-4">
                    <li>{{ t('guide.step1') }}</li>
                    <li>{{ t('guide.step2') }}</li>
                    <li>{{ t('guide.step3') }}</li>
                    <li>{{ t('guide.step4') }}</li>
                    <li>{{ t('guide.step5') }}</li>
                  </ol>
                </div>

                <div class="rounded-lg bg-gray-50 p-4 dark:bg-dark-800">
                  <p class="mb-2 text-sm font-semibold text-gray-900 dark:text-white">
                    {{ t('guide.keyCreationNote') }}
                  </p>
                  <img
                    src="/docs/images/64fb15c4720297bb212bf4e83c7ec398.png"
                    alt="密钥创建注意事项"
                    class="h-auto max-h-48 max-w-full rounded-md shadow-md"
                  />
                </div>

                <div class="rounded-lg border border-amber-200 bg-amber-50 p-4 dark:border-amber-800 dark:bg-amber-950/30">
                  <p class="mb-2 text-sm font-semibold text-amber-900 dark:text-amber-100">
                    注意：使用余额充值，请选余额分组
                  </p>
                  <img
                    src="/docs/images/image%201.png"
                    alt="余额分组选择"
                    class="h-auto max-h-64 max-w-full rounded-md shadow-md"
                  />
                </div>

                <div class="rounded-lg bg-gray-50 p-4 dark:bg-dark-800">
                  <p class="mb-3 text-sm text-gray-700 dark:text-gray-300">
                    如需<strong>指定模型</strong>，可更新 CC-SWITCH（CCS）配置，或手动更新到用户目录下的
                    <code class="rounded bg-gray-200 px-1 py-0.5 text-xs dark:bg-dark-700">.claude/settings.json</code>：
                  </p>
                  <GuideCodeBlock>
                    <code>{
  "env": {
    "ANTHROPIC_AUTH_TOKEN": "sk-你的APIKey",
    "ANTHROPIC_BASE_URL": "https://api.superwhv.me",
    "ANTHROPIC_MODEL": "claude-opus-4-6",
    "ANTHROPIC_DEFAULT_SONNET_MODEL": "claude-sonnet-4-6",
    "ANTHROPIC_DEFAULT_HAIKU_MODEL": "claude-haiku-4-5-20251001",
    "ANTHROPIC_DEFAULT_OPUS_MODEL": "claude-opus-4-6"
  },
  "model": "opus"
}</code>
                  </GuideCodeBlock>
                </div>
              </div>
            </section>

            <!-- Model List Reference -->
            <section id="models" class="rounded-lg border border-gray-200 bg-white p-6 dark:border-dark-700 dark:bg-dark-900">
              <h2 class="mb-6 text-xl font-bold text-gray-900 dark:text-white">
                {{ t('guide.modelList') }}（不限于）
              </h2>

              <div class="space-y-6">
                <div class="rounded-lg border border-sky-200 bg-sky-50 p-4 dark:border-sky-800 dark:bg-sky-950/30">
                  <p class="text-sm text-sky-900 dark:text-sky-100">
                    {{ t('guide.modelListHint') }}
                  </p>
                  <GuideCodeBlock class="mt-2">
                    <code>curl -H "Authorization: Bearer sk-你的APIKey" https://api.superwhv.me/v1/models</code>
                  </GuideCodeBlock>
                </div>

                <div class="overflow-x-auto">
                  <table class="w-full text-left text-sm">
                    <thead class="border-b border-gray-200 bg-gray-50 dark:border-dark-700 dark:bg-dark-800">
                      <tr>
                        <th class="px-4 py-3 font-semibold text-gray-900 dark:text-white">{{ t('guide.modelName') }}</th>
                        <th class="px-4 py-3 font-semibold text-gray-900 dark:text-white">{{ t('guide.modelDescription') }}</th>
                        <th class="px-4 py-3 font-semibold text-gray-900 dark:text-white">{{ t('guide.modelContextWindow') }}</th>
                      </tr>
                    </thead>
                    <tbody class="divide-y divide-gray-200 dark:divide-dark-700">
                      <tr v-for="model in models" :key="model.name">
                        <td class="px-4 py-3 font-mono text-xs text-gray-900 dark:text-white">{{ model.name }}</td>
                        <td class="px-4 py-3 text-gray-700 dark:text-gray-300">{{ model.description }}</td>
                        <td class="px-4 py-3 text-gray-700 dark:text-gray-300">{{ model.context }}</td>
                      </tr>
                    </tbody>
                  </table>
                </div>
              </div>
            </section>

            <!-- Context Window Note -->
            <section id="context-window" class="rounded-lg border border-amber-200 bg-amber-50 p-6 dark:border-amber-800 dark:bg-amber-950/30">
              <h3 class="mb-4 text-lg font-semibold text-amber-900 dark:text-amber-100">400 状态码 context window exceeded 问题</h3>
              <div class="space-y-4 text-sm text-amber-900 dark:text-amber-100">
                <p>上下文超限，即便是支持 1M context 的模型，在反代环境下仍可能触发 400 状态码。不光我们，其他站也有这个问题。</p>
                <p class="font-medium">永久解决方案（添加到环境变量）：</p>
                <GuideCodeBlock>
                  <code>CLAUDE_CODE_AUTO_COMPACT_WINDOW=200000
CLAUDE_CODE_DISABLE_1M_CONTEXT=1</code>
                </GuideCodeBlock>
                <p>让这个生效，超长的上下文会让模型性能下降（降智），如非必要，不用 1M 上下文，反而更聪明。</p>
              </div>
            </section>

            <!-- Client Installation Guides -->
            <section id="installation" class="rounded-lg border border-gray-200 bg-white p-6 dark:border-dark-700 dark:bg-dark-900">
              <h2 class="mb-6 text-xl font-bold text-gray-900 dark:text-white">
                {{ t('guide.installationGuide') }}
              </h2>

              <div class="space-y-6">
                <div class="rounded-lg border border-sky-200 bg-sky-50 p-4 dark:border-sky-800 dark:bg-sky-950/30">
                  <p class="text-sm text-sky-900 dark:text-sky-100">
                    {{ t('guide.installationHint') }}
                  </p>
                </div>

                <!-- Client Tabs -->
                <div class="space-y-4">
                  <div class="flex flex-wrap gap-2">
                    <button
                      v-for="client in clients"
                      :key="client.id"
                      type="button"
                      class="rounded-lg px-4 py-2 text-sm font-medium transition-colors"
                      :class="activeClient === client.id
                        ? 'bg-primary-600 text-white'
                        : 'bg-gray-100 text-gray-700 hover:bg-gray-200 dark:bg-dark-800 dark:text-gray-300 dark:hover:bg-dark-700'
                      "
                      @click="activeClient = client.id"
                    >
                      {{ client.name }}
                    </button>
                  </div>

                  <!-- Client Content -->
                  <div class="space-y-4">
                    <!-- CC-SWITCH -->
                    <div v-show="activeClient === 'cc-switch'" class="space-y-4">
                      <div class="rounded-lg border border-blue-200 bg-blue-50 p-4 dark:border-blue-800 dark:bg-blue-950/30">
                        <h4 class="mb-2 text-sm font-semibold text-blue-900 dark:text-blue-100">桌面管理工具（强烈建议先安装）</h4>
                        <p class="text-sm text-blue-800 dark:text-blue-200">
                          CC Switch 提供了一个桌面应用程序，用于管理 Claude Code、Codex、OpenCode、OpenClaw、Gemini CLI、Hermes Agent 等工具。
                          无需手动编辑配置文件，即可通过可视化界面一键导入提供商，并在它们之间即时切换。
                          同时平台 API 密钥可以通过 CCS 直接一键导入。
                        </p>
                      </div>

                      <div>
                        <h4 class="mb-3 font-semibold text-gray-900 dark:text-white">下载地址</h4>
                        <p class="text-sm text-gray-600 dark:text-gray-400">
                          开源地址：<a href="https://github.com/farion1231/cc-switch" target="_blank" rel="noopener" class="text-blue-600 underline dark:text-blue-400">https://github.com/farion1231/cc-switch</a>
                        </p>
                      </div>

                      <div>
                        <h4 class="mb-3 font-semibold text-gray-900 dark:text-white">配置步骤</h4>
                        <ol class="list-decimal space-y-2 pl-5 text-sm text-gray-600 dark:text-gray-400">
                          <li>下载并安装 CC-SWITCH</li>
                          <li>打开应用，点击"添加配置"</li>
                          <li>输入配置名称（如 "Superwhv"）</li>
                          <li>填入 API Key 和 Base URL（可通过 CCS 一键导入平台密钥）</li>
                          <li>保存并切换到该配置</li>
                        </ol>
                      </div>

                      <div>
                        <h4 class="mb-3 font-semibold text-gray-900 dark:text-white">配置参数</h4>
                        <ul class="list-disc space-y-1 pl-5 text-sm text-gray-600 dark:text-gray-400">
                          <li>API Key: 你的密钥</li>
                          <li>Base URL: <code class="rounded bg-gray-100 px-1 dark:bg-dark-800">https://api.superwhv.me</code></li>
                        </ul>
                      </div>
                    </div>

                    <!-- Claude Code -->
                    <div v-show="activeClient === 'claude-code'" class="space-y-6">
                      <div class="rounded-lg border border-blue-200 bg-blue-50 p-4 dark:border-blue-800 dark:bg-blue-950/30">
                        <h4 class="mb-2 text-sm font-semibold text-blue-900 dark:text-blue-100">Claude Code 安装教程</h4>
                        <p class="text-sm text-blue-800 dark:text-blue-200">支持 Claude Code / Codex 等官方客户端，直连 Superwhv API 中转。选择你的系统：Windows / macOS / Linux。</p>
                      </div>

                      <!-- Windows -->
                      <div>
                        <h4 class="mb-3 text-lg font-semibold text-gray-900 dark:text-white">Windows</h4>

                        <div class="mb-4 rounded-lg border border-gray-200 bg-gray-50 p-4 dark:border-dark-700 dark:bg-dark-800">
                          <p class="mb-2 text-sm font-semibold text-gray-900 dark:text-white">前置要求：安装 Git</p>
                          <p class="mb-2 text-sm text-gray-600 dark:text-gray-400">
                            Claude Code 需要 Git 才能运行。如果启动时遇到以下错误：
                          </p>
                          <GuideCodeBlock>
                            <code>Claude Code on Windows requires git-bash (https://git-scm.com/downloads/win).
If installed but not in PATH, set environment variable pointing to your bash.exe,
similar to: CLAUDE_CODE_GIT_BASH_PATH=C:\Program Files\Git\bin\bash.exe</code>
                          </GuideCodeBlock>
                          <p class="mt-2 text-sm text-gray-600 dark:text-gray-400">请先安装 Git：</p>
                          <ul class="mt-1 list-disc space-y-1 pl-5 text-sm text-gray-600 dark:text-gray-400">
                            <li>官方下载：<a href="https://git-scm.com/downloads/win" target="_blank" rel="noopener" class="text-blue-600 underline dark:text-blue-400">https://git-scm.com/downloads/win</a></li>
                            <li>国内加速下载：<a href="https://registry.npmmirror.com/binary.html?path=git-for-windows/" target="_blank" rel="noopener" class="text-blue-600 underline dark:text-blue-400">npmmirror 镜像</a></li>
                          </ul>
                          <p class="mt-2 text-sm text-gray-600 dark:text-gray-400">
                            如果已安装但不在 PATH 中，设置环境变量（路径改为你实际的 Git 安装路径）：
                          </p>
                          <GuideCodeBlock class="mt-2">
                            <code>CLAUDE_CODE_GIT_BASH_PATH=C:\Program Files\Git\bin\bash.exe</code>
                          </GuideCodeBlock>
                          <p class="mt-2 text-sm text-gray-600 dark:text-gray-400">验证安装：</p>
                          <GuideCodeBlock class="mt-2">
                            <code>git --version</code>
                          </GuideCodeBlock>
                        </div>

                        <div class="mb-4 rounded-lg border border-red-200 bg-red-50 p-4 dark:border-red-800 dark:bg-red-950/30">
                          <p class="text-sm font-medium text-red-900 dark:text-red-100">
                            ⚠️ 必须用 PowerShell 或 Windows Terminal！CMD 是上古遗物，用它必出问题！
                          </p>
                        </div>

                        <div class="mb-4">
                          <p class="mb-3 font-semibold text-gray-900 dark:text-white">安装方式</p>
                          <div class="space-y-3">
                            <div>
                              <p class="mb-2 text-sm font-medium text-gray-900 dark:text-white">方式一：原生安装（推荐）</p>
                              <p class="mb-2 text-sm text-gray-600 dark:text-gray-400">无需 Node.js，自动后台更新，官方推荐。</p>
                              <GuideCodeBlock>
                                <code>irm https://claude.ai/install.ps1 | iex</code>
                              </GuideCodeBlock>
                            </div>
                            <div>
                              <p class="mb-2 text-sm font-medium text-gray-900 dark:text-white">方式二：WinGet 安装</p>
                              <GuideCodeBlock>
                                <code>winget install Anthropic.ClaudeCode
winget upgrade Anthropic.ClaudeCode</code>
                              </GuideCodeBlock>
                            </div>
                            <div>
                              <p class="mb-2 text-sm font-medium text-gray-900 dark:text-white">方式三：npm 安装（需要 Node.js 18+，推荐）</p>
                              <p class="mb-2 text-sm text-gray-600 dark:text-gray-400">如果网络不好，可以使用国内镜像：</p>
                              <GuideCodeBlock>
                                <code>npm install -g @anthropic-ai/claude-code
npm install -g @anthropic-ai/claude-code --registry=https://registry.npmmirror.com</code>
                              </GuideCodeBlock>
                            </div>
                          </div>
                          <p class="mt-3 text-sm text-gray-600 dark:text-gray-400">验证安装：</p>
                          <GuideCodeBlock class="mt-2">
                            <code>claude --version</code>
                          </GuideCodeBlock>
                          <p class="mt-2 text-sm text-gray-600 dark:text-gray-400">显示版本号说明安装成功！</p>
                        </div>

                        <div class="mb-4">
                          <p class="mb-3 font-semibold text-gray-900 dark:text-white">设置环境变量（配置连接到中转服务）</p>
                          <p class="mb-2 text-sm text-gray-600 dark:text-gray-400">永久设置（系统级别，需管理员权限）：</p>
                          <GuideCodeBlock>
                            <code>[System.Environment]::SetEnvironmentVariable("ANTHROPIC_BASE_URL", "https://api.superwhv.me", [System.EnvironmentVariableTarget]::Machine)
[System.Environment]::SetEnvironmentVariable("ANTHROPIC_AUTH_TOKEN", "sk-你的APIKey", [System.EnvironmentVariableTarget]::Machine)
[System.Environment]::SetEnvironmentVariable("API_TIMEOUT_MS", "3000000", [System.EnvironmentVariableTarget]::Machine)
[System.Environment]::SetEnvironmentVariable("CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC", "1", [System.EnvironmentVariableTarget]::Machine)</code>
                          </GuideCodeBlock>
                          <p class="mb-2 mt-3 text-sm text-gray-600 dark:text-gray-400">永久设置（用户级别）：</p>
                          <GuideCodeBlock>
                            <code>[System.Environment]::SetEnvironmentVariable("ANTHROPIC_BASE_URL", "https://api.superwhv.me", [System.EnvironmentVariableTarget]::User)
[System.Environment]::SetEnvironmentVariable("ANTHROPIC_AUTH_TOKEN", "sk-你的APIKey", [System.EnvironmentVariableTarget]::User)
[System.Environment]::SetEnvironmentVariable("API_TIMEOUT_MS", "3000000", [System.EnvironmentVariableTarget]::User)
[System.Environment]::SetEnvironmentVariable("CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC", "1", [System.EnvironmentVariableTarget]::User)</code>
                          </GuideCodeBlock>
                          <p class="mb-2 mt-3 text-sm text-gray-600 dark:text-gray-400">临时设置（仅当前窗口有效）：</p>
                          <GuideCodeBlock>
                            <code>$env:ANTHROPIC_BASE_URL = "https://api.superwhv.me"
$env:ANTHROPIC_AUTH_TOKEN = "sk-你的APIKey"
$env:API_TIMEOUT_MS = "3000000"
$env:CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC = "1"</code>
                          </GuideCodeBlock>
                          <p class="mb-2 mt-3 text-sm text-gray-600 dark:text-gray-400">验证环境变量：</p>
                          <GuideCodeBlock>
                            <code>Get-ChildItem Env:ANTHROPIC_*</code>
                          </GuideCodeBlock>
                          <p class="mb-2 mt-3 text-sm text-gray-600 dark:text-gray-400">不再使用？删除环境变量：</p>
                          <GuideCodeBlock>
                            <code>[System.Environment]::SetEnvironmentVariable("ANTHROPIC_BASE_URL", $null, [System.EnvironmentVariableTarget]::Machine)
[System.Environment]::SetEnvironmentVariable("ANTHROPIC_AUTH_TOKEN", $null, [System.EnvironmentVariableTarget]::Machine)
[System.Environment]::SetEnvironmentVariable("API_TIMEOUT_MS", $null, [System.EnvironmentVariableTarget]::Machine)
[System.Environment]::SetEnvironmentVariable("CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC", $null, [System.EnvironmentVariableTarget]::Machine)</code>
                          </GuideCodeBlock>
                        </div>

                        <div>
                          <p class="mb-3 font-semibold text-gray-900 dark:text-white">开始使用</p>
                          <GuideCodeBlock>
                            <code>cd C:\path\to\your\project
claude</code>
                          </GuideCodeBlock>
                        </div>
                      </div>

                      <!-- macOS -->
                      <div>
                        <h4 class="mb-3 text-lg font-semibold text-gray-900 dark:text-white">macOS</h4>

                        <div class="mb-4 rounded-lg border border-gray-200 bg-gray-50 p-4 dark:border-dark-700 dark:bg-dark-800">
                          <p class="mb-2 text-sm font-semibold text-gray-900 dark:text-white">前置要求：安装 Git</p>
                          <GuideCodeBlock>
                            <code>brew install git
git --version</code>
                          </GuideCodeBlock>
                        </div>

                        <div class="mb-4">
                          <p class="mb-3 font-semibold text-gray-900 dark:text-white">安装方式</p>
                          <div class="space-y-3">
                            <div>
                              <p class="mb-2 text-sm font-medium text-gray-900 dark:text-white">方式一：原生安装（推荐）</p>
                              <GuideCodeBlock>
                                <code>curl -fsSL https://claude.ai/install.sh | sh</code>
                              </GuideCodeBlock>
                            </div>
                            <div>
                              <p class="mb-2 text-sm font-medium text-gray-900 dark:text-white">方式二：Homebrew 安装</p>
                              <GuideCodeBlock>
                                <code>brew install claude</code>
                              </GuideCodeBlock>
                            </div>
                            <div>
                              <p class="mb-2 text-sm font-medium text-gray-900 dark:text-white">方式三：npm 安装</p>
                              <GuideCodeBlock>
                                <code>npm install -g @anthropic-ai/claude-code</code>
                              </GuideCodeBlock>
                            </div>
                          </div>
                          <p class="mt-3 text-sm text-gray-600 dark:text-gray-400">验证安装：</p>
                          <GuideCodeBlock class="mt-2">
                            <code>claude --version</code>
                          </GuideCodeBlock>
                        </div>

                        <div class="mb-4">
                          <p class="mb-3 font-semibold text-gray-900 dark:text-white">设置环境变量</p>
                          <GuideCodeBlock>
                            <code># 永久设置（添加到 ~/.zshrc 或 ~/.bash_profile）
echo 'export ANTHROPIC_BASE_URL="https://api.superwhv.me"' >> ~/.zshrc
echo 'export ANTHROPIC_AUTH_TOKEN="sk-你的APIKey"' >> ~/.zshrc
echo 'export API_TIMEOUT_MS="3000000"' >> ~/.zshrc
echo 'export CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC="1"' >> ~/.zshrc

# 使配置生效
source ~/.zshrc</code>
                          </GuideCodeBlock>
                        </div>

                        <div>
                          <p class="mb-3 font-semibold text-gray-900 dark:text-white">开始使用</p>
                          <GuideCodeBlock>
                            <code>cd /path/to/your/project
claude</code>
                          </GuideCodeBlock>
                        </div>
                      </div>

                      <!-- Linux -->
                      <div>
                        <h4 class="mb-3 text-lg font-semibold text-gray-900 dark:text-white">Linux</h4>

                        <div class="mb-4 rounded-lg border border-gray-200 bg-gray-50 p-4 dark:border-dark-700 dark:bg-dark-800">
                          <p class="mb-2 text-sm font-semibold text-gray-900 dark:text-white">前置要求：安装 Git</p>
                          <GuideCodeBlock>
                            <code>sudo apt update && sudo apt install git
git --version</code>
                          </GuideCodeBlock>
                        </div>

                        <div class="mb-4">
                          <p class="mb-3 font-semibold text-gray-900 dark:text-white">安装方式</p>
                          <div class="space-y-3">
                            <div>
                              <p class="mb-2 text-sm font-medium text-gray-900 dark:text-white">方式一：原生安装（推荐）</p>
                              <GuideCodeBlock>
                                <code>curl -fsSL https://claude.ai/install.sh | sh</code>
                              </GuideCodeBlock>
                            </div>
                            <div>
                              <p class="mb-2 text-sm font-medium text-gray-900 dark:text-white">方式二：npm 安装</p>
                              <GuideCodeBlock>
                                <code>sudo npm install -g @anthropic-ai/claude-code</code>
                              </GuideCodeBlock>
                            </div>
                          </div>
                          <p class="mt-3 text-sm text-gray-600 dark:text-gray-400">验证安装：</p>
                          <GuideCodeBlock class="mt-2">
                            <code>claude --version</code>
                          </GuideCodeBlock>
                        </div>

                        <div class="mb-4">
                          <p class="mb-3 font-semibold text-gray-900 dark:text-white">设置环境变量</p>
                          <GuideCodeBlock>
                            <code># 永久设置（添加到 ~/.bashrc 或 ~/.zshrc）
echo 'export ANTHROPIC_BASE_URL="https://api.superwhv.me"' >> ~/.bashrc
echo 'export ANTHROPIC_AUTH_TOKEN="sk-你的APIKey"' >> ~/.bashrc
echo 'export API_TIMEOUT_MS="3000000"' >> ~/.bashrc
echo 'export CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC="1"' >> ~/.bashrc

# 使配置生效
source ~/.bashrc</code>
                          </GuideCodeBlock>
                        </div>

                        <div>
                          <p class="mb-3 font-semibold text-gray-900 dark:text-white">开始使用</p>
                          <GuideCodeBlock>
                            <code>cd /path/to/your/project
claude</code>
                          </GuideCodeBlock>
                        </div>
                      </div>

                      <div class="rounded-lg border border-amber-200 bg-amber-50 p-4 dark:border-amber-800 dark:bg-amber-950/30">
                        <p class="text-sm text-amber-900 dark:text-amber-100">💡 `sk-你的APIKey` 请替换为你在后台申请的 API Key！</p>
                      </div>
                    </div>

                    <!-- Claude APP -->
                    <div v-show="activeClient === 'claude-app'" class="space-y-4">
                      <div class="rounded-lg border border-purple-200 bg-purple-50 p-4 dark:border-purple-800 dark:bg-purple-950/30">
                        <h4 class="mb-2 text-sm font-semibold text-purple-900 dark:text-purple-100">Claude 官方桌面应用</h4>
                        <p class="text-sm text-purple-800 dark:text-purple-200">免账号登录模式（国内用户推荐）</p>
                      </div>

                      <div>
                        <h4 class="mb-3 font-semibold text-gray-900 dark:text-white">安装地址</h4>
                        <p class="mb-2 text-sm text-gray-600 dark:text-gray-400">
                          <a href="https://claude.com/download" target="_blank" rel="noopener" class="text-blue-600 underline dark:text-blue-400">Download Claude | Claude by Anthropic</a>
                        </p>
                        <div class="rounded-lg border border-amber-200 bg-amber-50 p-4 dark:border-amber-800 dark:bg-amber-950/30">
                          <p class="text-sm text-amber-900 dark:text-amber-100">
                            ⚠️ 下载之后是一个安装包，点击后会再次下载内容，这一步需要打开梯子的 TUN（虚拟网卡）模式
                          </p>
                        </div>
                      </div>

                      <div>
                        <h4 class="mb-3 font-semibold text-gray-900 dark:text-white">开启开发者模式</h4>
                        <ol class="list-decimal space-y-2 pl-5 text-sm text-gray-600 dark:text-gray-400">
                          <li>首次打开，先不着急登录，教程全程不需要 Claude 账号</li>
                          <li>打开左上角的菜单按钮 —— 一开始可能无法点击，需要通过鼠标点击邮箱输入框（不用输入内容，只是为了触发聚焦），连按键盘的 Tab 键就能选中此按钮了</li>
                          <li>菜单出现之后就可以用鼠标选择了，依次点击 Help → Troubleshooting → Enable Developer Mode</li>
                        </ol>
                        <img
                          src="/docs/images/tVVLh9ni.png"
                          alt="开启开发者模式"
                          class="mt-3 h-auto max-h-96 max-w-full rounded-md shadow-md"
                        />
                      </div>

                      <div>
                        <h4 class="mb-3 font-semibold text-gray-900 dark:text-white">配置第三方提供商</h4>
                        <p class="mb-2 text-sm text-gray-600 dark:text-gray-400">
                          打开界面左上角的 <code class="rounded bg-gray-100 px-1 dark:bg-dark-800">developer</code>，选择
                          <code class="rounded bg-gray-100 px-1 dark:bg-dark-800">Configure Third-Party Inference</code>，
                          进入第三方接口配置界面（确保已开启开发者模式）。
                        </p>
                        <img
                          src="/docs/images/E1ZVuXlc.png"
                          alt="开发者菜单"
                          class="mt-3 h-auto max-h-96 max-w-full rounded-md shadow-md"
                        />
                        <img
                          src="/docs/images/image%202.png"
                          alt="第三方接口配置"
                          class="mt-3 h-auto max-h-96 max-w-full rounded-md shadow-md"
                        />
                        <p class="mt-3 text-sm text-gray-600 dark:text-gray-400">三个配置更改后：</p>
                        <ul class="mt-2 list-disc space-y-1 pl-5 text-sm text-gray-600 dark:text-gray-400">
                          <li>Base URL: <code class="rounded bg-gray-100 px-1 dark:bg-dark-800">https://api.superwhv.me</code></li>
                          <li>API Key: <code class="rounded bg-gray-100 px-1 dark:bg-dark-800">sk-你的APIKey</code></li>
                          <li>Model: 选择你需要的模型</li>
                        </ul>
                        <p class="mt-3 text-sm font-medium text-gray-900 dark:text-white">
                          Apply locally（！！！一定要应用重启），然后就可以原版体验了。
                        </p>
                        <img
                          src="/docs/images/image.png"
                          alt="配置完成界面"
                          class="mt-3 h-auto max-h-96 max-w-full rounded-md shadow-md"
                        />
                      </div>

                      <div class="rounded-lg border border-amber-200 bg-amber-50 p-4 dark:border-amber-800 dark:bg-amber-950/30">
                        <p class="text-sm text-amber-900 dark:text-amber-100">💡 确保开启开发者模式，否则第三方配置不生效</p>
                      </div>
                    </div>

                    <!-- Codex TUI -->
                    <div v-show="activeClient === 'codex'" class="space-y-6">
                      <div class="rounded-lg border border-blue-200 bg-blue-50 p-4 dark:border-blue-800 dark:bg-blue-950/30">
                        <h4 class="mb-2 text-sm font-semibold text-blue-900 dark:text-blue-100">Codex TUI 安装教程</h4>
                        <p class="text-sm text-blue-800 dark:text-blue-200">支持 Windows / macOS / Linux，TUI 模式，适合 CLI 用户。选择你的系统：Windows / macOS / Linux。</p>
                      </div>

                      <!-- Windows -->
                      <div>
                        <h4 class="mb-3 text-lg font-semibold text-gray-900 dark:text-white">Windows</h4>

                        <div class="mb-4 rounded-lg border border-gray-200 bg-gray-50 p-4 dark:border-dark-700 dark:bg-dark-800">
                          <p class="mb-2 text-sm font-semibold text-gray-900 dark:text-white">前置要求：安装 Node.js</p>
                          <p class="mb-2 text-sm text-gray-600 dark:text-gray-400">检查是否已安装：</p>
                          <GuideCodeBlock>
                            <code>node --version
npm --version</code>
                          </GuideCodeBlock>
                          <p class="mt-2 text-sm text-gray-600 dark:text-gray-400">
                            如果命令不存在，访问 Node.js 官网下载 Windows 对应的 LTS 安装包。
                          </p>
                        </div>

                        <div class="mb-4">
                          <p class="mb-3 font-semibold text-gray-900 dark:text-white">安装 codex</p>
                          <GuideCodeBlock>
                            <code>npm i -g @openai/codex --registry=https://registry.npmmirror.com
codex --version</code>
                          </GuideCodeBlock>
                        </div>

                        <div class="mb-4">
                          <p class="mb-3 font-semibold text-gray-900 dark:text-white">配置 codex</p>
                          <p class="mb-2 text-sm text-gray-600 dark:text-gray-400">
                            打开文件资源管理器，找到 <code class="rounded bg-gray-100 px-1 dark:bg-dark-800">C:\Users\你的用户名\.codex</code> 文件夹（不存在则创建）。
                          </p>
                          <p class="mb-2 text-sm text-gray-600 dark:text-gray-400">创建 <code class="rounded bg-gray-100 px-1 dark:bg-dark-800">config.toml</code> 文件：</p>
                          <GuideCodeBlock>
                            <code>model_provider = "superwhv"
model = "gpt-5.3-codex"
model_reasoning_effort = "high"
network_access = "enabled"
disable_response_storage = true
windows_wsl_setup_acknowledged = true
model_verbosity = "high"

[model_providers.superwhv]
name = "superwhv"
base_url = "https://api.superwhv.me"
wire_api = "responses"
requires_openai_auth = true</code>
                          </GuideCodeBlock>
                          <p class="mb-2 mt-3 text-sm text-gray-600 dark:text-gray-400">创建 <code class="rounded bg-gray-100 px-1 dark:bg-dark-800">auth.json</code> 文件：</p>
                          <GuideCodeBlock>
                            <code>{
  "OPENAI_API_KEY": "sk-你的APIKey"
}</code>
                          </GuideCodeBlock>
                          <p class="mt-2 text-sm font-medium text-red-600 dark:text-red-400">⚠️ 将 `sk-你的APIKey` 替换为你的 API Key！</p>
                        </div>

                        <div>
                          <p class="mb-3 font-semibold text-gray-900 dark:text-white">启动 codex</p>
                          <GuideCodeBlock>
                            <code>cd C:\path\to\your\project
codex</code>
                          </GuideCodeBlock>
                          <p class="mt-2 text-sm text-gray-600 dark:text-gray-400">首次启动时，codex 会进行初始化配置。</p>
                        </div>

                        <div class="mt-4">
                          <p class="mb-3 font-semibold text-gray-900 dark:text-white">编辑器插件使用</p>
                          <p class="mb-2 text-sm text-gray-600 dark:text-gray-400">
                            在插件市场搜索并安装 <code class="rounded bg-gray-100 px-1 dark:bg-dark-800">Codex - OpenAI's coding agent</code>。
                          </p>
                          <p class="mb-2 text-sm text-gray-600 dark:text-gray-400">
                            VS Code、Cursor、Trae、Kiro 操作基本一样，只要前面的 Codex 配置已经完成，插件打开后就能直接使用。
                          </p>
                          <p class="text-sm text-gray-600 dark:text-gray-400">
                            第一次使用时，记得在输入框右上角把 <strong>Sandbox（沙盒）</strong>打开。
                          </p>
                        </div>
                      </div>

                      <!-- macOS -->
                      <div>
                        <h4 class="mb-3 text-lg font-semibold text-gray-900 dark:text-white">macOS</h4>

                        <div class="mb-4 rounded-lg border border-gray-200 bg-gray-50 p-4 dark:border-dark-700 dark:bg-dark-800">
                          <p class="mb-2 text-sm font-semibold text-gray-900 dark:text-white">前置要求：安装 Node.js</p>
                          <GuideCodeBlock>
                            <code>brew install node
node --version
npm --version</code>
                          </GuideCodeBlock>
                        </div>

                        <div class="mb-4">
                          <p class="mb-3 font-semibold text-gray-900 dark:text-white">安装 codex</p>
                          <GuideCodeBlock>
                            <code>npm i -g @openai/codex --registry=https://registry.npmmirror.com
codex --version</code>
                          </GuideCodeBlock>
                        </div>

                        <div class="mb-4">
                          <p class="mb-3 font-semibold text-gray-900 dark:text-white">配置 codex</p>
                          <GuideCodeBlock>
                            <code>mkdir -p ~/.codex</code>
                          </GuideCodeBlock>
                          <p class="mb-2 mt-3 text-sm text-gray-600 dark:text-gray-400">创建 <code class="rounded bg-gray-100 px-1 dark:bg-dark-800">~/.codex/config.toml</code> 文件：</p>
                          <GuideCodeBlock>
                            <code>model_provider = "superwhv"
model = "gpt-5.3-codex"
model_reasoning_effort = "high"
network_access = "enabled"
disable_response_storage = true
model_verbosity = "high"

[model_providers.superwhv]
name = "superwhv"
base_url = "https://api.superwhv.me"
wire_api = "responses"
requires_openai_auth = true</code>
                          </GuideCodeBlock>
                          <p class="mb-2 mt-3 text-sm text-gray-600 dark:text-gray-400">创建 <code class="rounded bg-gray-100 px-1 dark:bg-dark-800">~/.codex/auth.json</code> 文件：</p>
                          <GuideCodeBlock>
                            <code>{
  "OPENAI_API_KEY": "sk-你的APIKey"
}</code>
                          </GuideCodeBlock>
                          <p class="mt-2 text-sm font-medium text-red-600 dark:text-red-400">⚠️ 将 `sk-你的APIKey` 替换为你的 API Key！</p>
                        </div>

                        <div>
                          <p class="mb-3 font-semibold text-gray-900 dark:text-white">启动 codex</p>
                          <GuideCodeBlock>
                            <code>cd /path/to/your/project
codex</code>
                          </GuideCodeBlock>
                        </div>
                      </div>

                      <!-- Linux -->
                      <div>
                        <h4 class="mb-3 text-lg font-semibold text-gray-900 dark:text-white">Linux</h4>

                        <div class="mb-4 rounded-lg border border-gray-200 bg-gray-50 p-4 dark:border-dark-700 dark:bg-dark-800">
                          <p class="mb-2 text-sm font-semibold text-gray-900 dark:text-white">前置要求：安装 Node.js</p>
                          <GuideCodeBlock>
                            <code>sudo apt update && sudo apt install nodejs npm
node --version
npm --version</code>
                          </GuideCodeBlock>
                        </div>

                        <div class="mb-4">
                          <p class="mb-3 font-semibold text-gray-900 dark:text-white">安装 codex</p>
                          <GuideCodeBlock>
                            <code>sudo npm i -g @openai/codex --registry=https://registry.npmmirror.com
codex --version</code>
                          </GuideCodeBlock>
                        </div>

                        <div class="mb-4">
                          <p class="mb-3 font-semibold text-gray-900 dark:text-white">配置 codex</p>
                          <GuideCodeBlock>
                            <code>mkdir -p ~/.codex</code>
                          </GuideCodeBlock>
                          <p class="mb-2 mt-3 text-sm text-gray-600 dark:text-gray-400">创建 <code class="rounded bg-gray-100 px-1 dark:bg-dark-800">~/.codex/config.toml</code> 文件：</p>
                          <GuideCodeBlock>
                            <code>model_provider = "superwhv"
model = "gpt-5.3-codex"
model_reasoning_effort = "high"
network_access = "enabled"
disable_response_storage = true
model_verbosity = "high"

[model_providers.superwhv]
name = "superwhv"
base_url = "https://api.superwhv.me"
wire_api = "responses"
requires_openai_auth = true</code>
                          </GuideCodeBlock>
                          <p class="mb-2 mt-3 text-sm text-gray-600 dark:text-gray-400">创建 <code class="rounded bg-gray-100 px-1 dark:bg-dark-800">~/.codex/auth.json</code> 文件：</p>
                          <GuideCodeBlock>
                            <code>{
  "OPENAI_API_KEY": "sk-你的APIKey"
}</code>
                          </GuideCodeBlock>
                          <p class="mt-2 text-sm font-medium text-red-600 dark:text-red-400">⚠️ 将 `sk-你的APIKey` 替换为你的 API Key！</p>
                        </div>

                        <div>
                          <p class="mb-3 font-semibold text-gray-900 dark:text-white">启动 codex</p>
                          <GuideCodeBlock>
                            <code>cd /path/to/your/project
codex</code>
                          </GuideCodeBlock>
                        </div>
                      </div>

                      <div>
                        <p class="mb-3 font-semibold text-gray-900 dark:text-white">Codex 常见报错</p>
                        <div class="overflow-x-auto">
                          <table class="w-full text-left text-sm">
                            <thead class="border-b border-gray-200 bg-gray-50 dark:border-dark-700 dark:bg-dark-800">
                              <tr>
                                <th class="px-4 py-3 font-semibold text-gray-900 dark:text-white">错误码</th>
                                <th class="px-4 py-3 font-semibold text-gray-900 dark:text-white">可能原因</th>
                                <th class="px-4 py-3 font-semibold text-gray-900 dark:text-white">解决方案</th>
                              </tr>
                            </thead>
                            <tbody class="divide-y divide-gray-200 dark:divide-dark-700">
                              <tr>
                                <td class="px-4 py-3 font-mono text-xs text-red-600 dark:text-red-400">401</td>
                                <td class="px-4 py-3 text-gray-700 dark:text-gray-300">分组选错（如选成了 Codex 分组）</td>
                                <td class="px-4 py-3 text-gray-700 dark:text-gray-300">检查后台分组设置，选择正确的每天额度分组</td>
                              </tr>
                              <tr>
                                <td class="px-4 py-3 font-mono text-xs text-red-600 dark:text-red-400">403</td>
                                <td class="px-4 py-3 text-gray-700 dark:text-gray-300">把兑换码当成密钥配置</td>
                                <td class="px-4 py-3 text-gray-700 dark:text-gray-300">应填写后台创建出来的 API 密钥</td>
                              </tr>
                              <tr>
                                <td class="px-4 py-3 font-mono text-xs text-red-600 dark:text-red-400">429</td>
                                <td class="px-4 py-3 text-gray-700 dark:text-gray-300">当前额度已用完</td>
                                <td class="px-4 py-3 text-gray-700 dark:text-gray-300">去后台查看剩余额度，额度恢复后再试</td>
                              </tr>
                            </tbody>
                          </table>
                        </div>
                      </div>

                      <div class="rounded-lg border border-amber-200 bg-amber-50 p-4 dark:border-amber-800 dark:bg-amber-950/30">
                        <p class="text-sm text-amber-900 dark:text-amber-100">💡 `sk-你的APIKey` 请替换为你的 API Key！</p>
                      </div>
                    </div>

                    <!-- Codex APP -->
                    <div v-show="activeClient === 'codex-app'" class="space-y-4">
                      <div class="rounded-lg border border-green-200 bg-green-50 p-4 dark:border-green-800 dark:bg-green-950/30">
                        <h4 class="mb-2 text-sm font-semibold text-green-900 dark:text-green-100">Codex 桌面应用</h4>
                        <p class="text-sm text-green-800 dark:text-green-200">Codex 的图形界面客户端，提供更友好的交互体验。</p>
                      </div>

                      <div>
                        <h4 class="mb-3 font-semibold text-gray-900 dark:text-white">下载地址</h4>
                        <p class="mb-2 text-sm text-gray-600 dark:text-gray-400">
                          <a href="https://chatgpt.com/codex" target="_blank" rel="noopener" class="text-blue-600 underline dark:text-blue-400">https://chatgpt.com/codex</a>，按照系统需要下载对应版本。
                        </p>
                      </div>

                      <div>
                        <h4 class="mb-3 font-semibold text-gray-900 dark:text-white">配置说明</h4>
                        <p class="mb-2 text-sm text-gray-600 dark:text-gray-400">
                          Codex APP 与 Codex TUI 的 <code class="rounded bg-gray-100 px-1 dark:bg-dark-800">config.toml</code>、
                          <code class="rounded bg-gray-100 px-1 dark:bg-dark-800">auth.json</code> 配置公用，
                          因此可以参考上面 Codex TUI 的配置步骤，使用 cc-switch 同样可以生效。
                        </p>
                        <img
                          src="/docs/images/image%203.png"
                          alt="Codex APP 界面"
                          class="mt-3 h-auto max-h-96 max-w-full rounded-md shadow-md"
                        />
                      </div>
                    </div>

                    <!-- Codex Remote -->
                    <div v-show="activeClient === 'codex-remote'" class="space-y-4">
                      <div class="rounded-lg border border-indigo-200 bg-indigo-50 p-4 dark:border-indigo-800 dark:bg-indigo-950/30">
                        <h4 class="mb-2 text-sm font-semibold text-indigo-900 dark:text-indigo-100">Codex Remote（远程控制）</h4>
                        <p class="text-sm text-indigo-800 dark:text-indigo-200">通过 Web 界面远程控制服务器上的 Codex 实例。</p>
                      </div>

                      <div class="rounded-lg border border-amber-200 bg-amber-50 p-4 dark:border-amber-800 dark:bg-amber-950/30">
                        <p class="text-sm text-amber-900 dark:text-amber-100">
                          ⚠️ 仅做指导，自测可用但不是很稳定，还请自己探索，不提供咨询服务
                        </p>
                      </div>

                      <div>
                        <h4 class="mb-3 font-semibold text-gray-900 dark:text-white">使用说明</h4>
                        <p class="mb-2 text-sm text-gray-600 dark:text-gray-400">
                          Codex 远程控制目前局限于 ChatGPT 登录用户。若想使用第三方 API 实现远程操作，可参考以下步骤：
                        </p>
                        <ol class="list-decimal space-y-2 pl-5 text-sm text-gray-600 dark:text-gray-400">
                          <li><strong>初始设置</strong>：先用 ChatGPT 账号登录并开启 Remote Control。</li>
                          <li><strong>修改配置</strong>：打开 <code class="rounded bg-gray-100 px-1 dark:bg-dark-800">config.toml</code>，添加你的 API 信息，并删除 <code class="rounded bg-gray-100 px-1 dark:bg-dark-800">requires_OpenAI_auth = True</code> 这一行。这一行命令的作用是使用 ChatGPT 账户登录，这会导致 Codex App 在初始状态时优先读取 auth.json 文件，而并非 API。</li>
                          <li><strong>设置环境变量</strong>：在 API 参数中增加 <code class="rounded bg-gray-100 px-1 dark:bg-dark-800">ENV_KEY = "OPENAI_API_KEY"</code>。</li>
                          <li><strong>绕过检测</strong>：由于系统存在文件一致性检测，不能直接替换 auth.json。请在 .codex 目录新建 <code class="rounded bg-gray-100 px-1 dark:bg-dark-800">.env</code> 文件并写入 API Key。</li>
                          <li><strong>激活环境</strong>：在终端执行 <code class="rounded bg-gray-100 px-1 dark:bg-dark-800">source</code> 启动该 .env 文件，并通过 <code class="rounded bg-gray-100 px-1 dark:bg-dark-800">echo</code> 确认。建议将此命令加入 Codex 环境设置中以自动执行。</li>
                          <li><strong>保持登录</strong>：不要退出已登录的 ChatGPT 账号，直接对话即可通过 API 通信。此时远程连接状态应显示为 Online。</li>
                        </ol>
                        <p class="mt-3 text-sm text-gray-600 dark:text-gray-400">
                          注意：设置完成后请勿关闭 Codex App，否则重新读取 auth.json 可能导致配置失效，需重新设置。
                        </p>
                      </div>

                      <div class="rounded-lg bg-gray-50 p-4 dark:bg-dark-800">
                        <p class="text-sm text-gray-600 dark:text-gray-400">
                          这个方法的原理是避开了原有的 Codex 对于 auth.json 文件一致性的检测。通过另外的 .env 文件，
                          或者直接 export 环境变量的方式，可以在保持 ChatGPT 账户登录状态的同时，切换为 API 供应模式。
                        </p>
                      </div>
                    </div>

                    <!-- Gemini CLI -->
                    <div v-show="activeClient === 'gemini'" class="space-y-6">
                      <div class="rounded-lg border border-blue-200 bg-blue-50 p-4 dark:border-blue-800 dark:bg-blue-950/30">
                        <h4 class="mb-2 text-sm font-semibold text-blue-900 dark:text-blue-100">Gemini CLI 安装教程</h4>
                        <p class="text-sm text-blue-800 dark:text-blue-200">支持 Google Gemini CLI，直连 Superwhv API。选择你的系统：Windows / macOS / Linux。</p>
                      </div>

                      <!-- Windows -->
                      <div>
                        <h4 class="mb-3 text-lg font-semibold text-gray-900 dark:text-white">Windows</h4>

                        <div class="mb-4">
                          <p class="mb-3 font-semibold text-gray-900 dark:text-white">安装 Gemini CLI</p>
                          <GuideCodeBlock>
                            <code>npm i -g @google/gemini-cli --registry=https://registry.npmmirror.com</code>
                          </GuideCodeBlock>
                        </div>

                        <div class="mb-4">
                          <p class="mb-3 font-semibold text-gray-900 dark:text-white">配置环境变量（永久，需管理员权限）</p>
                          <GuideCodeBlock>
                            <code>[System.Environment]::SetEnvironmentVariable("GEMINI_BASE_URL", "https://api.superwhv.me/v1beta", [System.EnvironmentVariableTarget]::Machine)
[System.Environment]::SetEnvironmentVariable("GEMINI_API_KEY", "sk-你的APIKey", [System.EnvironmentVariableTarget]::Machine)</code>
                          </GuideCodeBlock>
                        </div>

                        <div>
                          <p class="mb-3 font-semibold text-gray-900 dark:text-white">开始使用</p>
                          <GuideCodeBlock>
                            <code>gemini</code>
                          </GuideCodeBlock>
                        </div>
                      </div>

                      <!-- macOS -->
                      <div>
                        <h4 class="mb-3 text-lg font-semibold text-gray-900 dark:text-white">macOS</h4>

                        <div class="mb-4">
                          <p class="mb-3 font-semibold text-gray-900 dark:text-white">安装 Gemini CLI</p>
                          <GuideCodeBlock>
                            <code>npm i -g @google/gemini-cli --registry=https://registry.npmmirror.com</code>
                          </GuideCodeBlock>
                        </div>

                        <div class="mb-4">
                          <p class="mb-3 font-semibold text-gray-900 dark:text-white">配置环境变量</p>
                          <GuideCodeBlock>
                            <code>echo 'export GEMINI_BASE_URL="https://api.superwhv.me/v1beta"' >> ~/.zshrc
echo 'export GEMINI_API_KEY="sk-你的APIKey"' >> ~/.zshrc
source ~/.zshrc</code>
                          </GuideCodeBlock>
                        </div>

                        <div>
                          <p class="mb-3 font-semibold text-gray-900 dark:text-white">开始使用</p>
                          <GuideCodeBlock>
                            <code>gemini</code>
                          </GuideCodeBlock>
                        </div>
                      </div>

                      <!-- Linux -->
                      <div>
                        <h4 class="mb-3 text-lg font-semibold text-gray-900 dark:text-white">Linux</h4>

                        <div class="mb-4">
                          <p class="mb-3 font-semibold text-gray-900 dark:text-white">安装 Gemini CLI</p>
                          <GuideCodeBlock>
                            <code>sudo npm i -g @google/gemini-cli --registry=https://registry.npmmirror.com</code>
                          </GuideCodeBlock>
                        </div>

                        <div class="mb-4">
                          <p class="mb-3 font-semibold text-gray-900 dark:text-white">配置环境变量</p>
                          <GuideCodeBlock>
                            <code>echo 'export GEMINI_BASE_URL="https://api.superwhv.me/v1beta"' >> ~/.bashrc
echo 'export GEMINI_API_KEY="sk-你的APIKey"' >> ~/.bashrc
source ~/.bashrc</code>
                          </GuideCodeBlock>
                        </div>

                        <div>
                          <p class="mb-3 font-semibold text-gray-900 dark:text-white">开始使用</p>
                          <GuideCodeBlock>
                            <code>gemini</code>
                          </GuideCodeBlock>
                        </div>
                      </div>

                      <div class="rounded-lg border border-amber-200 bg-amber-50 p-4 dark:border-amber-800 dark:bg-amber-950/30">
                        <p class="text-sm text-amber-900 dark:text-amber-100">💡 确保环境变量生效后重启终端或编辑器</p>
                      </div>
                    </div>

                    <!-- OpenClaw -->
                    <div v-show="activeClient === 'openclaw'" class="space-y-4">
                      <div class="rounded-lg border border-indigo-200 bg-indigo-50 p-4 dark:border-indigo-800 dark:bg-indigo-950/30">
                        <h4 class="mb-2 text-sm font-semibold text-indigo-900 dark:text-indigo-100">OpenClaw 安装教程</h4>
                        <p class="text-sm text-indigo-800 dark:text-indigo-200">
                          OpenClaw 是一个开源个人 AI 助手系统，支持通过自定义 provider 连接中转服务，
                          可接入 Telegram、Discord 等消息渠道，也可直接使用 TUI 界面。平台已经支持 OpenClaw 接入。
                        </p>
                      </div>

                      <img
                        src="/docs/images/cef784f8eb6512eddd6afe9730dc3d85.png"
                        alt="OpenClaw 平台接入"
                        class="h-auto max-h-96 max-w-full rounded-md shadow-md"
                      />

                      <div>
                        <h4 class="mb-3 font-semibold text-gray-900 dark:text-white">系统要求</h4>
                        <ul class="list-disc space-y-1 pl-5 text-sm text-gray-600 dark:text-gray-400">
                          <li><code class="rounded bg-gray-100 px-1 dark:bg-dark-800">Node.js &gt;= 22</code></li>
                          <li>支持 macOS / Linux / Windows（Windows 推荐 WSL2 或 PowerShell）</li>
                          <li>需要先安装 Git</li>
                        </ul>
                      </div>

                      <div>
                        <h4 class="mb-3 font-semibold text-gray-900 dark:text-white">安装 OpenClaw</h4>
                        <GuideCodeBlock>
                          <code># 全局安装
npm install -g openclaw@latest

# 国内镜像加速
npm install -g openclaw@latest --registry=https://registry.npmmirror.com

# 验证安装
openclaw --version</code>
                        </GuideCodeBlock>
                      </div>

                      <div>
                        <h4 class="mb-3 font-semibold text-gray-900 dark:text-white">运行初始化向导</h4>
                        <GuideCodeBlock>
                          <code>openclaw onboard --install-daemon</code>
                        </GuideCodeBlock>
                        <p class="mb-2 mt-3 text-sm text-gray-600 dark:text-gray-400">向导中按以下方式选择：</p>
                        <div class="overflow-x-auto">
                          <table class="w-full text-left text-sm">
                            <thead class="border-b border-gray-200 bg-gray-50 dark:border-dark-700 dark:bg-dark-800">
                              <tr>
                                <th class="px-4 py-3 font-semibold text-gray-900 dark:text-white">选项</th>
                                <th class="px-4 py-3 font-semibold text-gray-900 dark:text-white">建议选择</th>
                              </tr>
                            </thead>
                            <tbody class="divide-y divide-gray-200 dark:divide-dark-700">
                              <tr><td class="px-4 py-3 text-gray-700 dark:text-gray-300">Security 安全确认</td><td class="px-4 py-3 text-gray-700 dark:text-gray-300"><code class="rounded bg-gray-100 px-1 dark:bg-dark-800">Yes</code></td></tr>
                              <tr><td class="px-4 py-3 text-gray-700 dark:text-gray-300">Onboarding mode</td><td class="px-4 py-3 text-gray-700 dark:text-gray-300"><code class="rounded bg-gray-100 px-1 dark:bg-dark-800">Manual</code>（手动配置）</td></tr>
                              <tr><td class="px-4 py-3 text-gray-700 dark:text-gray-300">What do you want to set up?</td><td class="px-4 py-3 text-gray-700 dark:text-gray-300"><code class="rounded bg-gray-100 px-1 dark:bg-dark-800">Local gateway (this machine)</code></td></tr>
                              <tr><td class="px-4 py-3 text-gray-700 dark:text-gray-300">Model / auth provider</td><td class="px-4 py-3 text-gray-700 dark:text-gray-300"><code class="rounded bg-gray-100 px-1 dark:bg-dark-800">Skip for now</code></td></tr>
                              <tr><td class="px-4 py-3 text-gray-700 dark:text-gray-300">Gateway port</td><td class="px-4 py-3 text-gray-700 dark:text-gray-300"><code class="rounded bg-gray-100 px-1 dark:bg-dark-800">18789</code>（默认）</td></tr>
                              <tr><td class="px-4 py-3 text-gray-700 dark:text-gray-300">Gateway bind</td><td class="px-4 py-3 text-gray-700 dark:text-gray-300"><code class="rounded bg-gray-100 px-1 dark:bg-dark-800">LAN (0.0.0.0)</code></td></tr>
                              <tr><td class="px-4 py-3 text-gray-700 dark:text-gray-300">Gateway auth</td><td class="px-4 py-3 text-gray-700 dark:text-gray-300"><code class="rounded bg-gray-100 px-1 dark:bg-dark-800">Token</code></td></tr>
                              <tr><td class="px-4 py-3 text-gray-700 dark:text-gray-300">Tailscale exposure</td><td class="px-4 py-3 text-gray-700 dark:text-gray-300"><code class="rounded bg-gray-100 px-1 dark:bg-dark-800">Off</code></td></tr>
                              <tr><td class="px-4 py-3 text-gray-700 dark:text-gray-300">Configure chat channels now?</td><td class="px-4 py-3 text-gray-700 dark:text-gray-300"><code class="rounded bg-gray-100 px-1 dark:bg-dark-800">No</code></td></tr>
                              <tr><td class="px-4 py-3 text-gray-700 dark:text-gray-300">Set GOOGLE_PLACES_API_KEY?</td><td class="px-4 py-3 text-gray-700 dark:text-gray-300"><code class="rounded bg-gray-100 px-1 dark:bg-dark-800">No</code></td></tr>
                              <tr><td class="px-4 py-3 text-gray-700 dark:text-gray-300">Configure skills now?</td><td class="px-4 py-3 text-gray-700 dark:text-gray-300"><code class="rounded bg-gray-100 px-1 dark:bg-dark-800">No</code></td></tr>
                            </tbody>
                          </table>
                        </div>
                        <p class="mt-2 text-sm text-gray-600 dark:text-gray-400">
                          向导结束后提示 <code class="rounded bg-gray-100 px-1 dark:bg-dark-800">Model check: No auth configured</code> 是正常的，后续手动配置中转即可。
                        </p>
                      </div>

                      <div>
                        <h4 class="mb-3 font-semibold text-gray-900 dark:text-white">配置中转服务</h4>
                        <p class="mb-2 text-sm text-gray-600 dark:text-gray-400">编辑配置文件，添加自定义 provider：</p>
                        <GuideCodeBlock>
                          <code># macOS / Linux
~/.openclaw/openclaw.json

# Windows
%USERPROFILE%\.openclaw\openclaw.json</code>
                        </GuideCodeBlock>
                        <p class="mt-2 text-sm text-gray-600 dark:text-gray-400">
                          可选分组：按你的订阅分组选择对应配置，baseUrl 和 API Key 需对应你自己的分组地址。
                        </p>
                      </div>

                      <div>
                        <h4 class="mb-3 font-semibold text-gray-900 dark:text-white">Claude 分组配置（使用 Claude 系列模型）</h4>
                        <p class="mb-2 text-sm text-gray-600 dark:text-gray-400">
                          将以下内容覆盖写入配置文件，把 <code class="rounded bg-gray-100 px-1 dark:bg-dark-800">sk-你的APIKey</code> 替换为你的密钥，
                          <code class="rounded bg-gray-100 px-1 dark:bg-dark-800">your-token</code> 替换为自定义 Gateway Token：
                        </p>
                        <GuideCodeBlock>
                          <code>{
  "agents": {
    "defaults": {
      "workspace": "~/.openclaw",
      "model": {
        "primary": "superwhv/claude-opus-4-6"
      },
      "compaction": {
        "mode": "safeguard",
        "reserveTokensFloor": 40000
      },
      "thinkingDefault": "high",
      "timeoutSeconds": 3000,
      "maxConcurrent": 3,
      "subagents": {
        "maxConcurrent": 3
      }
    }
  },
  "models": {
    "mode": "merge",
    "providers": {
      "superwhv": {
        "baseUrl": "https://api.superwhv.me",
        "apiKey": "sk-你的APIKey",
        "api": "anthropic-messages",
        "models": [
          {
            "id": "claude-sonnet-4-6",
            "name": "claude-sonnet-4-6",
            "reasoning": true,
            "cost": {
              "input": 3,
              "output": 15,
              "cacheRead": 0.3,
              "cacheWrite": 3.75
            },
            "contextWindow": 200000,
            "maxTokens": 64000
          },
          {
            "id": "claude-opus-4-6",
            "name": "claude-opus-4-6",
            "reasoning": true,
            "cost": {
              "input": 3,
              "output": 15,
              "cacheRead": 0.3,
              "cacheWrite": 3.75
            },
            "contextWindow": 200000,
            "maxTokens": 64000
          },
          {
            "id": "claude-haiku-4-5",
            "name": "claude-haiku-4-5",
            "reasoning": true,
            "cost": {
              "input": 3,
              "output": 15,
              "cacheRead": 0.3,
              "cacheWrite": 3.75
            },
            "contextWindow": 200000,
            "maxTokens": 64000
          }
        ]
      }
    }
  },
  "messages": {
    "ackReactionScope": "group-mentions"
  },
  "commands": {
    "native": "auto",
    "nativeSkills": "auto"
  },
  "gateway": {
    "port": 18789,
    "mode": "local",
    "bind": "lan",
    "auth": {
      "mode": "token",
      "token": "your-token"
    },
    "tailscale": {
      "mode": "off",
      "resetOnExit": false
    }
  },
  "skills": {
    "install": {
      "nodeManager": "npm"
    }
  }
}</code>
                        </GuideCodeBlock>
                      </div>

                      <div>
                        <h4 class="mb-3 font-semibold text-gray-900 dark:text-white">Codex 分组配置（使用 GPT 系列模型）</h4>
                        <p class="mb-2 text-sm text-gray-600 dark:text-gray-400">
                          将以下内容覆盖写入配置文件，把 <code class="rounded bg-gray-100 px-1 dark:bg-dark-800">sk-你的APIKey</code> 替换为你的密钥，
                          <code class="rounded bg-gray-100 px-1 dark:bg-dark-800">your-token</code> 替换为自定义 Gateway Token：
                        </p>
                        <GuideCodeBlock>
                          <code>{
  "agents": {
    "defaults": {
      "workspace": "~/.openclaw",
      "model": {
        "primary": "superwhv/gpt-5.4"
      },
      "compaction": {
        "mode": "safeguard",
        "reserveTokensFloor": 40000
      },
      "thinkingDefault": "high",
      "timeoutSeconds": 900,
      "maxConcurrent": 4,
      "subagents": {
        "maxConcurrent": 8
      }
    }
  },
  "models": {
    "mode": "merge",
    "providers": {
      "superwhv": {
        "apiKey": "sk-你的APIKey",
        "baseUrl": "https://api.superwhv.me/v1",
        "api": "openai-responses",
        "models": [
          {
            "contextWindow": 400000,
            "cost": {
              "cacheRead": 0.175,
              "cacheWrite": 0,
              "input": 1.75,
              "output": 14
            },
            "id": "gpt-5.4",
            "maxTokens": 128000,
            "name": "gpt-5.4",
            "reasoning": true
          },
          {
            "contextWindow": 272000,
            "cost": {
              "cacheRead": 0.25,
              "cacheWrite": 0,
              "input": 2.5,
              "output": 15
            },
            "id": "gpt-5.3-codex",
            "maxTokens": 128000,
            "name": "gpt-5.3-codex",
            "reasoning": true
          }
        ]
      }
    }
  },
  "messages": {
    "ackReactionScope": "group-mentions"
  },
  "commands": {
    "native": "auto",
    "nativeSkills": "auto"
  },
  "gateway": {
    "port": 18789,
    "mode": "local",
    "bind": "lan",
    "auth": {
      "mode": "token",
      "token": "your-token"
    },
    "tailscale": {
      "mode": "off",
      "resetOnExit": false
    }
  },
  "skills": {
    "install": {
      "nodeManager": "npm"
    }
  }
}</code>
                        </GuideCodeBlock>
                      </div>

                      <div class="rounded-lg border border-amber-200 bg-amber-50 p-4 dark:border-amber-800 dark:bg-amber-950/30">
                        <p class="text-sm text-amber-900 dark:text-amber-100">
                          💡 `sk-你的APIKey` 请替换为你在后台申请的 API Key！
                          <br />
                          💡 `your-token` 请设置为一个强随机字符串作为 Gateway Token！
                        </p>
                      </div>

                      <div>
                        <h4 class="mb-3 font-semibold text-gray-900 dark:text-white">验证配置</h4>
                        <GuideCodeBlock>
                          <code># 检查配置问题
openclaw doctor

# 查看已配置的模型（Auth 列显示 yes 说明配置成功）
openclaw models list</code>
                        </GuideCodeBlock>
                      </div>

                      <div>
                        <h4 class="mb-3 font-semibold text-gray-900 dark:text-white">启动使用</h4>
                        <GuideCodeBlock>
                          <code># 安装并启动后台 Gateway 服务
openclaw gateway install
openclaw gateway start

# 或前台运行（可看日志）
openclaw gateway --verbose

# 启动 TUI 交互界面（Windows PowerShell）
$env:OPENCLAW_GATEWAY_TOKEN="your-token"
openclaw tui

# macOS / Linux
OPENCLAW_GATEWAY_TOKEN="your-token" openclaw tui

# 命令行直接对话
openclaw agent --agent main --message "Hello"</code>
                        </GuideCodeBlock>
                      </div>

                      <div>
                        <h4 class="mb-3 font-semibold text-gray-900 dark:text-white">常用命令速查</h4>
                        <p class="mb-2 text-sm font-medium text-gray-900 dark:text-white">Gateway 管理：</p>
                        <ul class="mb-4 list-disc space-y-1 pl-5 text-sm text-gray-600 dark:text-gray-400">
                          <li><code class="rounded bg-gray-100 px-1 dark:bg-dark-800">openclaw gateway</code> 前台启动 Gateway</li>
                          <li><code class="rounded bg-gray-100 px-1 dark:bg-dark-800">openclaw gateway --verbose</code> 前台启动并显示详细日志</li>
                          <li><code class="rounded bg-gray-100 px-1 dark:bg-dark-800">openclaw gateway install</code> 安装后台服务</li>
                          <li><code class="rounded bg-gray-100 px-1 dark:bg-dark-800">openclaw gateway start</code> 启动后台服务</li>
                          <li><code class="rounded bg-gray-100 px-1 dark:bg-dark-800">openclaw gateway stop</code> 停止后台服务</li>
                          <li><code class="rounded bg-gray-100 px-1 dark:bg-dark-800">openclaw gateway restart</code> 重启 Gateway</li>
                          <li><code class="rounded bg-gray-100 px-1 dark:bg-dark-800">openclaw gateway status</code> 查看运行状态</li>
                        </ul>
                        <p class="mb-2 text-sm font-medium text-gray-900 dark:text-white">诊断与模型：</p>
                        <ul class="mb-4 list-disc space-y-1 pl-5 text-sm text-gray-600 dark:text-gray-400">
                          <li><code class="rounded bg-gray-100 px-1 dark:bg-dark-800">openclaw doctor</code> 检查配置问题</li>
                          <li><code class="rounded bg-gray-100 px-1 dark:bg-dark-800">openclaw doctor --fix</code> 自动修复配置问题</li>
                          <li><code class="rounded bg-gray-100 px-1 dark:bg-dark-800">openclaw models list</code> 查看可用模型</li>
                          <li><code class="rounded bg-gray-100 px-1 dark:bg-dark-800">openclaw models set superwhv/claude-opus-4-6</code> 切换默认模型</li>
                          <li><code class="rounded bg-gray-100 px-1 dark:bg-dark-800">openclaw logs --follow</code> 实时查看日志</li>
                        </ul>
                        <p class="mb-2 text-sm font-medium text-gray-900 dark:text-white">系统管理：</p>
                        <ul class="mb-4 list-disc space-y-1 pl-5 text-sm text-gray-600 dark:text-gray-400">
                          <li><code class="rounded bg-gray-100 px-1 dark:bg-dark-800">openclaw update</code> 更新 OpenClaw</li>
                          <li><code class="rounded bg-gray-100 px-1 dark:bg-dark-800">openclaw dashboard</code> 打开 Web 管理面板</li>
                          <li><code class="rounded bg-gray-100 px-1 dark:bg-dark-800">openclaw reset</code> 重置本地配置</li>
                          <li><code class="rounded bg-gray-100 px-1 dark:bg-dark-800">openclaw uninstall</code> 完全卸载</li>
                        </ul>
                      </div>

                      <div>
                        <p class="mb-3 font-semibold text-gray-900 dark:text-white">常见问题</p>
                        <div class="overflow-x-auto">
                          <table class="w-full text-left text-sm">
                            <thead class="border-b border-gray-200 bg-gray-50 dark:border-dark-700 dark:bg-dark-800">
                              <tr>
                                <th class="px-4 py-3 font-semibold text-gray-900 dark:text-white">错误信息</th>
                                <th class="px-4 py-3 font-semibold text-gray-900 dark:text-white">解决方案</th>
                              </tr>
                            </thead>
                            <tbody class="divide-y divide-gray-200 dark:divide-dark-700">
                              <tr><td class="px-4 py-3 font-mono text-xs text-red-600 dark:text-red-400">No API key found</td><td class="px-4 py-3 text-gray-700 dark:text-gray-300">检查 apiKey 是否正确填写</td></tr>
                              <tr><td class="px-4 py-3 font-mono text-xs text-red-600 dark:text-red-400">Gateway start blocked</td><td class="px-4 py-3 text-gray-700 dark:text-gray-300">在 gateway 中加上 "mode": "local"</td></tr>
                              <tr><td class="px-4 py-3 font-mono text-xs text-red-600 dark:text-red-400">non-loopback Control UI requires allowedOrigins</td><td class="px-4 py-3 text-gray-700 dark:text-gray-300">补上 gateway.controlUi.allowedOrigins 字段</td></tr>
                              <tr><td class="px-4 py-3 font-mono text-xs text-red-600 dark:text-red-400">Connection refused</td><td class="px-4 py-3 text-gray-700 dark:text-gray-300">检查 baseUrl 和网络连接</td></tr>
                              <tr><td class="px-4 py-3 font-mono text-xs text-red-600 dark:text-red-400">401 Unauthorized</td><td class="px-4 py-3 text-gray-700 dark:text-gray-300">检查 API Key 是否正确</td></tr>
                              <tr><td class="px-4 py-3 font-mono text-xs text-red-600 dark:text-red-400">404 Not Found</td><td class="px-4 py-3 text-gray-700 dark:text-gray-300">检查 baseUrl 路径是否正确</td></tr>
                            </tbody>
                          </table>
                        </div>
                      </div>

                      <div>
                        <p class="text-sm text-gray-600 dark:text-gray-400">
                          官方文档：<a href="https://docs.openclaw.ai" target="_blank" rel="noopener" class="text-blue-600 underline dark:text-blue-400">https://docs.openclaw.ai</a>
                        </p>
                      </div>
                    </div>

                    <!-- VS Code / Cursor -->
                    <div v-show="activeClient === 'vscode'" class="space-y-4">
                      <div>
                        <h4 class="mb-3 font-semibold text-gray-900 dark:text-white">VSCode &amp; Cursor 配置</h4>

                        <div class="space-y-4">
                          <div>
                            <p class="mb-2 font-medium text-gray-900 dark:text-white">安装插件</p>
                            <p class="mb-2 text-sm text-gray-600 dark:text-gray-400">
                              在 VSCode 或 Cursor 的插件市场搜索并安装：
                            </p>
                            <ul class="list-disc space-y-1 pl-5 text-sm text-gray-600 dark:text-gray-400">
                              <li>Claude Code：搜索 <code class="rounded bg-gray-100 px-1 dark:bg-dark-800">claude</code> 或 <code class="rounded bg-gray-100 px-1 dark:bg-dark-800">Anthropic</code>（需先保证 Claude Code 可用）</li>
                              <li>Codex：搜索 <code class="rounded bg-gray-100 px-1 dark:bg-dark-800">Codex - OpenAI's coding agent</code>（需先保证 Codex 可用）</li>
                            </ul>
                          </div>

                          <div>
                            <p class="mb-2 font-medium text-gray-900 dark:text-white">配置</p>
                            <p class="mb-2 text-sm text-gray-600 dark:text-gray-400">
                              插件会自动使用系统环境变量中配置的 <code class="rounded bg-gray-100 px-1 dark:bg-dark-800">ANTHROPIC_BASE_URL</code> 和
                              <code class="rounded bg-gray-100 px-1 dark:bg-dark-800">ANTHROPIC_AUTH_TOKEN</code>。
                            </p>
                            <p class="text-sm text-gray-600 dark:text-gray-400">确保已正确设置环境变量后，重启编辑器即可使用。</p>
                          </div>
                        </div>
                      </div>
                    </div>

                    <!-- 通用提示 -->
                    <div class="mt-6 rounded-lg border border-sky-200 bg-sky-50 p-4 dark:border-sky-800 dark:bg-sky-950/30">
                      <p class="text-sm text-sky-900 dark:text-sky-100">
                        💡 以上为完整安装指南。若仍有疑问，可查看后台"使用记录"或联系技术支持。
                      </p>
                    </div>
                  </div>
                </div>
              </div>
            </section>

            <!-- Error Codes Reference -->
            <section id="errors" class="rounded-lg border border-gray-200 bg-white p-6 dark:border-dark-700 dark:bg-dark-900">
              <h2 class="mb-6 text-xl font-bold text-gray-900 dark:text-white">
                {{ t('guide.errorCodes') }}
              </h2>

              <div class="overflow-x-auto">
                <table class="w-full text-left text-sm">
                  <thead class="border-b border-gray-200 bg-gray-50 dark:border-dark-700 dark:bg-dark-800">
                    <tr>
                      <th class="px-4 py-3 font-medium text-gray-900 dark:text-white">{{ t('guide.errorCode') }}</th>
                      <th class="px-4 py-3 font-medium text-gray-900 dark:text-white">{{ t('guide.errorDescription') }}</th>
                    </tr>
                  </thead>
                  <tbody class="divide-y divide-gray-100 dark:divide-dark-700">
                    <tr v-for="error in errorCodes" :key="error.code">
                      <td class="px-4 py-3 font-mono text-red-600 dark:text-red-400">{{ error.code }}</td>
                      <td class="px-4 py-3 text-gray-700 dark:text-gray-300">{{ error.description }}</td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </section>

            <!-- Support -->
            <section id="support" class="rounded-lg border border-gray-200 bg-white p-6 dark:border-dark-700 dark:bg-dark-900">
              <h2 class="mb-4 text-xl font-bold text-gray-900 dark:text-white">
                {{ t('guide.technicalSupport') }}
              </h2>
              <p class="text-sm text-gray-600 dark:text-gray-400">
                {{ t('guide.supportHint') }}
              </p>
            </section>
          </main>
        </div>
      </div>
    </div>
  </AppLayout>
</template>

<script setup lang="ts">
import { ref, onMounted, onBeforeUnmount } from 'vue'
import { useI18n } from 'vue-i18n'
import AppLayout from '@/components/layout/AppLayout.vue'
import Icon from '@/components/icons/Icon.vue'
import GuideCodeBlock from '@/components/Guide/CodeBlock.vue'
import { getPublicSettings } from '@/api/auth'

const { t } = useI18n()

const activeSection = ref('purchase')
const activeClient = ref('claude-code')
const loadingSettings = ref(false)
const afterSalesQRCode = ref('')

const sections = [
  { id: 'purchase', title: t('guide.purchaseChannels') },
  { id: 'notice', title: t('guide.importantNotice') },
  { id: 'models', title: t('guide.modelList') },
  { id: 'installation', title: t('guide.installationGuide') },
  { id: 'errors', title: t('guide.errorCodes') },
  { id: 'support', title: t('guide.technicalSupport') },
]

const clients = [
  { id: 'cc-switch', name: 'CC-SWITCH' },
  { id: 'claude-code', name: 'Claude Code' },
  { id: 'claude-app', name: 'Claude APP' },
  { id: 'codex', name: 'Codex TUI' },
  { id: 'codex-app', name: 'Codex APP' },
  { id: 'codex-remote', name: 'Codex Remote' },
  { id: 'gemini', name: 'Gemini CLI' },
  { id: 'openclaw', name: 'OpenClaw' },
  { id: 'vscode', name: 'VS Code / Cursor' },
]

const models = [
  { name: 'claude-opus-4-6', description: '旗舰 Opus 模型，最强推理', context: '200K tokens' },
  { name: 'claude-opus-4-7', description: '上一代旗舰', context: '200K tokens' },
  { name: 'claude-sonnet-4-6', description: '平衡性能', context: '200K tokens' },
  { name: 'claude-sonnet-4-5-20250929', description: '平衡性能和速度', context: '200K tokens' },
  { name: 'claude-opus-4-5-20251101', description: 'Opus 4.5 系列，高性价比旗舰', context: '200K tokens' },
  { name: 'claude-haiku-4-5-20251001', description: '快速响应，轻量任务', context: '200K tokens' },
  { name: 'gpt-5.3-codex', description: 'Codex 专用 GPT 系列', context: '400K tokens' },
  { name: 'gpt-5.4', description: 'GPT 通用模型', context: '400K tokens' },
  { name: 'gpt-5.4-mini', description: 'GPT 轻量快速模型', context: '400K tokens' },
  { name: 'gpt-5.5', description: 'GPT 旗舰模型', context: '400K tokens' },
  { name: 'gpt-image-2', description: '图像生成模型', context: '—' },
]

const errorCodes = [
  { code: 401, description: t('guide.error401') },
  { code: 403, description: t('guide.error403') },
  { code: 429, description: t('guide.error429') },
  { code: 500, description: t('guide.error500') },
  { code: 502, description: t('guide.error502') },
  { code: 503, description: t('guide.error503') },
]

function scrollToSection(sectionId: string) {
  const element = document.getElementById(sectionId)
  if (element) {
    const offset = 80
    const elementPosition = element.getBoundingClientRect().top + window.scrollY
    window.scrollTo({
      top: elementPosition - offset,
      behavior: 'smooth',
    })
    activeSection.value = sectionId
  }
}

function handleScroll() {
  const scrollPosition = window.scrollY + 120
  for (const section of sections) {
    const element = document.getElementById(section.id)
    if (element) {
      const { offsetTop, offsetHeight } = element
      if (scrollPosition >= offsetTop && scrollPosition < offsetTop + offsetHeight) {
        activeSection.value = section.id
        break
      }
    }
  }
}

async function loadSettings() {
  loadingSettings.value = true
  try {
    // 从公开设置加载售后群二维码；未配置时回退到内置默认图
    const settings = await getPublicSettings()
    afterSalesQRCode.value = settings.after_sales_qr || ''
  } catch (error) {
    console.error('Failed to load settings:', error)
  } finally {
    loadingSettings.value = false
  }
}

onMounted(() => {
  window.addEventListener('scroll', handleScroll)
  void loadSettings()
})

onBeforeUnmount(() => {
  window.removeEventListener('scroll', handleScroll)
})
</script>
