<script setup lang="ts">
import { NCard } from 'naive-ui'
import { formatJSON } from '@/utils'
defineProps({
    title: String,
    object: Object
})
</script>

<template>
  <n-card style="margin-bottom: 16px" segmented embedded size="small" :title="title">
    <pre style="margin: 0" v-if="object">{{ formatJSON(object) }}</pre>
  </n-card>
</template>
