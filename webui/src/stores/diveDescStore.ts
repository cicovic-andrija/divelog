import { ref } from 'vue'
import { defineStore } from 'pinia'
import { type DiveTripDesc, type StatusResponse } from '@/types'

export const useDiveDescStore = defineStore('diveDescriptors', () => {
  const diveDescriptors = ref<DiveTripDesc[]>()

  async function fetchAll(): Promise<StatusResponse> {
    if (diveDescriptors.value === undefined) {
      try {
        const resp = await fetch('https://my-json-server.typicode.com/cicovic-andrija/jsonmock/trips')
        const body = resp.ok ? await resp.json() : undefined
        if (body) {
          diveDescriptors.value = body
        }
        return {
          ok: diveDescriptors.value !== undefined,
          status: resp.status,
          error: null
        }
      } catch (e) {
        return {
          ok: false,
          status: 0,
          error: e
        }
      }
    }
    return {
      ok: true,
      status: 200,
      error: null
    }
  }

  function firstId(): number | undefined {
    return diveDescriptors.value?.[0].id
  }

  return { diveDescriptors, fetchAll, firstId }
})
