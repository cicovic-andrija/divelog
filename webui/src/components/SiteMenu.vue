<script setup lang="ts">
import { onMounted, ref, computed, h } from 'vue'
import { RouterLink } from 'vue-router'
import { NMenu } from 'naive-ui'
import { useSiteDescStore } from '@/stores/siteDescStore'
import { sleep, paddedID, siteIdToRoute } from '@/utils'
import { DEV_PAUSE_MS } from '@/constants'
import GhostMenu from './GhostMenu.vue'
const store = useSiteDescStore()
const loaded = ref(false)
const props = defineProps({
  siteId: Number
})

const menuOptions = computed(() => store.siteDescriptors?.map(s => ({
  key: `o-s-${paddedID(s.id)}`,
  label: () => h(RouterLink, { to: siteIdToRoute(s.id) }, { default: () => s.label }),
})))

const selectedOption = computed(() => `o-s-${paddedID(props.siteId ?? 0)}`)

onMounted(async () => {
  if (import.meta.env.DEV) {
    await sleep(DEV_PAUSE_MS)
  }

  await store.fetchAll()
  loaded.value = true
})
</script>

<template>
  <n-menu v-if="loaded" :options="menuOptions" :value="selectedOption" />
  <ghost-menu v-else />
</template>
