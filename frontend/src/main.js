import 'boxicons'
import { createApp } from 'vue'
import App from './App.vue'

import './assets/main.css'



import { createRouter, createWebHashHistory, createWebHistory  } from 'vue-router'
import Settings from './Settings.vue'
import Home from './Home.vue'
import Cpu from './Cpu.vue'
import Scripts from './Scripts.vue'
import Store from './Store.vue'

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: '/settings',
      name: 'Settings',
      component: Settings
    },
    {
      path: '/',
      name: 'Home',
      component: Home
    },
    {
      path: '/scripts',
      name: 'Scripts',
      component: Scripts
    },
    {
      path: '/monitor',
      name: 'Monitor',
      component: Cpu
    },
    {
      path: '/store',
      name: 'Store',
      component: Store
    }

  ]
})
export default router

const app = createApp(App)

app.use(router)

app.mount('#app')

