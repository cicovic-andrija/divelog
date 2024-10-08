import './assets/main.css'

import { createApp } from 'vue'
import DiveLogApp from './DiveLogApp.vue'
import router from './router'

const app = createApp(DiveLogApp)

app.use(router)

app.mount('#app')
