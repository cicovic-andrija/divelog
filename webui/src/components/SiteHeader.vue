<script setup lang="ts">
import { ref, computed, h } from 'vue'
import { RouterLink, useRoute, useRouter } from 'vue-router'
import { NLayout, NLayoutHeader, NMenu, NButton, NText } from 'naive-ui'
const route = useRoute()
const router = useRouter()
const emit = defineEmits(['themeChangeRequested'])
const props = defineProps({
  themeName: String
})

const themeButtonLabel = computed(() => {
  return props.themeName === 'light' ? 'Dark' : 'Light'
})

const selectedMenu = computed(() => {
    if (route.name === 'home')
      return 'home'
    if (/\/dives/.test(route.path))
      return 'dives'
    if (/\/sites/.test(route.path))
      return 'sites'

    return null
})

const menuOptions = ref([
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
  <n-layout>
    <n-layout-header bordered class="nav">
      <n-text class="nav-title" @click="router.push('/')">
        <img width="32px" src="@/assets/scuba.svg">
        <span>MyDives</span>
      </n-text>
      <div class="flex-cont ofw-hidden">
        <n-menu
          responsive
          mode="horizontal"
          :value="selectedMenu"
          :options="menuOptions"
          class="nav-menu"
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
        <n-button
          size="small"
          quaternary
          tag="a"
          href="https://github.com/cicovic-andrija/divelog"
          target="_blank"
        >
          GitHub
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
  cursor: pointer;
  display: flex;
  align-items: center;
  font-size: 1rem;
  padding-left: 16px;
}

/* justify menu items that are nested in another flex container */
:deep(.nav-menu > div:first-child) {
  justify-content: center;
}

.nav-end {
  align-items: center;
  justify-content: flex-end;
  padding-right: 16px;
}
</style>
