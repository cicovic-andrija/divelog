import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/dives',
      name: 'dives',
      component: () => import('@/views/GroupFetchingView.vue')
    },
    {
      path: '/sites',
      name: 'sites',
      component: () => import('@/views/GroupFetchingView.vue')
    },
    {
      path: '/dives/:diveID(\\d{4})',
      name: 'one-dive',
      component: () => import('@/views/DataViewLayout.vue'),
      props: true
    },
    {
      path: '/sites/:siteID(\\d{4})',
      name: 'one-site',
      component: () => import('@/views/DataViewLayout.vue'),
      props: true
    },
    {
      path: '/:pathMatch(.*)*',
      name: 'not-found',
      component: () => import('@/views/NotFound.vue')
    }
  ]
})

export default router
