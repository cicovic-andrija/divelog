import { ref } from 'vue'
import { defineStore } from 'pinia'
import { type DiveSiteDesc, type StatusResponse } from '@/types'

export const useSiteDescStore = defineStore('siteDescriptors', () => {
  const siteDescriptors = ref<DiveSiteDesc[]>()

  async function fetchAll(): Promise<StatusResponse> {
    if (siteDescriptors.value === undefined) {
      try {
        const resp = await fetch('http://localhost:8072/data/sites?headonly=true')
        const body = resp.ok ? await resp.json() : undefined
        if (body) {
          siteDescriptors.value = body
        }
        return {
          ok: siteDescriptors.value !== undefined,
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
    return siteDescriptors.value?.[0].id
  }

  return { siteDescriptors, fetchAll, firstId }
})
