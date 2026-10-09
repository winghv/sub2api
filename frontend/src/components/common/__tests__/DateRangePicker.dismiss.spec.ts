import { mount } from '@vue/test-utils'
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'
import { nextTick, ref } from 'vue'
import DateRangePicker from '../DateRangePicker.vue'

// 下拉通过 Teleport 挂到 body，需从 document 上查找；且 Teleport 节点不会随
// wrapper 卸载移除，需在 afterEach 清理 body，否则跨用例读到残留节点。
const dateInputs = () => Array.from(document.body.querySelectorAll<HTMLInputElement>('input[type="date"]'))
const presetButtons = () => Array.from(document.body.querySelectorAll<HTMLButtonElement>('.date-picker-preset'))
const applyButton = () => document.body.querySelector<HTMLButtonElement>('.date-picker-apply')!

vi.mock('vue-i18n', () => ({ useI18n: () => ({ t: (key: string) => key, locale: ref('en') }) }))
let wrapper: ReturnType<typeof mount<typeof DateRangePicker>>
beforeEach(() => { vi.useFakeTimers(); vi.setSystemTime(new Date(2026, 8, 13, 12)) })
afterEach(() => {
  wrapper?.unmount()
  vi.useRealTimers()
  document.body.innerHTML = ''
})

async function chooseDraft() {
  wrapper = mount(DateRangePicker, {
    props: {
      startDate: '2026-09-13', endDate: '2026-09-13',
      'onUpdate:startDate': (startDate: string) => { void wrapper.setProps({ startDate }) },
      'onUpdate:endDate': (endDate: string) => { void wrapper.setProps({ endDate }) },
    },
  })
  await wrapper.get('.date-picker-trigger').trigger('click')
  await presetButtons().find(node => node.textContent === 'dates.last7Days')!.click()
}

describe('DateRangePicker unapplied changes', () => {
  it.each(['outside', 'escape', 'toggle'])('discards draft dates when dismissed with %s', async (method) => {
    await chooseDraft()
    if (method === 'outside') document.body.click()
    else if (method === 'escape') document.dispatchEvent(new KeyboardEvent('keydown', { key: 'Escape' }))
    else await wrapper.get('.date-picker-trigger').trigger('click')
    await nextTick()
    expect(wrapper.emitted('change')).toBeUndefined()
    expect(wrapper.get('.date-picker-trigger').text()).toContain('dates.today')
    await wrapper.get('.date-picker-trigger').trigger('click')
    expect(dateInputs().map(input => input.value))
      .toEqual(['2026-09-13', '2026-09-13'])
  })

  it('retains a newly applied range after the parent accepts both updates', async () => {
    await chooseDraft()
    await applyButton().click()
    await nextTick()
    expect(wrapper.emitted('change')).toEqual([[{ startDate: '2026-09-07', endDate: '2026-09-13', preset: '7days' }]])
    expect(wrapper.get('.date-picker-trigger').text()).toContain('dates.last7Days')
    await wrapper.get('.date-picker-trigger').trigger('click')
    expect(dateInputs().map(input => input.value))
      .toEqual(['2026-09-07', '2026-09-13'])
  })
})
