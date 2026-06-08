import { readFileSync } from 'node:fs'
import { dirname, resolve } from 'node:path'
import { fileURLToPath } from 'node:url'

import { describe, expect, it } from 'vitest'

const componentPath = resolve(dirname(fileURLToPath(import.meta.url)), '../TablePageLayout.vue')
const componentSource = readFileSync(componentPath, 'utf8')

describe('TablePageLayout mobile scrolling styles', () => {
  it('lets mobile admin table pages use document scrolling', () => {
    const mobileLayoutBlockMatch = componentSource.match(/\.table-page-layout\.mobile-mode\s*\{[\s\S]*?\n\}/)

    expect(mobileLayoutBlockMatch).not.toBeNull()
    expect(mobileLayoutBlockMatch?.[0]).toContain('height: auto;')
    expect(mobileLayoutBlockMatch?.[0]).toContain('overflow: visible;')
  })
})
