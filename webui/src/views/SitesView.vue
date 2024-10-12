<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted } from 'vue';
import { NText } from 'naive-ui'

const seconds = ref(0);
const timeout: number = 30;

let interval: number = 0;
const status = computed(() => {
  return {
    failed: seconds.value > timeout,
    message: seconds.value > timeout ? 'Download failed. Please refresh the page.' : 'Downloading dive site data from the server...',
    formattedTime: `${seconds.value}s`
  }
})

onMounted(() => {
  interval = setInterval(() => {
    if (seconds.value <= timeout) {
      seconds.value++
    }
  }, 1000);
})

onUnmounted(() => {
  clearInterval(interval)
})
</script>

<template>
  <n-text class="loading-msg">
    <div>
      <img v-if="status.failed" class="msg-icon" src="@/assets/warning.svg">
      <img v-else class="msg-icon" src="@/assets/download.svg">
      <span>{{ status.message }}</span>
    </div>
    <span v-if="!status.failed">{{ status.formattedTime }}</span>
  </n-text>
</template>
