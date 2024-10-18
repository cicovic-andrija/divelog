import { ref, type Ref } from 'vue'

export interface ResourceTitle {
    title: Ref<string | undefined>
    setResourceTitle: (newValue: string | undefined) => void
}

export function useResourceTitle(): ResourceTitle {
    const title = ref<string | undefined>(undefined)

    function setResourceTitle(newValue: string | undefined): void {
        title.value = newValue
    }

    return { title, setResourceTitle }
}
