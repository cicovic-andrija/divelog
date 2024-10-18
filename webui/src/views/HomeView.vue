<script setup lang="ts">
import { inject } from 'vue'
import { useRouter } from 'vue-router';
import { NLayout, NLayoutFooter, NH1, NP, NButton } from 'naive-ui'
const router = useRouter()
const themeName: string = inject('themeName', 'light' /* default */)
</script>

<template>
  <n-layout :native-scrollbar="false">
    <div class="banner">
      <n-h1 class="banner-title">
        <span>MyDives</span>
      </n-h1>
      <n-p style="font-size: 16px; margin-top: 0; margin-bottom: 0;">
        My personal dive log.
      </n-p>
      <n-p style="font-size: 16px; font-weight: 500; margin-top: 4px; margin-bottom: 16px;">
        Powered by
          <a href="https://go.dev/" target="_blank">Go</a>,
          <a href="https://vuejs.org/" target="_blank">Vue 3</a>,
          <a href="https://www.naiveui.com/" target="_blank">NaiveUI</a> and
          <a href="https://subsurface-divelog.org/" target="_blank">Subsurface</a>
      </n-p>
      <n-button
        type="primary"
        style="width: fit-content;"
        :ghost="themeName === 'dark'"
        @click="router.push('/dives')"
      >
        Explore
      </n-button>
    </div>
    <n-layout-footer bordered>
      <div class="footer-content">
        <n-p>
          Made by Andrija Cicović - 2025
        </n-p>
      </div>
    </n-layout-footer>
  </n-layout>
</template>

<style scoped>
.banner {
  height: calc(100vh - 64px);
  display: flex;
  flex-direction: column;
  position: relative;
  align-items: center;
  justify-content: center;
  overflow: hidden;
}

.banner::after {
  content: '';
  width: 100%;
  height: 64px;
}

.banner-title {
  line-height: 1;
  margin-top: 0;
  font-size: 64px;
  font-weight: 600;
}

.footer-content {
  display: flex;
  align-items: center;
  justify-content: center;
  line-height: 1;
  overflow: hidden;
}
</style>
