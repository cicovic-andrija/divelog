import { ref } from 'vue'
import { defineStore } from 'pinia'
import { type DiveTripDesc } from '@/types'

export const useDiveDescStore = defineStore('diveDescriptors', () => {
const diveDescriptors = ref<DiveTripDesc[]>()

async function fetchAll(): Promise<void> {
    if (diveDescriptors.value === undefined) {
        diveDescriptors.value = [
            {
                label: 'Malta',
                id: 1,
                descriptors: [
                    {
                        id: 1,
                        label: 'Um El Faroud',
                        cardinal: 1,
                    },
                    {
                        id: 2,
                        label: 'Tugboat Rozi',
                        cardinal: 2,
                    },
                ],
            },
            {
                label: 'Cyprus',
                id: 2,
                descriptors: [
                    {
                        id: 3,
                        label: 'Zenobia',
                        cardinal: 3,
                    },
                ],
            },
        ]
    }
}

return { diveDescriptors, fetchAll }
})
