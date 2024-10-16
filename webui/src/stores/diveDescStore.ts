import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import { type DiveDesc, type DiveTripDesc } from '@/types'

export const useDiveDescStore = defineStore('diveDescriptors', () => {
    const diveDescriptors = ref<DiveTripDesc[]>()

    const current = computed(() => diveDescriptors.value?.[0].descriptors[0])

    async function fetchAll(): Promise<void> {
        diveDescriptors.value = [
            {
                trip: 'Malta',
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
                trip: 'Cyprus',
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

    return { diveDescriptors, current, fetchAll }
})
