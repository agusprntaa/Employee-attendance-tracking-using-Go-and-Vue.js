import { createRouter, createWebHistory } from 'vue-router'

import Login from '../views/auth/Login.vue'

const routes = [
    {
        path: '/',
        name: 'Login',
        component: Login
    },

  // ADMIN
    {
        path: '/admin-pusat/dashboard',
        name: 'AdminPusatDashboard',
        component: () => import('../views/admin-pusat/Dashboard.vue')
    },

    {
        path: '/admin-cabang/dashboard',
        name: 'AdminCabangDashboard',
        component: () => import('../views/admin-cabang/Dashboard.vue')
    },

  // EMPLOYEE
    {
        path: '/employee/dashboard',
        name: 'EmployeeDashboard',
        component: () => import('../views/employee/Dashboard.vue')
    },

    { 
        path: '/employee/change-password',
        name: 'ChangePassword',
        component: () => import('../views/employee/ChangePassword.vue')
    },

    {
        path: '/employee/wfa',
        name: 'WFA',
        component: () => import('../views/employee/WFA.vue')
    },

    {
        path: '/employee/scan',
        name: 'ScanQR',
        component: () => import('../views/employee/ScanQR.vue')
    },

    {
        path: '/employee/success',
        name: 'Success',
        component: () => import('../views/employee/Success.vue')
    },

    {
        path: '/:pathMatch(.*)*',
        redirect: '/'
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

/*router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token')

  let user = null
  try {
    user = JSON.parse(localStorage.getItem('user'))
  } catch {}

  if (to.path === '/') {
    if (token && user) {
      const target =
        user.role === 'super_admin'
          ? '/admin-pusat/dashboard'
          : user.role === 'admin_cabang'
          ? '/admin-cabang/dashboard'
          : '/employee/dashboard'

      // NO LOOP
      if (to.path !== target) {
        return next(target)
      }
    }
    return next()
  }

  // BELUM LOGIN
  if (!token) {
    if (to.path !== '/') {
      return next('/')
    }
    return next()
  }

  // CEK ROLE
  if (user) {
    if (to.path.startsWith('/admin-pusat') && user.role !== 'super_admin') {
      return next('/')
    }

    if (to.path.startsWith('/admin-cabang') && user.role !== 'admin_cabang') {
      return next('/')
    }

    if (to.path.startsWith('/employee') && user.role !== 'karyawan') {
      return next('/')
    }
  }

  next()
})*/

export default router