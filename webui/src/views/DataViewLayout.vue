<script setup lang="ts">
import { ref, h } from 'vue'
import { RouterLink } from 'vue-router'
import { NLayout, NLayoutSider, NMenu, NText, NCard, NSwitch } from 'naive-ui';

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

const isPretty = ref(true);

</script>

<template>
  <n-layout :has-sider="true">
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

    <n-layout content-style="padding: 48px">
      <div class="data-repr-toggle">
        <n-text style="padding: 6px">Representation</n-text>
        <n-switch v-model:value="isPretty" :round="false">
          <template #checked>
            Pretty
          </template>
          <template #unchecked>
            JSON
          </template>
        </n-switch>
      </div>

      <n-card embedded title="Dive"></n-card>
    </n-layout>
  </n-layout>
</template>

<style scoped>
.data-repr-toggle {
  padding: 6px;
  display: flex;
  align-items: center;
  justify-content: right;
}
</style>
