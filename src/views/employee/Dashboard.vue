<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { computed } from 'vue'
import { useLocation } from '@/composables/useLocation'
import { useAuth } from '@/composables/useAuth'
import ProfileCard from '@/components/ProfileCard.vue'
import LocationBanner from '@/components/LocationBanner.vue'

const router = useRouter()
const { user, loadUser } = useAuth()

const canCheckIn = computed(() => {
  return isInRadius.value && !alreadyCheckedIn.value && !loading.value
})

// state
const loading = ref(false)
const currentTime = ref('')
const history = ref([])
const alreadyCheckedIn = ref(false)

// geo
const { isInRadius, getCurrentLocation } = useLocation()

// status
const locationStatus = ref('')

let interval = null

// init
onMounted(async () => {
  loadUser()
  startClock()

  try {
    const ok = await getCurrentLocation()
    locationStatus.value = ok
      ? (isInRadius.value ? 'inside' : 'outside')
      : ''

    fetchHistory()
  } catch (err) {
    console.error('ERROR:', err)
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

// HISTORY (sementara mock)
function fetchHistory() {
  history.value = [
    { id: 1, date: '2026-04-20', time: '08:45', type: 'HADIR' },
    { id: 2, date: '2026-04-19', time: '09:10', type: 'WFA' }
  ]

  const today = new Date().toLocaleDateString('id-ID')
  alreadyCheckedIn.value = history.value.some(
    h => h.date === today && h.type === 'HADIR'
  )
}

// CHECK IN
function goToScan() {
  if (loading.value || alreadyCheckedIn.value) return
  router.push('/employee/scan')
}

// WFA
function goToWFA() {
  if (loading.value) return
  router.push('/employee/wfa')
}
</script>

<template>
  <div class="container">

    <ProfileCard v-if="user" :user="user" />

     <LocationBanner :isInRadius="isInRadius" />

    <div class="clock">
      <h1>{{ currentTime }}</h1>
      <p>Waktu sekarang</p>
    </div>

    <!--
    GUARD
    <button
    class="btn"
    @click="goToScan"
    :disabled="!canCheckIn"
    >
    {{
    alreadyCheckedIn
      ? 'SUDAH ABSEN'
      : (!isInRadius
          ? 'DI LUAR RADIUS'
          : 'CHECK IN')
    }}
    </button> -->

    <button
      class="btn"
      @click="goToScan"
      :disabled="loading || alreadyCheckedIn"
    >
      {{ alreadyCheckedIn ? 'SUDAH ABSEN' : 'CHECK IN' }}
    </button>

    <button
      class="btn-outline"
      @click="goToWFA"
      :disabled="loading|| alreadyCheckedIn"
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
          <p>{{ item.time }}</p>
        </div>

        <span class="status" :class="item.type.toLowerCase()">
          {{ item.type }}
        </span>
      </div>
    </div>

  </div>
</template>

<style scoped>
.container {
  padding: 20px;
  max-width: 420px;
  margin: auto;
}

.clock {
  text-align: center;
  margin: 20px 0;
}

.clock h1 {
  font-size: 36px;
  font-weight: bold;
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

.btn:disabled,
.btn-outline:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.history {
  margin-top: 24px;
}

.item {
  background: #F0F5FE;
  border-radius: 16px;
  padding: 14px;
  display: flex;
  justify-content: space-between;
  margin-bottom: 10px;
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
</style>