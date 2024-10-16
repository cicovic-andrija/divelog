<script setup lang="ts">
import { onMounted, ref, computed, h } from 'vue'
import { RouterLink } from 'vue-router'
import { useDiveDescStore } from '@/stores/diveDescStore';
import { NLayout, NLayoutSider, NMenu, NText, NSwitch, NH1 } from 'naive-ui'
import { diveDescToLabel, diveIdToRoute, paddedID, sleep } from '@/utils'
import { DEV_PAUSE_MS } from '@/constants'
import DiveData from '@/components/DiveData.vue'
import LoadingMessage from '@/components/LoadingMessage.vue'
const store = useDiveDescStore()
const props = defineProps({
  diveID: String,
  siteID: String
})

const vertMenuOptions = computed(() => store.diveDescriptors?.map(t => ({
  key: `o-t-${paddedID(t.id)}`,
  type: 'group',
  label: `${t.trip}`,
  children: t.descriptors.map(d => ({
    key: `o-d-${paddedID(d.id)}`,
    label: () => h(
      RouterLink,
      { to: diveIdToRoute(d.id) },
      { default: () => diveDescToLabel(d) }
    ),
  })),
})))

const selectedOption = computed(() => {
  if (props.diveID !== undefined) {
    const num: number = Number(props.diveID)
    return isNaN(num) ? null : `o-d-${paddedID(num)}`
  }
  if (props.siteID !== undefined) {
    const num: number = Number(props.siteID)
    return isNaN(num) ? null : `o-s-${paddedID(num)}`
  }
  return null
})

onMounted(async () => {
  if (import.meta.env.DEV) {
    await sleep(DEV_PAUSE_MS)
  }

  if (store.diveDescriptors === undefined) {
    await store.fetchAll()
  }
})

const isPretty = ref(!import.meta.env.DEV)
</script>

<template>
  <n-layout :has-sider="true">
    <n-layout-sider
      v-if="store.current"
      :native-scrollbar="false"
      :collapsed-width="0"
      collapse-mode="transform"
      trigger-style="top: 196px;"
      collapsed-trigger-style="top: 196px; right: -16px;"
      bordered
      show-trigger="arrow-circle"
    >
      <n-menu
        :options="vertMenuOptions"
        :value="selectedOption"
      />
    </n-layout-sider>

    <n-layout content-style="padding: 16px 32px 16px 32px;">
      <n-h1 v-if="store.current">
        Dive {{ diveDescToLabel(store.current) }}
      </n-h1>

      <div v-if="store.current" class="data-repr-toggle">
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

      <dive-data v-if="store.current" :is-pretty="isPretty" />
      <loading-message v-else msg="Downloading details of the dive..." />
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
