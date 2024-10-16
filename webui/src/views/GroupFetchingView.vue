<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import LoadingMessage from '@/components/LoadingMessage.vue'
const route = useRoute()

const msg = computed(() => {
  if (route.name === 'dives') {
    return 'Downloading dive data from the server...'
  } else if (route.name === 'sites') {
    return 'Downloading dive site data from the server...'
  } else {
    return '???'
  }
})
</script>

<template>
  <loading-message :key="route.name" :msg="msg" />
</template>
