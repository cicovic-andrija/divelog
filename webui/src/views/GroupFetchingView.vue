<script setup lang="ts">
import { onMounted, computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useDiveDescStore } from '@/stores/diveDescStore'
import { diveIdToRoute, sleep } from '@/utils'
import { DEV_PAUSE_MS } from '@/constants'
import LoadingMessage from '@/components/LoadingMessage.vue'

const route = useRoute()
const router = useRouter()
const store = useDiveDescStore()

const msg = computed(() =>
  route.name === 'dives' ? 'Downloading dive data...' :
  route.name === 'sites' ? 'Downloading dive site data...' : '?'
)

onMounted(async () => {
  if (import.meta.env.DEV) {
    await sleep(DEV_PAUSE_MS)
  }

  if (store.diveDescriptors === undefined) {
    await store.fetchAll()
  }

  if (store.diveDescriptors !== undefined) {
    router.push(diveIdToRoute(store.diveDescriptors[0].id))
  }
})
</script>

<template>
  <loading-message :key="route.name" :msg="msg" />
</template>
