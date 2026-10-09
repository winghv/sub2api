import { enableAutoUnmount, mount } from '@vue/test-utils'
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest'
import { ref } from 'vue'
import DateRangePicker from '../DateRangePicker.vue'

// 下拉通过 Teleport 挂到 body，需从 document 上查找；且 Teleport 节点不会随
// wrapper 卸载移除，需在 afterEach 清理 body，否则跨用例读到残留节点。
const dateInputs = () => Array.from(document.body.querySelectorAll<HTMLInputElement>('input[type="date"]'))
const presetButtons = () => Array.from(document.body.querySelectorAll<HTMLButtonElement>('.date-picker-preset'))
const applyButton = () => document.body.querySelector<HTMLButtonElement>('.date-picker-apply')!
vi.mock('vue-i18n', () => ({ useI18n: () => ({ t: (key: string) => key, locale: ref('en') }) }))
enableAutoUnmount(afterEach)
beforeEach(() => { vi.useFakeTimers(); vi.setSystemTime(new Date(2026, 8, 30, 23, 59)) })
afterEach(() => {
  vi.useRealTimers()
  document.body.innerHTML = ''
})
async function reopenNextDay() {
  const w = mount(DateRangePicker, { props: { startDate: '2026-09-30', endDate: '2026-09-30' }, global: { stubs: { Icon: true } } })
  await w.get('.date-picker-trigger').trigger('click')
  expect(dateInputs()[1].getAttribute('max')).toBe('2026-10-01')
  await w.get('.date-picker-trigger').trigger('click')
  vi.setSystemTime(new Date(2026, 9, 1, 0, 1))
  await w.get('.date-picker-trigger').trigger('click')
  return w
}
describe('date presets after midnight', () => {
  it.each([
    ['dates.today', '2026-10-01'],
    ['dates.last7Days', '2026-09-25'],
    ['dates.thisMonth', '2026-10-01'],
  ])('refreshes %s when the page stays mounted overnight', async (label, startDate) => {
    const w = await reopenNextDay()
    await presetButtons().find(b => b.textContent === label)!.click()
    await applyButton().click()
    expect(w.emitted('change')?.[0]?.[0]).toMatchObject({ startDate, endDate: '2026-10-01' })
  })
  it('updates the maximum selectable date when reopened', async () => {
    await reopenNextDay()
    expect(dateInputs()[1].getAttribute('max')).toBe('2026-10-02')
  })
})
