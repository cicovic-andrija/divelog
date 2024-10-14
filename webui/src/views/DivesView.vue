<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue';
import { NText } from 'naive-ui'
import { TIMEOUT } from '@/utils'

const seconds = ref(0)

let interval: number = 0
const status = computed(() => {
  return {
    failed: seconds.value > TIMEOUT,
    message: seconds.value > TIMEOUT ? 'Download failed. Please refresh the page.' : 'Downloading dive data from the server...',
    formattedTime: `${seconds.value}s`
  }
})

onMounted(() => {
  interval = setInterval(() => {
    if (seconds.value <= TIMEOUT) {
      seconds.value++
    }
  }, 1000)
})

onUnmounted(() => {
  clearInterval(interval)
})
</script>

<template>
  <n-text class="loading-msg">
    <img v-if="status.failed" class="msg-icon" src="@/assets/warning.svg">
    <img v-else class="msg-icon" src="@/assets/download.svg">
    <span>{{ status.message }}</span>
    <span v-if="!status.failed">{{ status.formattedTime }}</span>
  </n-text>
</template>
