<template>
  <div class="app-shell dark min-h-screen">
    <div class="app-backdrop" aria-hidden="true">
      <span class="app-backdrop__beam app-backdrop__beam--cyan"></span>
      <span class="app-backdrop__beam app-backdrop__beam--violet"></span>
      <span class="app-backdrop__scan"></span>
    </div>

    <!-- Sidebar -->
    <AppSidebar />

    <!-- Main Content Area -->
    <div
      class="app-content relative min-h-screen transition-all duration-300"
      :class="[sidebarCollapsed ? 'lg:ml-[72px]' : 'lg:ml-64']"
    >
      <!-- Header -->
      <AppHeader />

      <!-- Main Content -->
      <main class="app-main p-4 md:p-6 lg:p-8">
        <slot />
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import '@/styles/onboarding.css'
import { computed, onMounted } from 'vue'
import { useAppStore } from '@/stores'
import { useAuthStore } from '@/stores/auth'
import { useOnboardingTour } from '@/composables/useOnboardingTour'
import { useOnboardingStore } from '@/stores/onboarding'
import AppSidebar from './AppSidebar.vue'
import AppHeader from './AppHeader.vue'

const appStore = useAppStore()
const authStore = useAuthStore()
const sidebarCollapsed = computed(() => appStore.sidebarCollapsed)
const isAdmin = computed(() => authStore.user?.role === 'admin')

const { replayTour } = useOnboardingTour({
  storageKey: isAdmin.value ? 'admin_guide' : 'user_guide',
  autoStart: true
})

const onboardingStore = useOnboardingStore()

onMounted(() => {
  onboardingStore.setReplayCallback(replayTour)
})

defineExpose({ replayTour })
</script>

<style scoped>
.app-shell {
  position: relative;
  overflow: hidden;
  background:
    radial-gradient(circle at 78% 16%, rgba(0, 229, 255, 0.13), transparent 30%),
    radial-gradient(circle at 18% 76%, rgba(57, 255, 20, 0.08), transparent 30%),
    linear-gradient(180deg, rgba(5, 7, 18, 0.95) 0%, rgba(7, 10, 23, 0.98) 46%, #050711 100%),
    #050711;
  color: #f5fbff;
  color-scheme: dark;
}

.app-shell::before,
.app-shell::after {
  position: fixed;
  inset: 0;
  pointer-events: none;
  content: "";
}

.app-shell::before {
  z-index: 0;
  background:
    repeating-linear-gradient(0deg, rgba(74, 222, 255, 0.055) 0 1px, transparent 1px 28px),
    repeating-linear-gradient(90deg, rgba(74, 222, 255, 0.04) 0 1px, transparent 1px 28px);
  mask-image: linear-gradient(to bottom, rgba(0, 0, 0, 0.84), transparent 92%);
}

.app-shell::after {
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

.app-backdrop {
  position: fixed;
  inset: 0;
  z-index: 2;
  pointer-events: none;
}

.app-backdrop__beam {
  position: absolute;
  width: 46vw;
  height: 170px;
  transform: rotate(-16deg);
  border: 1px solid rgba(255, 255, 255, 0.08);
  opacity: 0.42;
  clip-path: polygon(8% 0, 100% 0, 92% 100%, 0 100%);
}

.app-backdrop__beam--cyan {
  top: 14%;
  right: -24%;
  background: linear-gradient(90deg, transparent, rgba(0, 229, 255, 0.18), transparent);
}

.app-backdrop__beam--violet {
  bottom: 10%;
  left: -32%;
  background: linear-gradient(90deg, transparent, rgba(178, 91, 255, 0.14), transparent);
}

.app-backdrop__scan {
  position: absolute;
  inset: 0;
  opacity: 0.07;
  background-image:
    linear-gradient(90deg, rgba(57, 255, 20, 0.12) 1px, transparent 1px),
    linear-gradient(rgba(57, 255, 20, 0.08) 1px, transparent 1px);
  background-size: 7px 7px;
}

.app-content {
  position: relative;
  z-index: 3;
}

.app-main {
  position: relative;
  z-index: 3;
}

@media (prefers-reduced-motion: no-preference) {
  .app-shell::before {
    animation: app-grid-drift 22s linear infinite;
  }

  .app-backdrop__scan {
    animation: app-noise-shift 16s steps(8, end) infinite;
  }
}

@keyframes app-grid-drift {
  0% {
    background-position: 0 0, 0 0;
  }
  100% {
    background-position: 0 56px, 56px 0;
  }
}

@keyframes app-noise-shift {
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
</style>
