import { ref } from 'vue'
import { defineStore } from 'pinia'
import { type DiveSite, type StatusResponse } from '@/types'

export const useSiteStore = defineStore('sites', () => {
    const sites = ref<Map<number, DiveSite>>(new Map<number, DiveSite>())

    async function fetchAll(): Promise<StatusResponse> {
        if (sites.value.size === 0) {
            try {
                const resp = await fetch('http://localhost:8072/data/sites')
                const body = resp.ok ? await resp.json() : undefined
                if (body) {
                    for (const site of body) {
                        sites.value.set(site.id, site)
                    }
                }
                return {
                    ok: sites.value.size > 0,
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

    async function find(id: number): Promise<DiveSite | undefined> {
        return sites.value.get(id)
    }

    return { sites, fetchAll, find }
})
