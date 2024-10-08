<script setup lang="ts">
import { ref, computed } from 'vue'
import { NLayout, NLayoutHeader, NMenu, NButton, NText } from 'naive-ui'
const emit = defineEmits(['themeChangeRequested'])
const props = defineProps({
  themeName: String
})

const themeButtonLabel = computed(() => {
  return props.themeName === 'light' ? 'Dark' : 'Light'
})

const menuInstRef = ref()
const menuOptions = ref([
  {
    key: 'home',
    label: 'Home',
    path: '/'
  },
  {
    key: 'dives',
    label: 'Dives',
    path: '/dives'
  },
  {
    key: 'sites',
    label: 'Sites',
    path: '/sites'
  }
])
</script>

<template>
<n-layout>
  <n-layout-header bordered class="nav">
    <n-text class="nav-title">
      <img width="32px" src="./assets/scuba.svg">
      <span>MyDives</span>
    </n-text>
    <div class="flex-cont ofw-hidden">
      <n-menu
        ref="menuInstRef"
        responsive
        mode="horizontal"
        :options="menuOptions"
      />
    </div>
    <div class="flex-cont nav-end">
      <n-button
        size="small"
        quaternary
        @click="emit('themeChangeRequested')"
      >
        {{ themeButtonLabel }}
      </n-button>
    </div>
  </n-layout-header>
</n-layout>
</template>

<style scoped>
.nav {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
}

.nav-title {
  display: flex;
  align-items: center;
  font-size: 1rem;
  padding-left: 16px;
}

.nav-menu {
  align-items: center;
}

.nav-end {
  align-items: center;
}
</style>
