<script setup lang="ts">
import { NMenu } from 'naive-ui'

const ghostOptions = [...Array(20)].map((_, i) => ({
    key: `o-g-${i}`,
    label: 'Loading...',
    disabled: true
}))
</script>

<template>
  <n-menu :options="ghostOptions" />
</template>
