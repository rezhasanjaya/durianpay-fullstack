//Utils
import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '../stores/auth'

//Page
import Login from '../views/Login.vue'
import NotFound from '../views/errors/404.vue'
import Dashboard from '../views/Dashboard.vue'


const routes = [
  {
    path: '/',
    redirect: '/login',
  },
  {
    path: '/login',
    name: 'login',
    component: Login,
    meta: { guestOnly: true },
  },
  {
    path: '/dashboard',
    name: 'dashboard',
    component: Dashboard,
    meta: { requiresAuth: true }
  },
  { 
    path: '/:pathMatch(.*)*', 
    name: 'NotFound', 
    component: NotFound 
  },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
})

router.beforeEach((to) => {
  const auth = useAuthStore()

  if (to.meta.requiresAuth && !auth.isLogin) {
    return '/'
  }

  if (to.meta.guestOnly && auth.isLogin) {
    return '/dashboard'
  }

  return true
})

export default router
