<script setup lang="ts">
import { onMounted, ref, computed, h } from 'vue'
import { RouterLink, useRouter } from 'vue-router'
import { NMenu } from 'naive-ui'
import { useDiveDescStore } from '@/stores/diveDescStore'
import { paddedID, diveIdToRoute } from '@/utils'
import GhostMenu from './GhostMenu.vue'
const router = useRouter()
const store = useDiveDescStore()
const loaded = ref(false)
const props = defineProps({
  diveId: Number
})

const menuOptions = computed(() => store.diveDescriptors?.map(t => ({
  key: `o-t-${paddedID(t.id)}`,
  type: 'group',
  label: `${t.label}`,
  children: t.linked_dives.map(d => ({
    key: `o-d-${paddedID(d.id)}`,
    label: () => h(
      RouterLink,
      { to: diveIdToRoute(d.id) },
      { default: () => d.short_label }
    ),
  })),
})))

const selectedOption = computed(() => `o-d-${paddedID(props.diveId ?? 0)}`)

onMounted(async () => {
  // devtest: sleep
  await store.fetchAll()
  loaded.value = store.diveDescriptors !== undefined
  if (props.diveId == 0) {
    loadFirst()
  }
})

function loadFirst(): void {
  const first = store.firstId()
  if (first) {
    router.replace({ path: `/dives/${paddedID(first)}` })
  }
}
</script>

<template>
  <n-menu v-if="loaded" :options="menuOptions" :value="selectedOption" />
  <ghost-menu v-else />
</template>
