import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/views/Home.vue'
import FileConverter from '@/components/FileConverter.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: Home,
  },
  {
    path: '/converter/:type',
    name: 'converter',
    component: FileConverter,
    props: true,
  },
]

export const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
