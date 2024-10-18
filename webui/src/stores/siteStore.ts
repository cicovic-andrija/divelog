import { ref } from 'vue'
import { defineStore } from 'pinia'
import { type DiveSite } from '@/types'
import { sleep } from '@/utils'
import { DEV_PAUSE_MS } from '@/constants'

export const useSiteStore = defineStore('sites', () => {
    const sites = ref<Map<number, DiveSite>>(new Map<number, DiveSite>([
        [1, {
            id: 1,
            label: 'Malo Sidro',
        }],
        [2, {
            id: 2,
            label: 'Katamaran',
        }]
    ]))

    async function find(id: number) {
        await sleep(DEV_PAUSE_MS)
        return sites.value.get(id)
    }

    return { sites, find }
})
