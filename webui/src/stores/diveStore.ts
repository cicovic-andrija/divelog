import { ref } from 'vue'
import { defineStore } from 'pinia'
import { type Dive } from '@/types'
import { sleep } from '@/utils'
import { DEV_PAUSE_MS } from '@/constants'

export const useDiveStore = defineStore('dives', () => {
    const dives = ref<Map<number, Dive>>(new Map<number, Dive>([
        [1, {
            id: 1,
            label: 'Um El Faroud',
            cardinal: 1,
            max_depth: '33 m'
        }],
        [2, {
            id: 2,
            label: 'Tugboat Rozi',
            cardinal: 2,
            max_depth: '31 m'
        }],
        [3, {
            id: 3,
            label: 'Zenobia',
            cardinal: 3,
            max_depth: '16 m'
        }]
    ]))

    async function find(id: number) {
        await sleep(DEV_PAUSE_MS)
        return dives.value.get(id)
    }

    return { dives, find }
})
