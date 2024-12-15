<script setup lang="ts">
import { onMounted, inject, ref, watch } from 'vue'
import { useSiteStore } from '@/stores/siteStore'
import { type DiveSite } from '@/types'
import { type ResourceTitle } from '@/composables/useResourceTitle'
import JsonBlock from '@/components/JsonBlock.vue'
const store = useSiteStore()
const props = defineProps({
    siteId: Number,
    isPretty: Boolean
})

const site = ref<DiveSite | undefined>()
const resourceTitle = inject<ResourceTitle>('resourceTitle')

onMounted(async () => {
  site.value = await loadSite(props.siteId)
})

watch(() => props.siteId, async (newId) => {
  site.value = await loadSite(newId)
})

watch(site, (newValue) => {
  resourceTitle?.setResourceTitle(newValue?.name)
})

async function loadSite(id: number | undefined): Promise<DiveSite | undefined> {
  if (id === undefined) {
    return undefined
  }
  await store.fetchSite(id)
  return store.sites.get(id)
}
</script>

<template>
  <div v-if="!isPretty">
    <json-block title="Dive Site" :object="site" />
  </div>
</template>
