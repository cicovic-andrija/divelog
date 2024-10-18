<script setup lang="ts">
import { onMounted, inject, ref, watch } from 'vue'
import { NH2, NH3 } from 'naive-ui'
import { useDiveStore } from '@/stores/diveStore'
import JsonBlock from '@/components/JsonBlock.vue'
import { type Dive } from '@/types'
import { type ResourceTitle } from '@/composables/useResourceTitle'
const store = useDiveStore()
const props = defineProps({
    diveId: Number,
    isPretty: Boolean
})

const dive = ref<Dive | undefined>()
watch(() => props.diveId, async (newId) => {
  dive.value = undefined
  if (newId) {
    dive.value = await loadDive(newId)
  }

  if (import.meta.env.DEV) {
    console.log(`watch: diveId=${newId} => dive.id=${dive.value?.id} [DiveData]`)
  }
})
onMounted(async () => {
    if (import.meta.env.DEV) {
      console.log(`diveId=${props.diveId}, isPretty=${props.isPretty} [DiveData]`)
    }

    dive.value = await loadDive(props.diveId)
})

const resourceTitle = inject<ResourceTitle>('resourceTitle')
watch(dive, (newValue) => {
  resourceTitle?.setResourceTitle(newValue ? composeDiveTitle(newValue) : undefined)

  if (import.meta.env.DEV) {
    console.log(`watch: dive.id=${dive.value?.id} => resourceTitle="${resourceTitle?.title.value}" [DiveData]`)
  }
})

async function loadDive(id: number | undefined): Promise<Dive | undefined> {
  if (id === undefined) {
    return undefined
  }
  if (id === 0) {
    id = 1
  }
  if (id) {
    return await store.find(id)
  }
  return undefined
}

function composeDiveTitle(dive: Dive): string {
  return `Dive ${dive.cardinal}. ${dive.label}`
}
</script>

<template>
  <div v-if="dive">
    <div v-if="!isPretty">
      <json-block title="Details" :object="dive" />
    </div>
    <div v-else>
      <n-h2>Details</n-h2>
      <n-h3>{{ dive.label }}</n-h3>
    </div>
  </div>
</template>
