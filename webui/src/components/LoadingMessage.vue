<script setup lang="ts">
import { ref, computed,  onMounted, onUnmounted } from 'vue'
import { NText } from 'naive-ui'
import { DOWNLOAD_TIMEOUT, DOWNLOAD_FAILED_MSG } from '@/constants'
const props = defineProps({
    msg: String
})

const seconds = ref(0)
let interval: number = 0
const status = computed(() => {
  return {
    failed: seconds.value > DOWNLOAD_TIMEOUT,
    message: seconds.value > DOWNLOAD_TIMEOUT ? DOWNLOAD_FAILED_MSG : props.msg,
    formattedTime: `${seconds.value}s`
  }
})

onMounted(() => {
  interval = setInterval(() => {
    if (seconds.value <= DOWNLOAD_TIMEOUT) {
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
    <img v-else class="msg-icon" src="@/assets/gear.svg">
    <span>{{ status.message }}</span>
    <span v-if="!status.failed">{{ status.formattedTime }}</span>
  </n-text>
</template>
