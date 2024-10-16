import '@/assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import router from '@/router'
import DiveLogApp from '@/DiveLogApp.vue'

const app = createApp(DiveLogApp)
const pinia = createPinia()

app.use(router)
app.use(pinia)

app.mount('#app')
