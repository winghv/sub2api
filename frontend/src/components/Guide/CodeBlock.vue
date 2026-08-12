<template>
  <div class="group relative">
    <button
      type="button"
      class="absolute right-2 top-2 z-10 rounded-md bg-white/10 px-2 py-1 text-xs text-gray-300 transition-colors hover:bg-white/20 hover:text-white"
      @click="handleCopy"
    >
      {{ copied ? t('guide.copied') : t('guide.copy') }}
    </button>
    <pre ref="codeRef" class="overflow-x-auto rounded-lg bg-gray-900 p-4 text-sm leading-relaxed text-gray-100"><slot /></pre>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()
const codeRef = ref<HTMLPreElement | null>(null)
const copied = ref(false)

async function handleCopy() {
  if (!codeRef.value) return
  try {
    await navigator.clipboard.writeText(codeRef.value.innerText)
    copied.value = true
    setTimeout(() => {
      copied.value = false
    }, 2000)
  } catch {
    // 剪贴板不可用时静默失败，不影响阅读
  }
}
</script>
