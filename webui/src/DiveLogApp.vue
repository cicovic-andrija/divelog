<script setup lang="ts">
import { ref, computed, provide } from 'vue'
import {RouterView } from 'vue-router'
import { NConfigProvider, NGlobalStyle, darkTheme, lightTheme } from 'naive-ui'
import { saveStringToLocalStorage, loadStringFromLocalStorage } from '@/utils'
import DocumentHeader from '@/components/DocumentHeader.vue'

const themeName = ref(savedThemeNameOrDefault())
provide('themeName', themeName)

const theme = computed(() => {
  return themeName.value === 'light' ? lightTheme : darkTheme
})

function changeTheme(): void {
  if (themeName.value === 'light') {
    themeName.value = 'dark'
  } else {
    themeName.value = 'light'
  }
  saveThemeName(themeName.value)
}

function savedThemeNameOrDefault(): string {
  return loadStringFromLocalStorage('divelogThemeName') ?? 'light'
}

function saveThemeName(themeName: string): void {
  saveStringToLocalStorage('divelogThemeName', themeName)
}
</script>

<template>
  <n-config-provider :theme="theme">
    <document-header :themeName="themeName" @themeChangeRequested="changeTheme()"/>
    <router-view />
    <n-global-style />
  </n-config-provider>
</template>
