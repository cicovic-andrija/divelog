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
      redirect: '/dives/0000'
    },
    {
      path: '/sites',
      name: 'sites',
      redirect: '/sites/0000'
    },
    {
      path: '/dives/:diveID(\\d{4})',
      name: 'dive',
      component: () => import('@/views/DataViewLayout.vue'),
      props: route => ({ diveId: Number(route.params.diveID) })
    },
    {
      path: '/sites/:siteID(\\d{4})',
      name: 'dive-site',
      component: () => import('@/views/DataViewLayout.vue'),
      props: route => ({ siteId: Number(route.params.siteID) })
    },
    {
      path: '/:pathMatch(.*)*',
      name: 'not-found',
      component: () => import('@/views/NotFoundView.vue')
    }
  ]
})

export default router
