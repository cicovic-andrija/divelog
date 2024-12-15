import { ref } from 'vue'
import { defineStore } from 'pinia'
import { type Dive, type StatusResponse } from '@/types'

export const useDiveStore = defineStore('dives', () => {
    const dives = ref<Map<number, Dive>>(new Map<number, Dive>())

    async function fetchDive(id:number): Promise<StatusResponse> {
        if (dives.value.has(id)) {
            return {
                ok: true,
                status: 200,
                error: null
            }
        }
        try {
            const resp = await fetch(`http://localhost:8072/data/dives/${id}`)
            const body = resp.ok ? await resp.json() : undefined
            if (body) {
                dives.value.set(body.id, body)
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

    return { dives, fetchDive }
})
