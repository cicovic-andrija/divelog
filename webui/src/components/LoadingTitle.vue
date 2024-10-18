<script setup lang="ts">
import { onMounted, onUnmounted, computed, inject } from 'vue'
import { NH1 } from 'naive-ui'
import { useTimer } from '@/composables/useTimer'
import { type ResourceTitle } from '@/composables/useResourceTitle';
const { seconds, start, stop, timedOut } = useTimer()

const loadingSequence = [
  'Loading',
  'Loading.',
  'Loading..',
  'Loading...',
]
const resourceTitle = inject<ResourceTitle>('resourceTitle')
const title = computed(() =>
  resourceTitle?.title.value ?? (!timedOut.value ? loadingSequence[seconds.value % 4] : 'Failed.')
)

onMounted(() => start())
onUnmounted(() => stop())
</script>

<template>
  <n-h1>
    {{ title }}
  </n-h1>
</template>
