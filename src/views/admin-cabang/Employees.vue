<script setup>
import { ref, onMounted, watch } from 'vue'
import { useRouter, useRoute} from 'vue-router'
import { useAuth } from '@/composables/useAuth'
import AdminSidebar from '@/components/AdminSidebar.vue'
import {
  getEmployees,
  addEmployee,
  updateEmployee,
  toggleEmployeeStatus,
  deleteEmployee
} from '@/services/adminCabang'

const router = useRouter()
const route = useRoute()
const { user, logout } = useAuth()

const employees = ref('')
const search = ref('')
const status = ref('')
const loading = ref(false)

const page = ref(1)
const limit = ref()
const meta = ref({total:0, total_pages: 1})

const showModal = ref(false)
const modalMode = ref('add')
const modalLoading = ref(false)
const modalError = ref('')

const form = ref({
    id: null,
    username: '',
    role: 'karyawan',
    tipe: 'cabang',
    position: '',
    division_id: null,
    status: 'active'
})

onMounted(fetchEmployee)

watch(page, fetchEmployee)

async function fetchEmployee() {
    loading.value = true
    try {
        const res = await getEmployees ({
            search: search.value || undefined,
            status: status.value || undefined,
            page: page.value,
            limit: limit.value 
        })

        employees.value = res.data.data
        meta.value = res.data.meta || { total: 0, total_page: 1}

    } catch (err) {
        console.error(err)
    } finally {
        loading.value = false
    }
}

function handleFilter() {
    page.value = 1
    fetchEmployees()
}

function formatDate(iso) {
  if (!iso) return '-'
  return new Date(iso).toLocaleDateString('id-ID', {
    day: 'numeric',
    month: 'long',
    year: 'numeric'
  })
}

function empCode(id) {
    return 'EMP' + string(id).padStart(3, '0')
}

function openAdd() {
    modalMode.value = 'add'
    modalError.value = ''

    form.value = {
        id: null,
        username: '',
        password:'',
        role: 'karyawan',
        tipe: 'cabang',
        position: '',
        division_id: null,
        status: 'active'
    }
    showModal.value = true
}

function openEdit(emp) {
    modalMode.value = 'edit'
    modalError.value = ''

    form.value = {
        id: emp.id,
        username: emp.username,
        password: '',
        role: emp.role,
        tipe: emp.tipe,
        position: emp.position || '',
        division_id: emp.division_id,
        status: emp.status
    }
    showModal.value = true
}

function closeModal() {
    showModal.value = false
    modalError.value =''
}

async function submitModal() {
    if(!form.value.username.trim()) {
        modalError.error.value = 'Username wajib diisi'
        return
    }

    if(modalMode.value === 'add' && !form.value.password.trim()) {
        modalError.value ='Password wajib diisi'
        return
    }

    modalLoading.value = true
    modalError.value = ''
}
</script>
<template>
<div class="layout">
    <AdminSidebar></AdminSidebar>

    <main class="main">
        <div class="header">
            <div>
                <h2>Employees</h2>
                <p class="subtitle">
                    Manage employees in your branch
                </p>
            </div>

            <AdminProfile :user="user"/>
        </div>

        <div class="toolbar">
            <div class="search-wrap">
                <svg width="14" height="14" fill="none" viewBox="0 0 24 24">
                    <circle cx="11" cy="11" r="8" stroke="currentColor" stroke-width="2"/>
                    <path d="M21 21l-4.35-4.35" stroke="currentColor" stroke-width="2" stroke-linecap="round"/>
                </svg>
                <input
                    v-model="search"
                    placeholder="Search Employee..."
                    @keyup.enter="handleFilter"
                />
            </div>
        </div>
    </main>
</div>
</template>

<style scope>
.layout {
  display: flex;
  height: 100vh;
  background: #f0f2ff;
  font-family: 'Segoe UI', sans-serif;
  overflow: hidden;
}
</style>