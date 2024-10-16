<script setup lang="ts">
import { ref, h } from 'vue'
import { NH2 } from 'naive-ui'
import JsonBlock from '@/components/JsonBlock.vue'
defineProps({
    isPretty: Boolean
})

const jsonData = ref(
  {"a": 123, "b":456, "c": ["a", "b", "c"], "d": true}
)

</script>

<template>
  <div v-if="!isPretty">
    <json-block title="Details" :object="jsonData" />
  </div>
  <div v-else>
    <n-h2>Details</n-h2>
  </div>
</template>
