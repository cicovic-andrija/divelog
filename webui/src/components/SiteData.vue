<script setup lang="ts">
import { ref } from 'vue'
import { NH2 } from 'naive-ui'
import JsonBlock from '@/components/JsonBlock.vue'
defineProps({
    siteId: Number,
    isPretty: Boolean
})
const jsonData = ref(
  {"a": 123, "b":456, "c": ["a", "b", "c"], "d": true}
)
</script>

<template>
  <div v-if="!isPretty">
    <json-block title="Details Site" :object="jsonData" />
  </div>
  <div v-else>
    <n-h2>Site Details</n-h2>
  </div>
</template>
