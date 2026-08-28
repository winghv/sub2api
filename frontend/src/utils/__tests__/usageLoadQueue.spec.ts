import { describe, expect, it } from 'vitest'
import { enqueueUsageRequest } from '../usageLoadQueue'
import type { Account } from '@/types'

// enqueueUsageRequest 已退化为立即执行（后端改为被动采样，无需前端节流）。
// 本套用例守护该契约：请求立即执行、透传返回值、失败正常 reject。
function fakeAccount(id: number): Account {
  return { id } as Account
}

describe('usageLoadQueue', () => {
  it('请求立即执行（无排队）', async () => {
    const order: number[] = []

    const p1 = enqueueUsageRequest(fakeAccount(1), async () => {
      order.push(1)
      return 'a'
    })
    const p2 = enqueueUsageRequest(fakeAccount(2), async () => {
      order.push(2)
      return 'b'
    })
    const p3 = enqueueUsageRequest(fakeAccount(3), async () => {
      order.push(3)
      return 'c'
    })

    expect(order).toEqual([1, 2, 3])
    await Promise.all([p1, p2, p3])
    expect(order).toEqual([1, 2, 3])
  })

  it('请求失败时正常 reject', async () => {
    const p = enqueueUsageRequest(fakeAccount(99), async () => {
      throw new Error('fail')
    })

    await expect(p).rejects.toThrow('fail')
  })

  it('返回值正确透传', async () => {
    const result = await enqueueUsageRequest(fakeAccount(1), async () => {
      return { usage: 42 }
    })
    expect(result).toEqual({ usage: 42 })
  })
})
