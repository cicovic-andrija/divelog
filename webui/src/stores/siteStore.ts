import { ref } from 'vue'
import { defineStore } from 'pinia'
import { type DiveSite, type StatusResponse } from '@/types'

export const useSiteStore = defineStore('sites', () => {
    const sites = ref<Map<number, DiveSite>>(new Map<number, DiveSite>())

    async function fetchSite(id: number): Promise<StatusResponse> {
        if (sites.value.has(id)) {
            return {
                ok: true,
                status: 200,
                error: null
            }
        }
        try {
            const resp = await fetch(`http://localhost:8072/data/sites/${id}`)
            const body = resp.ok ? await resp.json() : undefined
            if (body) {
                sites.value.set(body.id, body)
                return {
                    ok: true,
                    status: 200,
                    error: null
                }
            }
            return {
                ok: false,
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

    return { sites, fetchSite }
})
