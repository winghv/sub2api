<template>
  <!-- Custom Home Content: Full Page Mode -->
  <div v-if="homeContent" class="min-h-screen">
    <iframe
      v-if="isHomeContentUrl"
      :src="homeContent.trim()"
      class="h-screen w-full border-0"
      allowfullscreen
    ></iframe>
    <!-- homeContent is an admin-only setting. -->
    <div v-else v-html="homeContent"></div>
  </div>

  <!-- Default Home Page -->
  <div v-else class="cyber-home">
    <div class="cyber-backdrop" aria-hidden="true">
      <span class="cyber-backdrop__beam cyber-backdrop__beam--cyan"></span>
      <span class="cyber-backdrop__beam cyber-backdrop__beam--violet"></span>
      <span class="cyber-backdrop__noise"></span>
    </div>

    <header class="cyber-header">
      <nav class="cyber-nav" :aria-label="siteName">
        <router-link to="/" class="cyber-brand" :aria-label="siteName">
          <span class="cyber-brand__mark">
            <img :src="siteLogo || DEFAULT_SITE_LOGO" :alt="siteName" />
          </span>
          <span class="cyber-brand__copy">
            <span class="cyber-brand__name">{{ siteName }}</span>
            <span class="cyber-brand__mode">{{ t('home.cyber.brandMode') }}</span>
          </span>
        </router-link>

        <div class="cyber-actions">
          <a
            v-if="docUrl"
            :href="docUrl"
            target="_blank"
            rel="noopener noreferrer"
            class="cyber-icon-button cyber-icon-button--wide"
            :title="t('home.viewDocs')"
          >
            <Icon name="book" size="sm" />
            <span>{{ t('home.docs') }}</span>
          </a>

          <div class="cyber-locale-control">
            <LocaleSwitcher />
          </div>

          <button
            type="button"
            class="cyber-icon-button"
            @click="toggleTheme"
            :title="isDark ? t('home.switchToLight') : t('home.switchToDark')"
          >
            <Icon v-if="isDark" name="sun" size="sm" />
            <Icon v-else name="moon" size="sm" />
          </button>

          <router-link
            :to="isAuthenticated ? dashboardPath : '/login'"
            class="cyber-login-link"
          >
            <span v-if="isAuthenticated" class="cyber-user-initial">{{ userInitial }}</span>
            <Icon v-else name="login" size="sm" />
            <span>{{ isAuthenticated ? t('home.dashboard') : t('home.login') }}</span>
          </router-link>
        </div>
      </nav>
    </header>

    <main class="cyber-main">
      <section class="cyber-hero" aria-labelledby="home-heading">
        <div class="cyber-hero__copy">
          <div class="cyber-eyebrow">
            <span class="cyber-eyebrow__pulse" aria-hidden="true"></span>
            <span>{{ t('home.cyber.eyebrow') }}</span>
          </div>

          <h1 id="home-heading" class="cyber-title">
            <span>{{ siteName }}</span>
            <strong>{{ siteSubtitle }}</strong>
          </h1>

          <p class="cyber-description">{{ t('home.cyber.description') }}</p>

          <div class="cyber-cta-row">
            <router-link
              :to="isAuthenticated ? dashboardPath : '/login'"
              class="cyber-button cyber-button--primary"
            >
              <Icon name="arrowRight" size="sm" :stroke-width="2" />
              <span>{{ isAuthenticated ? t('home.goToDashboard') : t('home.getStarted') }}</span>
            </router-link>

            <a
              v-if="docUrl"
              :href="docUrl"
              target="_blank"
              rel="noopener noreferrer"
              class="cyber-button cyber-button--secondary"
            >
              <Icon name="book" size="sm" />
              <span>{{ t('home.viewDocs') }}</span>
            </a>
          </div>

          <dl class="cyber-signal-list">
            <div v-for="signal in signals" :key="signal.label" class="cyber-signal">
              <dt>{{ signal.label }}</dt>
              <dd>{{ signal.value }}</dd>
            </div>
          </dl>
        </div>

        <div class="cyber-hero__visual" :aria-label="t('home.cyber.visualLabel')">
          <div class="gateway-panel">
            <div class="gateway-panel__header">
              <span>{{ t('home.cyber.networkTitle') }}</span>
              <strong>{{ t('home.cyber.statusOnline') }}</strong>
            </div>

            <div class="gateway-route-map" aria-hidden="true">
              <div class="gateway-node gateway-node--source">
                <span>Client</span>
                <small>Agent / App</small>
              </div>
              <div class="gateway-node gateway-node--core">
                <img :src="siteLogo || DEFAULT_SITE_LOGO" alt="" />
                <strong>{{ siteName }}</strong>
              </div>
              <div class="gateway-node gateway-node--anthropic">
                <span>Claude</span>
                <small>Messages</small>
              </div>
              <div class="gateway-node gateway-node--openai">
                <span>GPT</span>
                <small>Chat</small>
              </div>
              <span class="gateway-route gateway-route--source"></span>
              <span class="gateway-route gateway-route--anthropic"></span>
              <span class="gateway-route gateway-route--openai"></span>
            </div>
          </div>

          <div class="cyber-terminal">
            <div class="cyber-terminal__bar">
              <span class="cyber-terminal__dot"></span>
              <span>{{ t('home.cyber.terminalTitle') }}</span>
              <strong>200 OK</strong>
            </div>
            <div class="cyber-terminal__body">
              <p
                v-for="(line, index) in terminalLines"
                :key="line"
                :class="{ 'cyber-terminal__success': index === terminalLines.length - 1 }"
                :style="{ '--line': index }"
              >
                <span>&gt;</span> {{ line }}
              </p>
            </div>
          </div>
        </div>
      </section>

      <section class="cyber-capabilities" aria-labelledby="capabilities-heading">
        <div class="cyber-section-heading">
          <span>{{ t('home.cyber.capabilities.eyebrow') }}</span>
          <h2 id="capabilities-heading">{{ t('home.cyber.capabilities.title') }}</h2>
          <p>{{ t('home.cyber.capabilities.description') }}</p>
        </div>

        <div class="cyber-capability-grid">
          <article
            v-for="capability in capabilityCards"
            :key="capability.title"
            class="cyber-capability-card"
          >
            <span class="cyber-card-icon" aria-hidden="true">
              <Icon :name="capability.icon" size="md" />
            </span>
            <h3>{{ capability.title }}</h3>
            <p>{{ capability.description }}</p>
            <code>{{ capability.detail }}</code>
          </article>
        </div>
      </section>

      <section class="cyber-advantages" aria-labelledby="advantages-heading">
        <div class="cyber-section-heading cyber-section-heading--center">
          <span>{{ t('home.cyber.advantages.eyebrow') }}</span>
          <h2 id="advantages-heading">{{ t('home.cyber.advantages.title') }}</h2>
          <p>{{ t('home.cyber.advantages.description') }}</p>
        </div>

        <div class="cyber-advantage-grid">
          <article
            v-for="(advantage, index) in advantages"
            :key="advantage.title"
            class="cyber-advantage-card"
            :style="{ '--delay': `${index * 90}ms` }"
          >
            <div class="cyber-advantage-card__index">
              {{ String(index + 1).padStart(2, '0') }}
            </div>
            <div class="cyber-advantage-card__body">
              <h3>{{ advantage.title }}</h3>
              <p>{{ advantage.description }}</p>
            </div>
            <span class="cyber-advantage-card__meter" aria-hidden="true"></span>
          </article>
        </div>
      </section>

      <section class="cyber-workflow" aria-labelledby="workflow-heading">
        <div class="cyber-section-heading">
          <span>{{ t('home.cyber.workflow.eyebrow') }}</span>
          <h2 id="workflow-heading">{{ t('home.cyber.workflow.title') }}</h2>
          <p>{{ t('home.cyber.workflow.description') }}</p>
        </div>

        <ol class="cyber-workflow-list">
          <li
            v-for="(step, index) in workflowSteps"
            :key="step.title"
            class="cyber-workflow-step"
          >
            <span class="cyber-workflow-step__number">{{ String(index + 1).padStart(2, '0') }}</span>
            <div>
              <h3>{{ step.title }}</h3>
              <p>{{ step.description }}</p>
            </div>
          </li>
        </ol>
      </section>

      <section class="cyber-final-cta" aria-labelledby="final-cta-heading">
        <div>
          <span class="cyber-final-cta__eyebrow">{{ t('home.cyber.finalCta.eyebrow') }}</span>
          <h2 id="final-cta-heading">{{ t('home.cyber.finalCta.title') }}</h2>
          <p>{{ t('home.cyber.finalCta.description') }}</p>
        </div>

        <div class="cyber-cta-row cyber-cta-row--compact">
          <router-link
            :to="isAuthenticated ? dashboardPath : '/login'"
            class="cyber-button cyber-button--primary"
          >
            <Icon name="arrowRight" size="sm" :stroke-width="2" />
            <span>{{ isAuthenticated ? t('home.goToDashboard') : t('home.getStarted') }}</span>
          </router-link>

          <a
            v-if="docUrl"
            :href="docUrl"
            target="_blank"
            rel="noopener noreferrer"
            class="cyber-button cyber-button--secondary"
          >
            <Icon name="book" size="sm" />
            <span>{{ t('home.viewDocs') }}</span>
          </a>
        </div>
      </section>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useI18n } from 'vue-i18n'
import { useAuthStore, useAppStore } from '@/stores'
import LocaleSwitcher from '@/components/common/LocaleSwitcher.vue'
import Icon from '@/components/icons/Icon.vue'
import { DEFAULT_SITE_LOGO, DEFAULT_SITE_NAME, isDefaultSiteSubtitle } from '@/utils/brand'
import { sanitizeUrl } from '@/utils/url'

type IconName = InstanceType<typeof Icon>['$props']['name']

const { t } = useI18n()

const authStore = useAuthStore()
const appStore = useAppStore()

// Site settings - dev 品牌默认值 + upstream sanitizeUrl 安全包装
const siteName = computed(() => appStore.cachedPublicSettings?.site_name || appStore.siteName || DEFAULT_SITE_NAME)
const siteLogo = computed(() => sanitizeUrl(appStore.cachedPublicSettings?.site_logo || appStore.siteLogo || '', { allowRelative: true, allowDataUrl: true }))
const localizedDefaultSubtitle = computed(() => t('home.cyber.subtitle'))
const siteSubtitle = computed(() => {
  const configuredSubtitle = appStore.cachedPublicSettings?.site_subtitle?.trim()
  if (isDefaultSiteSubtitle(configuredSubtitle)) {
    return localizedDefaultSubtitle.value
  }
  return configuredSubtitle || localizedDefaultSubtitle.value
})
const docUrl = computed(() => sanitizeUrl(appStore.cachedPublicSettings?.doc_url || appStore.docUrl || ''))
const homeContent = computed(() => appStore.cachedPublicSettings?.home_content || '')

const terminalLines = computed(() => [
  t('home.cyber.terminalLines.chat'),
  t('home.cyber.terminalLines.messages'),
  t('home.cyber.terminalLines.route'),
  t('home.cyber.terminalLines.quota'),
  t('home.cyber.terminalLines.ready')
])

const isHomeContentUrl = computed(() => {
  const content = homeContent.value.trim()
  return content.startsWith('http://') || content.startsWith('https://')
})

const isDark = ref(document.documentElement.classList.contains('dark'))

const isAuthenticated = computed(() => authStore.isAuthenticated)
const isAdmin = computed(() => authStore.isAdmin)
const dashboardPath = computed(() => (isAdmin.value ? '/admin/dashboard' : '/dashboard'))
const userInitial = computed(() => {
  const user = authStore.user
  if (!user?.email) return ''
  return user.email.charAt(0).toUpperCase()
})

const signals = computed(() => [
  { label: t('home.cyber.signals.gateway'), value: t('home.cyber.signals.gatewayValue') },
  { label: t('home.cyber.signals.models'), value: t('home.cyber.signals.modelsValue') },
  { label: t('home.cyber.signals.billing'), value: t('home.cyber.signals.billingValue') }
])

const capabilityCards = computed<Array<{
  title: string
  description: string
  detail: string
  icon: IconName
}>>(() => [
  {
    title: t('home.cyber.capabilities.items.openai.title'),
    description: t('home.cyber.capabilities.items.openai.description'),
    detail: '/v1/chat/completions',
    icon: 'terminal'
  },
  {
    title: t('home.cyber.capabilities.items.anthropic.title'),
    description: t('home.cyber.capabilities.items.anthropic.description'),
    detail: '/v1/messages',
    icon: 'chat'
  },
  {
    title: t('home.cyber.capabilities.items.routing.title'),
    description: t('home.cyber.capabilities.items.routing.description'),
    detail: 'GPT / Claude',
    icon: 'swap'
  }
])

const advantages = computed(() => [
  {
    title: t('home.cyber.advantages.items.protocol.title'),
    description: t('home.cyber.advantages.items.protocol.description')
  },
  {
    title: t('home.cyber.advantages.items.routing.title'),
    description: t('home.cyber.advantages.items.routing.description')
  },
  {
    title: t('home.cyber.advantages.items.cost.title'),
    description: t('home.cyber.advantages.items.cost.description')
  },
  {
    title: t('home.cyber.advantages.items.observability.title'),
    description: t('home.cyber.advantages.items.observability.description')
  }
])

const workflowSteps = computed(() => [
  {
    title: t('home.cyber.workflow.steps.key.title'),
    description: t('home.cyber.workflow.steps.key.description')
  },
  {
    title: t('home.cyber.workflow.steps.route.title'),
    description: t('home.cyber.workflow.steps.route.description')
  },
  {
    title: t('home.cyber.workflow.steps.quota.title'),
    description: t('home.cyber.workflow.steps.quota.description')
  },
  {
    title: t('home.cyber.workflow.steps.observe.title'),
    description: t('home.cyber.workflow.steps.observe.description')
  }
])

function toggleTheme() {
  isDark.value = !isDark.value
  document.documentElement.classList.toggle('dark', isDark.value)
  localStorage.setItem('theme', isDark.value ? 'dark' : 'light')
}

function initTheme() {
  const savedTheme = localStorage.getItem('theme')
  if (
    savedTheme === 'dark' ||
    (!savedTheme && window.matchMedia('(prefers-color-scheme: dark)').matches)
  ) {
    isDark.value = true
    document.documentElement.classList.add('dark')
  }
}

onMounted(() => {
  initTheme()
  authStore.checkAuth()

  if (!appStore.publicSettingsLoaded) {
    appStore.fetchPublicSettings()
  }
})
</script>

<style scoped>
.cyber-home {
  position: relative;
  min-height: 100vh;
  overflow: hidden;
  background:
    radial-gradient(circle at 74% 18%, rgba(0, 229, 255, 0.14), transparent 30%),
    radial-gradient(circle at 12% 64%, rgba(57, 255, 20, 0.08), transparent 28%),
    linear-gradient(180deg, rgba(5, 7, 18, 0.94) 0%, rgba(7, 10, 23, 0.98) 46%, #050711 100%),
    #050711;
  color: #f5fbff;
  color-scheme: dark;
  font-family:
    Inter, ui-sans-serif, system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif;
  letter-spacing: 0;
}

.cyber-home::before,
.cyber-home::after {
  position: absolute;
  inset: 0;
  pointer-events: none;
  content: "";
}

.cyber-home::before {
  z-index: 0;
  background:
    repeating-linear-gradient(0deg, rgba(74, 222, 255, 0.055) 0 1px, transparent 1px 28px),
    repeating-linear-gradient(90deg, rgba(74, 222, 255, 0.04) 0 1px, transparent 1px 28px);
  mask-image: linear-gradient(to bottom, rgba(0, 0, 0, 0.82), transparent 92%);
}

.cyber-home::after {
  z-index: 1;
  opacity: 0.12;
  background: repeating-linear-gradient(
    180deg,
    transparent 0,
    transparent 6px,
    rgba(255, 255, 255, 0.08) 7px
  );
  mix-blend-mode: screen;
}

.cyber-backdrop,
.cyber-header,
.cyber-main {
  position: relative;
  z-index: 2;
}

.cyber-header {
  z-index: 10;
}

.cyber-backdrop {
  position: absolute;
  inset: 0;
  z-index: 1;
  pointer-events: none;
  overflow: hidden;
}

.cyber-backdrop__beam {
  position: absolute;
  width: 46vw;
  height: 170px;
  transform: rotate(-16deg);
  border: 1px solid rgba(255, 255, 255, 0.08);
  opacity: 0.42;
  clip-path: polygon(8% 0, 100% 0, 92% 100%, 0 100%);
}

.cyber-backdrop__beam--cyan {
  top: 14%;
  right: -24%;
  background: linear-gradient(90deg, transparent, rgba(0, 229, 255, 0.18), transparent);
}

.cyber-backdrop__beam--violet {
  bottom: 10%;
  left: -32%;
  background: linear-gradient(90deg, transparent, rgba(178, 91, 255, 0.14), transparent);
}

.cyber-backdrop__noise {
  position: absolute;
  inset: 0;
  opacity: 0.07;
  background-image:
    linear-gradient(90deg, rgba(57, 255, 20, 0.12) 1px, transparent 1px),
    linear-gradient(rgba(57, 255, 20, 0.08) 1px, transparent 1px);
  background-size: 7px 7px;
}

.cyber-header {
  padding: 18px 24px 0;
}

.cyber-nav {
  display: flex;
  width: min(1180px, 100%);
  max-width: 100%;
  min-height: 60px;
  align-items: center;
  justify-content: space-between;
  gap: 18px;
  margin: 0 auto;
  padding: 10px;
  border: 1px solid rgba(0, 229, 255, 0.22);
  border-radius: 8px;
  background: rgba(5, 10, 22, 0.78);
  box-shadow:
    0 0 0 1px rgba(178, 91, 255, 0.07),
    0 16px 48px rgba(0, 0, 0, 0.26);
  backdrop-filter: blur(16px);
}

.cyber-brand {
  display: inline-flex;
  min-width: 0;
  align-items: center;
  gap: 12px;
  color: inherit;
  text-decoration: none;
}

.cyber-brand__mark {
  display: grid;
  width: 46px;
  height: 46px;
  flex: 0 0 auto;
  place-items: center;
  border: 1px solid rgba(0, 229, 255, 0.58);
  border-radius: 6px;
  background:
    linear-gradient(135deg, rgba(0, 229, 255, 0.2), rgba(178, 91, 255, 0.14)),
    #07111f;
  box-shadow: 0 0 24px rgba(0, 229, 255, 0.18);
}

.cyber-brand__mark img {
  width: 36px;
  height: 36px;
  object-fit: contain;
  image-rendering: pixelated;
}

.cyber-brand__copy {
  display: grid;
  min-width: 0;
  gap: 1px;
}

.cyber-brand__name,
.cyber-brand__mode,
.cyber-eyebrow,
.cyber-terminal,
.cyber-signal-list,
.cyber-card-icon,
.cyber-section-heading span,
.cyber-advantage-card__index,
.cyber-workflow-step__number,
.cyber-final-cta__eyebrow {
  font-family: "SFMono-Regular", "Cascadia Code", "Roboto Mono", Consolas, monospace;
}

.cyber-brand__name {
  overflow: hidden;
  font-size: 15px;
  font-weight: 800;
  line-height: 1.2;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.cyber-brand__mode {
  overflow: hidden;
  color: #00e5ff;
  font-size: 11px;
  line-height: 1.2;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.cyber-actions {
  display: flex;
  flex: 0 0 auto;
  align-items: center;
  justify-content: flex-end;
  gap: 8px;
}

.cyber-icon-button,
.cyber-login-link {
  display: inline-flex;
  min-height: 38px;
  align-items: center;
  justify-content: center;
  gap: 8px;
  border: 1px solid rgba(0, 229, 255, 0.28);
  border-radius: 6px;
  background: rgba(8, 16, 34, 0.86);
  color: #d7f7ff;
  font-size: 13px;
  font-weight: 700;
  line-height: 1;
  text-decoration: none;
  transition:
    border-color 160ms ease,
    background 160ms ease,
    color 160ms ease,
    transform 160ms ease;
}

.cyber-icon-button {
  width: 38px;
  padding: 0;
}

.cyber-icon-button--wide {
  width: auto;
  padding: 0 12px;
}

.cyber-login-link {
  padding: 0 14px;
  border-color: rgba(57, 255, 20, 0.5);
  color: #e8ffe6;
  box-shadow: inset 0 0 18px rgba(57, 255, 20, 0.09);
}

.cyber-icon-button:hover,
.cyber-login-link:hover {
  border-color: #00e5ff;
  background: rgba(0, 229, 255, 0.12);
  color: #ffffff;
  transform: translateY(-1px);
}

.cyber-icon-button:focus-visible,
.cyber-login-link:focus-visible,
.cyber-button:focus-visible {
  outline: 2px solid #39ff14;
  outline-offset: 3px;
}

.cyber-user-initial {
  display: grid;
  width: 20px;
  height: 20px;
  place-items: center;
  border: 1px solid rgba(57, 255, 20, 0.7);
  border-radius: 4px;
  color: #39ff14;
  font-family: "SFMono-Regular", Consolas, monospace;
  font-size: 11px;
  font-weight: 800;
}

.cyber-locale-control :deep(button) {
  min-height: 38px;
  border: 1px solid rgba(0, 229, 255, 0.28);
  border-radius: 6px;
  background: rgba(8, 16, 34, 0.86);
  color: #d7f7ff;
}

.cyber-locale-control :deep(button:hover) {
  background: rgba(0, 229, 255, 0.12);
  color: #ffffff;
}

.cyber-locale-control :deep(.absolute) {
  border-color: rgba(0, 229, 255, 0.3);
  border-radius: 6px;
  background: rgba(7, 13, 28, 0.98);
  color: #e8fbff;
}

.cyber-main {
  width: min(1180px, calc(100% - 48px));
  max-width: 100%;
  margin: 0 auto;
  padding: 54px 0 48px;
}

.cyber-hero {
  display: grid;
  grid-template-columns: minmax(0, 1fr) minmax(370px, 0.86fr);
  min-height: calc(100vh - 214px);
  align-items: center;
  gap: 54px;
}

.cyber-hero__copy {
  min-width: 0;
  max-width: 100%;
}

.cyber-eyebrow {
  display: inline-flex;
  max-width: 100%;
  align-items: center;
  gap: 10px;
  margin-bottom: 22px;
  color: #39ff14;
  font-size: 13px;
  font-weight: 800;
  overflow-wrap: break-word;
}

.cyber-eyebrow__pulse {
  width: 10px;
  height: 10px;
  border: 2px solid #39ff14;
  background: #07111f;
  box-shadow: 0 0 16px rgba(57, 255, 20, 0.56);
}

.cyber-title {
  display: grid;
  max-width: 760px;
  gap: 12px;
  margin: 0;
  color: #ffffff;
  font-size: 76px;
  font-weight: 900;
  line-height: 0.98;
  overflow-wrap: break-word;
  text-shadow: 0 0 24px rgba(0, 229, 255, 0.24);
}

.cyber-title > span {
  width: fit-content;
  max-width: 100%;
  background: linear-gradient(
    90deg,
    #ffffff 0%,
    #00e5ff 23%,
    #39ff14 47%,
    #b25bff 72%,
    #ffffff 100%
  );
  background-size: 240% 100%;
  -webkit-background-clip: text;
  background-clip: text;
  color: transparent;
  -webkit-text-fill-color: transparent;
  text-shadow: 0 0 26px rgba(0, 229, 255, 0.24);
}

.cyber-title strong {
  color: #dffbff;
  font-size: 30px;
  font-weight: 800;
  line-height: 1.25;
}

.cyber-description {
  max-width: 680px;
  margin: 18px 0 0;
  color: #9db4c7;
  font-size: 17px;
  line-height: 1.8;
  overflow-wrap: break-word;
}

.cyber-cta-row {
  display: flex;
  flex-wrap: wrap;
  gap: 12px;
  margin-top: 34px;
}

.cyber-cta-row--compact {
  margin-top: 0;
}

.cyber-button {
  display: inline-flex;
  min-height: 50px;
  align-items: center;
  justify-content: center;
  gap: 10px;
  padding: 0 20px;
  border-radius: 6px;
  font-size: 15px;
  font-weight: 900;
  line-height: 1;
  text-decoration: none;
  transition:
    box-shadow 160ms ease,
    transform 160ms ease,
    background 160ms ease,
    border-color 160ms ease;
  clip-path: polygon(0 0, calc(100% - 12px) 0, 100% 12px, 100% 100%, 12px 100%, 0 calc(100% - 12px));
}

.cyber-button--primary {
  border: 1px solid rgba(57, 255, 20, 0.72);
  background: linear-gradient(135deg, #39ff14 0%, #00e5ff 100%);
  color: #031007;
  box-shadow: 0 0 30px rgba(57, 255, 20, 0.2);
}

.cyber-button--primary span,
.cyber-button--primary svg {
  color: inherit;
}

.cyber-button--secondary {
  border: 1px solid rgba(0, 229, 255, 0.45);
  background: rgba(7, 15, 32, 0.86);
  color: #d7f7ff;
}

.cyber-button:hover {
  transform: translateY(-2px);
}

.cyber-button--primary:hover {
  box-shadow: 0 0 38px rgba(0, 229, 255, 0.3);
}

.cyber-button--secondary:hover {
  border-color: rgba(178, 91, 255, 0.76);
  background: rgba(178, 91, 255, 0.12);
}

.cyber-signal-list {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: 10px;
  max-width: 680px;
  margin: 30px 0 0;
}

.cyber-signal {
  min-width: 0;
  padding: 12px;
  border: 1px solid rgba(0, 229, 255, 0.22);
  border-radius: 6px;
  background: rgba(8, 17, 36, 0.5);
}

.cyber-signal dt {
  color: #8096aa;
  font-size: 11px;
  line-height: 1.3;
}

.cyber-signal dd {
  margin: 5px 0 0;
  color: #f3fbff;
  font-size: 13px;
  font-weight: 800;
  line-height: 1.35;
  overflow-wrap: anywhere;
}

.cyber-hero__visual {
  display: grid;
  gap: 14px;
  min-width: 0;
  max-width: 100%;
}

.gateway-panel,
.cyber-terminal,
.cyber-capability-card,
.cyber-advantage-card,
.cyber-workflow-step,
.cyber-final-cta {
  max-width: 100%;
  border: 1px solid rgba(0, 229, 255, 0.24);
  border-radius: 8px;
  background:
    linear-gradient(135deg, rgba(0, 229, 255, 0.07), transparent 42%),
    rgba(5, 12, 27, 0.8);
  box-shadow:
    inset 0 0 0 1px rgba(255, 255, 255, 0.03),
    0 18px 50px rgba(0, 0, 0, 0.24);
}

.gateway-panel {
  padding: 14px;
}

.gateway-panel__header,
.cyber-terminal__bar {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
  color: #9db4c7;
  font-family: "SFMono-Regular", Consolas, monospace;
  font-size: 12px;
  line-height: 1.2;
}

.gateway-panel__header strong,
.cyber-terminal__bar strong {
  color: #39ff14;
  font-weight: 900;
}

.gateway-route-map {
  position: relative;
  height: 286px;
  margin-top: 14px;
  overflow: hidden;
  border: 1px solid rgba(178, 91, 255, 0.18);
  border-radius: 6px;
  background:
    linear-gradient(90deg, rgba(0, 229, 255, 0.05) 1px, transparent 1px),
    linear-gradient(rgba(0, 229, 255, 0.05) 1px, transparent 1px),
    #07101f;
  background-size: 20px 20px;
}

.gateway-node {
  position: absolute;
  z-index: 2;
  display: grid;
  min-width: 88px;
  min-height: 58px;
  place-items: center;
  padding: 8px;
  border: 1px solid rgba(0, 229, 255, 0.62);
  border-radius: 6px;
  background: rgba(8, 20, 38, 0.96);
  color: #f8fdff;
  font-family: "SFMono-Regular", Consolas, monospace;
  box-shadow: 0 0 20px rgba(0, 229, 255, 0.22);
}

.gateway-node span,
.gateway-node strong {
  font-size: 13px;
  font-weight: 900;
}

.gateway-node small {
  color: #8096aa;
  font-size: 10px;
}

.gateway-node--source {
  top: 112px;
  left: 22px;
}

.gateway-node--core {
  top: 98px;
  left: calc(50% - 53px);
  width: 106px;
  min-height: 86px;
  border-color: rgba(57, 255, 20, 0.72);
  box-shadow: 0 0 30px rgba(57, 255, 20, 0.2);
}

.gateway-node--core img {
  width: 34px;
  height: 34px;
  object-fit: contain;
  image-rendering: pixelated;
}

.gateway-node--anthropic {
  top: 32px;
  right: 28px;
  border-color: rgba(255, 176, 32, 0.72);
}

.gateway-node--openai {
  right: 34px;
  top: 120px;
  border-color: rgba(178, 91, 255, 0.72);
}

.gateway-route {
  position: absolute;
  z-index: 1;
  height: 2px;
  transform-origin: left center;
  background: repeating-linear-gradient(90deg, #00e5ff 0 8px, transparent 8px 14px);
  background-size: 28px 2px;
  box-shadow: 0 0 14px rgba(0, 229, 255, 0.3);
}

.gateway-route--source {
  top: 141px;
  left: 110px;
  width: 31%;
}

.gateway-route--anthropic {
  top: 112px;
  left: 50%;
  width: 34%;
  transform: rotate(-28deg);
}

.gateway-route--openai {
  top: 142px;
  left: 50%;
  width: 34%;
}

.cyber-terminal {
  overflow: hidden;
}

.cyber-terminal__bar {
  min-height: 42px;
  padding: 0 14px;
  border-bottom: 1px solid rgba(0, 229, 255, 0.18);
  background: rgba(0, 0, 0, 0.18);
}

.cyber-terminal__dot {
  width: 10px;
  height: 10px;
  border: 2px solid #b25bff;
  background: #00e5ff;
  box-shadow: 0 0 16px rgba(178, 91, 255, 0.5);
}

.cyber-terminal__body {
  display: grid;
  gap: 10px;
  padding: 18px;
  color: #d7f7ff;
  font-size: 13px;
  line-height: 1.45;
}

.cyber-terminal__body p {
  margin: 0;
  overflow-wrap: anywhere;
}

.cyber-terminal__body span {
  color: #39ff14;
  font-weight: 900;
}

.cyber-terminal__success {
  color: #39ff14;
}

.cyber-capabilities,
.cyber-advantages,
.cyber-workflow,
.cyber-final-cta {
  margin-top: 56px;
}

.cyber-section-heading {
  display: grid;
  max-width: 760px;
  gap: 10px;
  margin-bottom: 20px;
}

.cyber-section-heading--center {
  margin-right: auto;
  margin-left: auto;
  text-align: center;
}

.cyber-section-heading span,
.cyber-final-cta__eyebrow {
  color: #39ff14;
  font-size: 12px;
  font-weight: 900;
  letter-spacing: 0;
}

.cyber-section-heading h2,
.cyber-final-cta h2 {
  margin: 0;
  color: #f7fcff;
  font-size: 34px;
  font-weight: 900;
  line-height: 1.15;
  overflow-wrap: break-word;
  text-shadow: 0 0 20px rgba(0, 229, 255, 0.18);
}

.cyber-section-heading p,
.cyber-final-cta p {
  margin: 0;
  color: #9db4c7;
  font-size: 15px;
  line-height: 1.75;
  overflow-wrap: break-word;
}

.cyber-capability-grid {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: 14px;
}

.cyber-capability-card {
  position: relative;
  min-width: 0;
  overflow: hidden;
  padding: 20px;
  transition:
    border-color 180ms ease,
    box-shadow 180ms ease,
    transform 180ms ease;
}

.cyber-card-icon {
  display: grid;
  width: 42px;
  height: 42px;
  place-items: center;
  border: 1px solid rgba(57, 255, 20, 0.48);
  border-radius: 6px;
  color: #39ff14;
  background: rgba(57, 255, 20, 0.08);
}

.cyber-capability-card h3,
.cyber-advantage-card h3,
.cyber-workflow-step h3 {
  margin: 16px 0 0;
  color: #ffffff;
  font-size: 18px;
  font-weight: 900;
  line-height: 1.3;
  overflow-wrap: break-word;
}

.cyber-capability-card p,
.cyber-advantage-card p,
.cyber-workflow-step p {
  margin: 10px 0 0;
  color: #9db4c7;
  font-size: 13px;
  line-height: 1.68;
  overflow-wrap: break-word;
}

.cyber-capability-card code {
  display: inline-flex;
  max-width: 100%;
  margin-top: 16px;
  padding: 6px 8px;
  border: 1px solid rgba(0, 229, 255, 0.24);
  border-radius: 5px;
  color: #d7f7ff;
  background: rgba(0, 0, 0, 0.2);
  font-family: "SFMono-Regular", Consolas, monospace;
  font-size: 12px;
  overflow-wrap: anywhere;
}

.cyber-advantage-grid {
  display: grid;
  grid-template-columns: repeat(4, minmax(0, 1fr));
  gap: 14px;
}

.cyber-advantage-card {
  position: relative;
  display: grid;
  min-height: 214px;
  overflow: hidden;
  padding: 18px;
  transition:
    border-color 180ms ease,
    box-shadow 180ms ease,
    transform 180ms ease;
}

.cyber-advantage-card::before,
.cyber-capability-card::before,
.cyber-workflow-step::before {
  position: absolute;
  inset: 0;
  pointer-events: none;
  background:
    linear-gradient(110deg, transparent 0 38%, rgba(0, 229, 255, 0.11) 48%, transparent 58%),
    repeating-linear-gradient(0deg, rgba(255, 255, 255, 0.04) 0 1px, transparent 1px 10px);
  opacity: 0;
  transform: translateX(-80%);
  content: "";
}

.cyber-advantage-card__index {
  color: #00e5ff;
  font-size: 12px;
  font-weight: 900;
}

.cyber-advantage-card__body {
  align-self: end;
}

.cyber-advantage-card__meter {
  position: absolute;
  right: 16px;
  bottom: 14px;
  left: 16px;
  height: 3px;
  overflow: hidden;
  background: rgba(0, 229, 255, 0.12);
}

.cyber-advantage-card__meter::before {
  display: block;
  width: 62%;
  height: 100%;
  background: linear-gradient(90deg, #39ff14, #00e5ff);
  box-shadow: 0 0 14px rgba(0, 229, 255, 0.4);
  content: "";
}

.cyber-advantage-card:hover,
.cyber-capability-card:hover,
.cyber-workflow-step:hover {
  border-color: rgba(57, 255, 20, 0.5);
  box-shadow:
    inset 0 0 0 1px rgba(57, 255, 20, 0.1),
    0 20px 56px rgba(0, 229, 255, 0.12);
  transform: translateY(-3px);
}

.cyber-advantage-card:hover::before,
.cyber-capability-card:hover::before,
.cyber-workflow-step:hover::before {
  opacity: 1;
}

.cyber-workflow {
  display: grid;
  grid-template-columns: minmax(0, 0.85fr) minmax(0, 1.15fr);
  align-items: start;
  gap: 28px;
}

.cyber-workflow-list {
  display: grid;
  gap: 12px;
  margin: 0;
  padding: 0;
  list-style: none;
}

.cyber-workflow-step {
  position: relative;
  display: grid;
  grid-template-columns: auto minmax(0, 1fr);
  gap: 14px;
  overflow: hidden;
  padding: 18px;
  transition:
    border-color 180ms ease,
    box-shadow 180ms ease,
    transform 180ms ease;
}

.cyber-workflow-step__number {
  display: grid;
  width: 42px;
  height: 42px;
  place-items: center;
  border: 1px solid rgba(0, 229, 255, 0.48);
  border-radius: 6px;
  color: #00e5ff;
  background: rgba(0, 229, 255, 0.08);
  font-size: 12px;
  font-weight: 900;
}

.cyber-workflow-step h3 {
  margin-top: 0;
}

.cyber-final-cta {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 24px;
  padding: 28px;
}

.cyber-final-cta div {
  min-width: 0;
}

.cyber-final-cta p {
  max-width: 640px;
  margin-top: 10px;
}

@media (max-width: 980px) {
  .cyber-hero,
  .cyber-workflow {
    grid-template-columns: 1fr;
    min-height: 0;
    gap: 34px;
  }

  .cyber-title {
    font-size: 56px;
  }

  .cyber-capability-grid,
  .cyber-advantage-grid {
    grid-template-columns: repeat(2, minmax(0, 1fr));
  }

  .cyber-final-cta {
    align-items: flex-start;
    flex-direction: column;
  }
}

@media (max-width: 720px) {
  .cyber-header {
    padding: 12px 12px 0;
  }

  .cyber-nav {
    align-items: flex-start;
    flex-direction: column;
  }

  .cyber-actions {
    display: flex;
    width: 100%;
    flex-wrap: wrap;
    justify-content: start;
  }

  .cyber-main {
    width: calc(100% - 24px);
    max-width: 1180px;
    padding: 36px 0 30px;
  }

  .cyber-title {
    font-size: 42px;
  }

  .cyber-title strong {
    font-size: 22px;
  }

  .cyber-description {
    font-size: 15px;
  }

  .cyber-signal-list,
  .cyber-capability-grid,
  .cyber-advantage-grid {
    grid-template-columns: 1fr;
  }

  .cyber-section-heading h2,
  .cyber-final-cta h2 {
    font-size: 28px;
  }

  .gateway-route-map {
    height: 254px;
  }

  .gateway-node {
    min-width: 76px;
  }

  .gateway-node--source {
    left: 12px;
  }

  .gateway-node--anthropic,
  .gateway-node--openai {
    right: 14px;
  }
}

@media (max-width: 460px) {
  .cyber-brand__mode,
  .cyber-icon-button--wide span,
  .cyber-login-link span {
    display: none;
  }

  .cyber-actions {
    width: 100%;
  }

  .cyber-icon-button--wide,
  .cyber-login-link {
    width: auto;
    min-width: 38px;
    padding: 0;
  }

  .cyber-title {
    font-size: 36px;
  }

  .cyber-title strong {
    font-size: 19px;
  }

  .cyber-cta-row {
    display: grid;
    width: 100%;
  }

  .cyber-button {
    width: 100%;
  }
}

@media (prefers-reduced-motion: no-preference) {
  .cyber-home::before {
    animation: cyber-grid-drift 22s linear infinite;
  }

  .cyber-backdrop__noise {
    animation: cyber-noise-shift 16s steps(8, end) infinite;
  }

  .cyber-eyebrow__pulse {
    animation: cyber-pulse 1.8s steps(2, end) infinite;
  }

  .cyber-title > span {
    animation: cyber-title-shift 7s linear infinite;
  }

  .gateway-route {
    animation: gateway-route-flow 1.8s linear infinite;
  }

  .cyber-terminal__body p {
    opacity: 0;
    transform: translateY(8px);
    animation: terminal-line-in 460ms ease forwards;
    animation-delay: calc(var(--line) * 100ms);
  }

  .cyber-advantage-card,
  .cyber-capability-card,
  .cyber-workflow-step {
    animation: card-in 520ms ease both;
    animation-delay: var(--delay, 0ms);
  }

  .cyber-advantage-card:hover::before,
  .cyber-capability-card:hover::before,
  .cyber-workflow-step:hover::before {
    animation: card-scan 1.5s linear;
  }

  .cyber-advantage-card__meter::before {
    animation: advantage-meter 2.8s ease-in-out infinite;
    animation-delay: var(--delay);
  }
}

@keyframes cyber-grid-drift {
  0% {
    background-position: 0 0, 0 0;
  }
  100% {
    background-position: 0 56px, 56px 0;
  }
}

@keyframes cyber-noise-shift {
  0%,
  100% {
    transform: translate3d(0, 0, 0);
  }
  35% {
    transform: translate3d(7px, -5px, 0);
  }
  70% {
    transform: translate3d(-4px, 6px, 0);
  }
}

@keyframes cyber-pulse {
  0%,
  100% {
    opacity: 1;
  }
  50% {
    opacity: 0.45;
  }
}

@keyframes cyber-title-shift {
  0% {
    background-position: 0% 50%;
  }
  100% {
    background-position: 240% 50%;
  }
}

@keyframes gateway-route-flow {
  0% {
    background-position: 0 0;
  }
  100% {
    background-position: 28px 0;
  }
}

@keyframes terminal-line-in {
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes card-in {
  from {
    opacity: 0;
    transform: translateY(14px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes card-scan {
  0% {
    transform: translateX(-90%);
  }
  100% {
    transform: translateX(90%);
  }
}

@keyframes advantage-meter {
  0%,
  100% {
    width: 48%;
  }
  50% {
    width: 82%;
  }
}
</style>
