import { ref } from 'vue'
import { defineStore } from 'pinia'
import { type DiveSiteDesc } from '@/types'

export const useSiteDescStore = defineStore('siteDescriptors', () => {
const siteDescriptors = ref<DiveSiteDesc[]>()

async function fetchAll(): Promise<void> {
    if (siteDescriptors.value === undefined) {
        siteDescriptors.value = [
            {
                id: 1,
                label: 'Malo Sidro',
            },
            {
                id: 2,
                label: 'Katamaran',
            },
        ]
    }
}

return { siteDescriptors, fetchAll }
})
