import { createRouter, createWebHistory } from 'vue-router'

import Login from '../views/auth/Login.vue'
import AdminPusatDashboard from '../views/admin-pusat/Dashboard.vue'
import AdminCabangDashboard from '../views/admin-cabang/Dashboard.vue'
import EmployeeDashboard from '../views/employee/Dashboard.vue'

const routes = [
    {
    path: '/',
    component: Login
},
{
    path: '/admin-pusat/dashboard',
    component: AdminPusatDashboard
},
{
    path: '/admin-cabang/dashboard',
    component: AdminCabangDashboard
},
{
    path: '/employee/dashboard',
    component: EmployeeDashboard
}
]

export default createRouter({
    history: createWebHistory(),
    routes
})