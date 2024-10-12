<script setup lang="ts">
import { ref, h } from 'vue'
import { RouterLink } from 'vue-router'
import { NLayout, NLayoutSider, NMenu } from 'naive-ui';

const vertMenuOptions = ref([
  {
    key: 'home',
    label: () => h(RouterLink, { to: '/' }, { default: () => 'Home' })
  },
  {
    key: 'dives',
    label: () => h(RouterLink, { to: '/dives' }, { default: () => 'Dives' })
  },
  {
    key: 'sites',
    label: () => h(RouterLink, { to: '/sites' }, { default: () => 'Sites' })
  }
])

</script>

<template>
  <n-layout
    :has-sider="true"
  >
    <n-layout-sider
      :native-scrollbar="false"
      :collapsed-width="0"
      collapse-mode="transform"
      collapsed-trigger-style="right: -16px;"
      bordered
      show-trigger="arrow-circle"
    >
      <n-menu
        :options="vertMenuOptions"
      />
    </n-layout-sider>
  </n-layout>
</template>
