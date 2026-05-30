<template>
  <div class="auth-shell">
    <div class="auth-backdrop" aria-hidden="true">
      <span class="auth-backdrop__beam auth-backdrop__beam--cyan"></span>
      <span class="auth-backdrop__beam auth-backdrop__beam--violet"></span>
      <span class="auth-backdrop__scan"></span>
    </div>

    <main class="auth-layout">
      <section class="auth-brand-panel" :aria-label="siteName">
        <router-link to="/home" class="auth-brand-link">
          <span class="auth-brand-mark">
            <img :src="siteLogo || DEFAULT_SITE_LOGO" :alt="siteName" />
          </span>
          <span class="auth-brand-copy">
            <span class="auth-brand-name">{{ siteName }}</span>
            <span class="auth-brand-mode">{{ t('home.cyber.brandMode') }}</span>
          </span>
        </router-link>

        <div class="auth-hero-copy">
          <p class="auth-eyebrow">{{ t('home.cyber.eyebrow') }}</p>
          <h1>{{ siteSubtitle }}</h1>
          <p>{{ t('home.cyber.description') }}</p>
        </div>

        <div class="auth-route-card" aria-hidden="true">
          <div class="auth-route-card__header">
            <span>{{ t('home.cyber.networkTitle') }}</span>
            <strong>{{ t('home.cyber.statusOnline') }}</strong>
          </div>
          <div class="auth-route-map">
            <span class="auth-route-node auth-route-node--client">Client</span>
            <span class="auth-route-line auth-route-line--in"></span>
            <span class="auth-route-node auth-route-node--core">{{ siteName }}</span>
            <span class="auth-route-line auth-route-line--out-one"></span>
            <span class="auth-route-line auth-route-line--out-two"></span>
            <span class="auth-route-node auth-route-node--claude">Claude</span>
            <span class="auth-route-node auth-route-node--gpt">GPT</span>
          </div>
        </div>

        <dl class="auth-signal-grid">
          <div v-for="signal in authSignals" :key="signal.label">
            <dt>{{ signal.label }}</dt>
            <dd>{{ signal.value }}</dd>
          </div>
        </dl>
      </section>

      <section class="auth-form-panel">
        <div class="auth-card">
          <slot />
        </div>

        <div class="auth-footer">
          <slot name="footer" />
        </div>

        <div class="auth-copyright">
          &copy; {{ currentYear }} {{ siteName }}. {{ t('home.footer.allRightsReserved') }}
        </div>
      </section>
    </main>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted } from 'vue'
import { useI18n } from 'vue-i18n'
import { useAppStore } from '@/stores'
import { sanitizeUrl } from '@/utils/url'
import { DEFAULT_SITE_LOGO, DEFAULT_SITE_NAME, isDefaultSiteSubtitle } from '@/utils/brand'

const appStore = useAppStore()
const { t } = useI18n()

const siteName = computed(() => appStore.siteName || DEFAULT_SITE_NAME)
const siteLogo = computed(() => sanitizeUrl(appStore.siteLogo || '', { allowRelative: true, allowDataUrl: true }))
const localizedDefaultSubtitle = computed(() => t('home.cyber.subtitle'))
const siteSubtitle = computed(() => {
  const configuredSubtitle = appStore.cachedPublicSettings?.site_subtitle?.trim()
  if (isDefaultSiteSubtitle(configuredSubtitle)) {
    return localizedDefaultSubtitle.value
  }
  return configuredSubtitle || localizedDefaultSubtitle.value
})
const authSignals = computed(() => [
  { label: t('home.cyber.signals.gateway'), value: t('home.cyber.signals.gatewayValue') },
  { label: t('home.cyber.signals.models'), value: t('home.cyber.signals.modelsValue') },
  { label: t('home.cyber.signals.billing'), value: t('home.cyber.signals.billingValue') }
])

const currentYear = computed(() => new Date().getFullYear())

onMounted(() => {
  appStore.fetchPublicSettings()
})
</script>

<style scoped>
.auth-shell {
  position: relative;
  min-height: 100vh;
  overflow: hidden;
  background:
    radial-gradient(circle at 78% 18%, rgba(0, 229, 255, 0.13), transparent 30%),
    radial-gradient(circle at 16% 72%, rgba(57, 255, 20, 0.08), transparent 30%),
    linear-gradient(180deg, rgba(5, 7, 18, 0.95) 0%, rgba(7, 10, 23, 0.98) 46%, #050711 100%),
    #050711;
  color: #f5fbff;
  color-scheme: dark;
}

.auth-shell::before,
.auth-shell::after {
  position: absolute;
  inset: 0;
  pointer-events: none;
  content: "";
}

.auth-shell::before {
  z-index: 0;
  background:
    repeating-linear-gradient(0deg, rgba(74, 222, 255, 0.055) 0 1px, transparent 1px 28px),
    repeating-linear-gradient(90deg, rgba(74, 222, 255, 0.04) 0 1px, transparent 1px 28px);
  mask-image: linear-gradient(to bottom, rgba(0, 0, 0, 0.84), transparent 92%);
}

.auth-shell::after {
  z-index: 1;
  opacity: 0.1;
  background: repeating-linear-gradient(
    180deg,
    transparent 0,
    transparent 6px,
    rgba(255, 255, 255, 0.08) 7px
  );
  mix-blend-mode: screen;
}

.auth-backdrop,
.auth-layout {
  position: relative;
  z-index: 2;
}

.auth-backdrop {
  position: absolute;
  inset: 0;
  pointer-events: none;
}

.auth-backdrop__beam {
  position: absolute;
  width: 46vw;
  height: 170px;
  transform: rotate(-16deg);
  border: 1px solid rgba(255, 255, 255, 0.08);
  opacity: 0.42;
  clip-path: polygon(8% 0, 100% 0, 92% 100%, 0 100%);
}

.auth-backdrop__beam--cyan {
  top: 14%;
  right: -24%;
  background: linear-gradient(90deg, transparent, rgba(0, 229, 255, 0.18), transparent);
}

.auth-backdrop__beam--violet {
  bottom: 10%;
  left: -32%;
  background: linear-gradient(90deg, transparent, rgba(178, 91, 255, 0.14), transparent);
}

.auth-backdrop__scan {
  position: absolute;
  inset: 0;
  opacity: 0.07;
  background-image:
    linear-gradient(90deg, rgba(57, 255, 20, 0.12) 1px, transparent 1px),
    linear-gradient(rgba(57, 255, 20, 0.08) 1px, transparent 1px);
  background-size: 7px 7px;
}

.auth-layout {
  display: grid;
  grid-template-columns: minmax(0, 1.05fr) minmax(390px, 0.72fr);
  gap: 42px;
  width: min(1180px, calc(100% - 48px));
  max-width: 100%;
  min-height: 100vh;
  align-items: center;
  margin: 0 auto;
  padding: 42px 0;
}

.auth-brand-panel,
.auth-form-panel {
  min-width: 0;
}

.auth-brand-link {
  display: inline-flex;
  max-width: 100%;
  align-items: center;
  gap: 12px;
  color: inherit;
  text-decoration: none;
}

.auth-brand-mark {
  display: grid;
  width: 50px;
  height: 50px;
  flex: 0 0 auto;
  place-items: center;
  border: 1px solid rgba(0, 229, 255, 0.58);
  border-radius: 6px;
  background:
    linear-gradient(135deg, rgba(0, 229, 255, 0.2), rgba(178, 91, 255, 0.14)),
    #07111f;
  box-shadow: 0 0 24px rgba(0, 229, 255, 0.18);
}

.auth-brand-mark img {
  width: 39px;
  height: 39px;
  object-fit: contain;
  image-rendering: pixelated;
}

.auth-brand-copy {
  display: grid;
  min-width: 0;
  gap: 2px;
}

.auth-brand-name,
.auth-brand-mode,
.auth-eyebrow,
.auth-route-card,
.auth-signal-grid {
  font-family: "SFMono-Regular", "Cascadia Code", "Roboto Mono", Consolas, monospace;
}

.auth-brand-name {
  overflow: hidden;
  color: #ffffff;
  font-size: 16px;
  font-weight: 900;
  line-height: 1.2;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.auth-brand-mode {
  color: #00e5ff;
  font-size: 11px;
  font-weight: 800;
  line-height: 1.2;
}

.auth-hero-copy {
  display: grid;
  max-width: 700px;
  gap: 16px;
  margin-top: 58px;
}

.auth-eyebrow {
  margin: 0;
  color: #39ff14;
  font-size: 12px;
  font-weight: 900;
  line-height: 1.4;
}

.auth-hero-copy h1 {
  max-width: 680px;
  margin: 0;
  color: #f7fcff;
  font-size: 56px;
  font-weight: 900;
  line-height: 1.02;
  overflow-wrap: break-word;
  text-shadow: 0 0 24px rgba(0, 229, 255, 0.22);
}

.auth-hero-copy p:last-child {
  max-width: 660px;
  margin: 0;
  color: #9db4c7;
  font-size: 16px;
  line-height: 1.8;
  overflow-wrap: break-word;
}

.auth-route-card,
.auth-card {
  max-width: 100%;
  border: 1px solid rgba(0, 229, 255, 0.24);
  border-radius: 8px;
  background:
    linear-gradient(135deg, rgba(0, 229, 255, 0.07), transparent 42%),
    rgba(5, 12, 27, 0.82);
  box-shadow:
    inset 0 0 0 1px rgba(255, 255, 255, 0.03),
    0 18px 50px rgba(0, 0, 0, 0.24);
  backdrop-filter: blur(18px);
}

.auth-route-card {
  width: min(560px, 100%);
  margin-top: 32px;
  padding: 14px;
}

.auth-route-card__header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 12px;
  color: #9db4c7;
  font-size: 12px;
  line-height: 1.2;
}

.auth-route-card__header strong {
  color: #39ff14;
  font-weight: 900;
}

.auth-route-map {
  position: relative;
  height: 172px;
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

.auth-route-node {
  position: absolute;
  z-index: 2;
  display: grid;
  min-width: 86px;
  min-height: 42px;
  place-items: center;
  padding: 8px 10px;
  border: 1px solid rgba(0, 229, 255, 0.62);
  border-radius: 6px;
  background: rgba(8, 20, 38, 0.96);
  color: #f8fdff;
  font-size: 12px;
  font-weight: 900;
  box-shadow: 0 0 20px rgba(0, 229, 255, 0.22);
}

.auth-route-node--client {
  top: 64px;
  left: 18px;
}

.auth-route-node--core {
  top: 58px;
  left: calc(50% - 52px);
  min-height: 54px;
  border-color: rgba(57, 255, 20, 0.72);
  box-shadow: 0 0 30px rgba(57, 255, 20, 0.2);
}

.auth-route-node--claude {
  top: 28px;
  right: 18px;
  border-color: rgba(255, 176, 32, 0.72);
}

.auth-route-node--gpt {
  right: 18px;
  bottom: 28px;
  border-color: rgba(178, 91, 255, 0.72);
}

.auth-route-line {
  position: absolute;
  z-index: 1;
  height: 2px;
  transform-origin: left center;
  background: repeating-linear-gradient(90deg, #00e5ff 0 8px, transparent 8px 14px);
  background-size: 28px 2px;
  box-shadow: 0 0 14px rgba(0, 229, 255, 0.3);
}

.auth-route-line--in {
  top: 85px;
  left: 104px;
  width: 28%;
}

.auth-route-line--out-one {
  top: 74px;
  left: 50%;
  width: 36%;
  transform: rotate(-18deg);
}

.auth-route-line--out-two {
  top: 97px;
  left: 50%;
  width: 36%;
  transform: rotate(18deg);
}

.auth-signal-grid {
  display: grid;
  grid-template-columns: repeat(3, minmax(0, 1fr));
  gap: 10px;
  width: min(560px, 100%);
  margin: 18px 0 0;
}

.auth-signal-grid div {
  min-width: 0;
  padding: 12px;
  border: 1px solid rgba(0, 229, 255, 0.22);
  border-radius: 6px;
  background: rgba(8, 17, 36, 0.5);
}

.auth-signal-grid dt {
  color: #8096aa;
  font-size: 11px;
  line-height: 1.3;
}

.auth-signal-grid dd {
  margin: 5px 0 0;
  color: #f3fbff;
  font-size: 13px;
  font-weight: 800;
  line-height: 1.35;
  overflow-wrap: anywhere;
}

.auth-form-panel {
  width: min(100%, 480px);
  justify-self: end;
}

.auth-card {
  padding: 28px;
}

.auth-footer {
  margin-top: 18px;
  text-align: center;
  font-size: 14px;
}

.auth-copyright {
  margin-top: 24px;
  color: #617386;
  font-size: 12px;
  line-height: 1.5;
  text-align: center;
}

:deep(.auth-form) {
  color: #f5fbff;
}

:deep(.auth-form > .text-center) {
  text-align: left;
}

:deep(.auth-form h2) {
  color: #ffffff;
  font-size: 28px;
  font-weight: 900;
  line-height: 1.15;
  text-shadow: 0 0 20px rgba(0, 229, 255, 0.18);
}

:deep(.auth-form h2 + p) {
  margin-top: 10px;
  color: #9db4c7;
  font-size: 14px;
  line-height: 1.6;
}

:deep(.auth-form .input-label) {
  margin-bottom: 8px;
  color: #cde4f2;
  font-family: "SFMono-Regular", "Cascadia Code", "Roboto Mono", Consolas, monospace;
  font-size: 12px;
  font-weight: 900;
}

:deep(.auth-form .input) {
  min-height: 46px;
  border-color: rgba(0, 229, 255, 0.28);
  border-radius: 6px;
  background: rgba(4, 12, 26, 0.78);
  color: #f7fcff;
  box-shadow: inset 0 0 0 1px rgba(255, 255, 255, 0.025);
}

:deep(.auth-form .input::placeholder) {
  color: #617386;
}

:deep(.auth-form .input:focus) {
  border-color: rgba(57, 255, 20, 0.62);
  --tw-ring-color: rgba(57, 255, 20, 0.18);
}

:deep(.auth-form .input-error) {
  border-color: rgba(248, 113, 113, 0.9);
  --tw-ring-color: rgba(248, 113, 113, 0.2);
}

:deep(.auth-form .input-hint) {
  color: #8096aa;
  line-height: 1.5;
}

:deep(.auth-form .btn-primary) {
  min-height: 50px;
  border: 1px solid rgba(57, 255, 20, 0.72);
  border-radius: 6px;
  background: linear-gradient(135deg, #39ff14 0%, #00e5ff 100%);
  color: #031007;
  font-weight: 900;
  box-shadow: 0 0 30px rgba(57, 255, 20, 0.2);
  clip-path: polygon(0 0, calc(100% - 12px) 0, 100% 12px, 100% 100%, 12px 100%, 0 calc(100% - 12px));
}

:deep(.auth-form .btn-primary:hover) {
  box-shadow: 0 0 38px rgba(0, 229, 255, 0.3);
}

:deep(.auth-form .btn-primary svg) {
  color: inherit;
}

:deep(.auth-form a),
.auth-footer :deep(a) {
  color: #39ff14;
  font-weight: 800;
}

:deep(.auth-form a:hover),
.auth-footer :deep(a:hover) {
  color: #00e5ff;
}

:deep(.auth-form .bg-gray-200),
:deep(.auth-form .dark\:bg-dark-700) {
  background-color: rgba(0, 229, 255, 0.18);
}

:deep(.auth-form .text-gray-500),
:deep(.auth-form .dark\:text-dark-400),
.auth-footer :deep(.text-gray-500),
.auth-footer :deep(.dark\:text-dark-400) {
  color: #9db4c7;
}

:deep(.auth-card h2),
:deep(.auth-card h1) {
  color: #ffffff !important;
  letter-spacing: 0;
}

:deep(.auth-card p),
:deep(.auth-card .text-gray-300),
:deep(.auth-card .text-gray-400),
:deep(.auth-card .text-gray-500),
:deep(.auth-card .text-gray-600),
:deep(.auth-card .text-gray-700),
:deep(.auth-card .dark\:text-gray-300),
:deep(.auth-card .dark\:text-gray-400),
:deep(.auth-card .dark\:text-dark-400) {
  color: #9db4c7 !important;
}

:deep(.auth-card .text-gray-900),
:deep(.auth-card .dark\:text-white) {
  color: #f7fcff !important;
}

:deep(.auth-card .bg-gray-50),
:deep(.auth-card .bg-white),
:deep(.auth-card .dark\:bg-dark-800\/60),
:deep(.auth-card .dark\:bg-dark-900\/50) {
  border-color: rgba(0, 229, 255, 0.18) !important;
  background: rgba(8, 16, 34, 0.66) !important;
}

:deep(.auth-card .border-gray-200),
:deep(.auth-card .dark\:border-dark-600) {
  border-color: rgba(0, 229, 255, 0.2) !important;
}

:deep(.auth-card label.bg-white),
:deep(.auth-card label.dark\:bg-dark-900\/50) {
  border-radius: 8px;
  box-shadow: inset 0 0 0 1px rgba(255, 255, 255, 0.025);
}

:deep(.auth-card input[type='checkbox']) {
  accent-color: #39ff14;
}

@media (max-width: 980px) {
  .auth-layout {
    grid-template-columns: 1fr;
    align-items: start;
    gap: 28px;
    padding: 28px 0;
  }

  .auth-form-panel {
    width: min(100%, 560px);
    justify-self: stretch;
  }

  .auth-hero-copy {
    margin-top: 34px;
  }

  .auth-hero-copy h1 {
    font-size: 42px;
  }
}

@media (max-width: 640px) {
  .auth-layout {
    width: calc(100% - 24px);
  }

  .auth-route-card,
  .auth-signal-grid {
    display: none;
  }

  .auth-hero-copy {
    gap: 12px;
    margin-top: 28px;
  }

  .auth-hero-copy h1 {
    font-size: 32px;
  }

  .auth-hero-copy p:last-child {
    font-size: 14px;
  }

  .auth-card {
    padding: 20px;
  }

  :deep(.auth-form h2) {
    font-size: 24px;
  }
}

@media (prefers-reduced-motion: no-preference) {
  .auth-shell::before {
    animation: auth-grid-drift 22s linear infinite;
  }

  .auth-backdrop__scan {
    animation: auth-noise-shift 16s steps(8, end) infinite;
  }

  .auth-route-line {
    animation: auth-route-flow 1.8s linear infinite;
  }
}

@keyframes auth-grid-drift {
  0% {
    background-position: 0 0, 0 0;
  }
  100% {
    background-position: 0 56px, 56px 0;
  }
}

@keyframes auth-noise-shift {
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

@keyframes auth-route-flow {
  0% {
    background-position: 0 0;
  }
  100% {
    background-position: 28px 0;
  }
}
</style>
