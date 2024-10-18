<script setup lang="ts">
import { onMounted, ref, computed, h } from 'vue'
import { RouterLink } from 'vue-router'
import { NMenu } from 'naive-ui'
import { useDiveDescStore } from '@/stores/diveDescStore'
import { sleep, paddedID, diveIdToRoute, diveDescToLabel } from '@/utils'
import { DEV_PAUSE_MS } from '@/constants'
import GhostMenu from './GhostMenu.vue'
const store = useDiveDescStore()
const loaded = ref(false)
const props = defineProps({
  diveId: Number
})

const menuOptions = computed(() => store.diveDescriptors?.map(t => ({
  key: `o-t-${paddedID(t.id)}`,
  type: 'group',
  label: `${t.label}`,
  children: t.descriptors.map(d => ({
    key: `o-d-${paddedID(d.id)}`,
    label: () => h(
      RouterLink,
      { to: diveIdToRoute(d.id) },
      { default: () => diveDescToLabel(d) }
    ),
  })),
})))

const selectedOption = computed(() => `o-d-${paddedID(props.diveId ?? 0)}`)

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
