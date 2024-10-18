<script setup lang="ts">
import { onMounted, ref, computed, h } from 'vue'
import { RouterLink, useRouter } from 'vue-router'
import { NMenu } from 'naive-ui'
import { useSiteDescStore } from '@/stores/siteDescStore'
import { paddedID, siteIdToRoute } from '@/utils'
import GhostMenu from './GhostMenu.vue'
const router = useRouter()
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
  // devtest: sleep
  await store.fetchAll()
  loaded.value = true
  if (props.siteId == 0) {
    loadFirst()
  }
})

function loadFirst(): void {
  const first = store.firstId()
  if (first) {
    router.replace({ path: `/sites/${paddedID(first)}` })
  }
}
</script>

<template>
  <n-menu v-if="loaded" :options="menuOptions" :value="selectedOption" />
  <ghost-menu v-else />
</template>
