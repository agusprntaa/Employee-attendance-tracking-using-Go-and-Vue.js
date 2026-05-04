<script setup>
import { ref, onMounted, onUnmounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useLocation } from '@/composables/useLocation'
import { useAuth } from '@/composables/useAuth'
import { getAttendanceHistory } from '@/services/attendance'
import API from '@/services/api'

import ProfileCard from '@/components/ProfileCard.vue'
import LocationBanner from '@/components/LocationBanner.vue'

const router = useRouter()
const { user, loadUser } = useAuth()

const loading = ref(false)
const currentTime = ref('')
const history = ref([])
const todayData = ref(null)

const {
  isInRadius,
  getCurrentLocation,
  distance,
  nearestOffice
} = useLocation()

let interval = null

const alreadyCheckedIn = computed(() => {
  return todayData.value?.has_checked_in || false
})

const canCheckIn = computed(() => {
  return isInRadius.value && !alreadyCheckedIn.value && !loading.value
})

// INIT
onMounted(async () => {
  loadUser()
  startClock()

  try {
    await getCurrentLocation()
    await fetchToday()
    await fetchHistory()
  } catch (err) {
    console.error(err)
    alert('Gagal memuat data')
  }
})

onUnmounted(() => {
  clearInterval(interval)
})

// CLOCK
function startClock() {
  interval = setInterval(() => {
    currentTime.value = new Date().toLocaleTimeString('id-ID')
  }, 1000)
}

// TODAY
async function fetchToday() {
  try {
    const res = await API.get('/attendance/today')
    todayData.value = res.data.data
  } catch (err) {
    console.error('TODAY ERROR:', err)
  }
}

// HISTORY
async function fetchHistory() {
  try {
    const res = await getAttendanceHistory(1, 10)
    history.value = res.data.data || []
  } catch (err) {
    console.error('HISTORY ERROR:', err)
  }
}

// NAVIGATION
function goToScan() {
  if (!canCheckIn.value) return
  router.push('/employee/scan')
}

function goToWFA() {
  if (loading.value) return
  router.push('/employee/wfa')
}

// TIMEZONE
function formatTime(utc) {
  if (!utc) return '-'

  return new Date(utc).toLocaleTimeString('id-ID', {
    hour: '2-digit',
    minute: '2-digit',
    timeZone: 'Asia/Jakarta'
  })
}

// FORMAT STATUS
function formatStatus(item) {
  if (item.work_type === 'WFA') return 'WFA'
  if (item.status === 'PRESENT') return 'HADIR'
  if (item.status === 'LATE') return 'TERLAMBAT'
  return item.status || '-'
}

function statusClass(item) {
  if (item.work_type === 'WFA') return 'wfa'
  if (item.status === 'PRESENT') return 'hadir'
  if (item.status === 'LATE') return 'late'
  return ''
}
</script>

<template>
  <div class="wrapper">
    <div class="content">
      
      <ProfileCard v-if="user" :user="user" />
      
      <LocationBanner
        :isInRadius="isInRadius"
        :distance="distance"
        :nearestOffice="nearestOffice"
      />

      <div class="clock">
        <h1>{{ currentTime }}</h1>
        <p>Waktu sekarang</p>
      </div>

      <button
        class="btn"
        @click="goToScan"
        :disabled="!canCheckIn"
      >
        {{
          alreadyCheckedIn
            ? 'SUDAH ABSEN'
            : (!isInRadius ? 'DI LUAR RADIUS' : 'CHECK IN')
        }}
      </button>

      <button
        class="btn-outline"
        @click="goToWFA"
        :disabled="loading || alreadyCheckedIn"
      >
        {{ alreadyCheckedIn ? 'SUDAH ABSEN' : 'WFA' }}
      </button>

      <div class="history">
        <h3>Riwayat Absensi</h3>

        <div
          class="item"
          v-for="item in history"
          :key="item.id"
        >
          <div>
            <strong>{{ item.date }}</strong>
            <p>{{ formatTime(item.check_in) }}</p>
          </div>

          <span class="status" :class="statusClass(item)">
            {{ formatStatus(item) }}
          </span>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.wrapper {
  min-height: 100vh;
  background: #f3f4f6;
}

.content {
  width: 100%;
  max-width: 640px;
  margin: 0 auto;
  padding: 20px;
}

@media (min-width: 1024px) {
  .content {
    max-width: 820px;
    padding: 30px 40px;
  }
}

.clock {
  text-align: center;
  margin: 24px 0;
}

.clock h1 {
  font-size: 36px;
  font-weight: bold;
}

@media (min-width: 1024px) {
  .clock h1 {
    font-size: 42px;
  }
}

.btn {
  width: 100%;
  padding: 16px;
  border-radius: 14px;
  background: #4f46e5;
  color: white;
  border: none;
  font-weight: 600;
  cursor: pointer;
}

.btn-outline {
  width: 100%;
  padding: 16px;
  border-radius: 14px;
  border: 2px solid #4f46e5;
  background: transparent;
  color: #4f46e5;
  font-weight: 600;
  cursor: pointer;
  margin-top: 12px;
}

@media (min-width: 1024px) {
  .btn,
  .btn-outline {
    padding: 18px;
  }
}

.btn:disabled,
.btn-outline:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.history {
  margin-top: 28px;
}

.history h3 {
  margin-bottom: 12px;
}

.item {
  background: #F0F5FE;
  border-radius: 16px;
  padding: 14px;
  display: flex;
  justify-content: space-between;
  margin-bottom: 12px;
}

@media (min-width: 768px) {
  .item {
    padding: 16px;
  }
}

.status {
  padding: 10px 12px;
  border-radius: 20px;
  font-size: 12px;
}

.status.hadir {
  background: #dbeafe;
  color: #1d4ed8;
}

.status.wfa {
  background: #e0e7ff;
  color: #3730a3;
}

.status.late {
  background: #fee2e2;
  color: #b91c1c;
}
</style>