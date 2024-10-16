<script setup lang="ts">
import { ref, h } from 'vue'
import { RouterLink } from 'vue-router'
import { NLayout, NLayoutSider, NMenu, NText, NSwitch, NH1 } from 'naive-ui';
import DiveData from '@/components/DiveData.vue';

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

const isPretty = ref(!import.meta.env.DEV)

</script>

<template>
  <n-layout :has-sider="true">
    <n-layout-sider
      :native-scrollbar="false"
      :collapsed-width="0"
      collapse-mode="transform"
      trigger-style="top: 240px;"
      collapsed-trigger-style="top: 240px; right: -16px;"
      bordered
      show-trigger="arrow-circle"
    >
      <n-menu
        :options="vertMenuOptions"
      />
    </n-layout-sider>

    <n-layout content-style="padding: 16px 32px 16px 32px;">
      <n-h1>Dive #2</n-h1>

      <div class="data-repr-toggle">
        <n-text style="padding-right: 8px">Data View</n-text>
        <n-switch v-model:value="isPretty" :round="false">
          <template #checked>
            Pretty
          </template>
          <template #unchecked>
            JSON
          </template>
        </n-switch>
      </div>

      <dive-data :is-pretty="isPretty" />
    </n-layout>
  </n-layout>
</template>

<style scoped>
.data-repr-toggle {
  padding-bottom: 16px;
  display: flex;
  align-items: center;
  justify-content: right;
}
</style>
