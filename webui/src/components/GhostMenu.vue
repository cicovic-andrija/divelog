<script setup lang="ts">
import { NMenu } from 'naive-ui'

const ghostOptions = [...Array(20)].map(() => ({
    key: `o-g`,
    label: '',
    disabled: true
}))
</script>

<template>
  <n-menu :options="ghostOptions" value="o-g" />
</template>
