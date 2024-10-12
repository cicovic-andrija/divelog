<script setup lang="ts">
import { ref, computed, provide } from 'vue'
import {RouterView } from 'vue-router'
import { NConfigProvider, NGlobalStyle, darkTheme, lightTheme } from 'naive-ui'
import SiteHeader from '@/SiteHeader.vue';

const themeName = ref('light')
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
}
</script>

<template>
  <n-config-provider :theme="theme">
    <site-header :themeName="themeName" @themeChangeRequested="changeTheme()"/>
    <router-view />
    <n-global-style />
  </n-config-provider>
</template>
