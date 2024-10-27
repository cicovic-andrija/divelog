<script setup lang="ts">
import { onMounted, inject, ref, watch } from 'vue'
import { NCard, NGrid, NGi } from 'naive-ui'
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
    <div v-else>
      <n-card title="Details" size="large">
        <n-grid :cols="2">
          <n-gi>a</n-gi>
          <n-gi>b</n-gi>
        </n-grid>
      </n-card>
    </div>
  </div>
</template>
