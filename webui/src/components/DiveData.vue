<script setup lang="ts">
import { onMounted, inject, ref, watch } from 'vue'
import { useDiveStore } from '@/stores/diveStore'
import JsonBlock from '@/components/JsonBlock.vue'
import { type Dive } from '@/types'
import { type ResourceTitle } from '@/composables/useResourceTitle'
import { composeDiveTitle } from '@/utils'
const store = useDiveStore()
const props = defineProps({
    diveId: Number,
    isPretty: Boolean
})

const dive = ref<Dive | undefined>()
const resourceTitle = inject<ResourceTitle>('resourceTitle')

onMounted(async () => {
    dive.value = await loadDive(props.diveId)
})

watch(() => props.diveId, async (newId) => {
  dive.value = await loadDive(newId)
})

watch(dive, (newValue) => {
  resourceTitle?.setResourceTitle(newValue !== undefined ? composeDiveTitle(newValue) : undefined)
})

async function loadDive(id: number | undefined): Promise<Dive | undefined> {
  return (id === undefined) ? undefined : await store.find(id)
}
</script>

<template>
  <div v-if="dive">
    <div v-if="!isPretty">
      <json-block title="Details" :object="dive" />
    </div>
  </div>
</template>
