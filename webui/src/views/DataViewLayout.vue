<script setup lang="ts">
import { ref, computed, provide, watch } from 'vue'
import { NLayout, NLayoutSider, NText, NSwitch } from 'naive-ui'
import { useResourceTitle } from '@/composables/useResourceTitle'
import LoadingTitle from '@/components/LoadingTitle.vue'
import DiveMenu from '@/components/DiveMenu.vue'
import DiveData from '@/components/DiveData.vue'
import SiteMenu from '@/components/SiteMenu.vue'
import SiteData from '@/components/SiteData.vue'
const props = defineProps({
  diveId: Number,
  siteId: Number
})
const isDiveSite = computed(() => props.diveId === undefined)
const resourceId = computed(() => props.diveId ?? props.siteId ?? 0)
const key = computed(() => `${isDiveSite.value ? 's' : 'd'}-${resourceId.value}`)
const isPretty = ref<boolean>(!import.meta.env.DEV)

const resourceTitle = useResourceTitle()
watch(key, () => { resourceTitle.setResourceTitle(undefined) })
provide('resourceTitle', resourceTitle)
</script>

<template>
  <n-layout :has-sider="true">
    <!-- Side Menu -->
    <n-layout-sider
      :native-scrollbar="false"
      :collapsed-width="0"
      collapse-mode="transform"
      trigger-style="top: 40vh;"
      collapsed-trigger-style="top: 40vh; right: -16px;"
      bordered
      show-trigger="arrow-circle"
    >
      <site-menu v-if="isDiveSite" :site-id="resourceId" />
      <dive-menu v-else :dive-id="resourceId" />
    </n-layout-sider>

    <!-- content-class doesn't work for some reason -->
    <n-layout content-style="padding: 16px 32px 16px 32px;">
      <!-- Title -->
      <loading-title :key="key" />
      <!-- Toggle -->
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
      <!-- Main -->
      <site-data v-if="isDiveSite" :site-id="resourceId" :is-pretty="isPretty" />
      <dive-data v-else :dive-id="resourceId" :is-pretty="isPretty" />
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
