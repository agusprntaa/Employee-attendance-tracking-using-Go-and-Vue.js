<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useLocation } from '@/composables/useLocation'
import LocationBanner from '@/components/LocationBanner.vue'

const router = useRouter()

const note = ref('')
const error = ref('')

const {
  latitude,
  longitude,
  isInRadius,
  getCurrentLocation
} = useLocation()

onMounted(async () => {
  await getCurrentLocation()
})

function submitWFA() {
  const text = note.value?.trim()

  if (!text) {
    error.value = 'Catatan wajib diisi'
    return
  }

  if (text.length < 5) {
    error.value = 'Minimal 5 karakter'
    return
  }

  if (text.length > 200) {
    error.value = 'Maksimal 200 karakter'
    return
  }

  error.value = ''

  console.log({
    note: text,
    lat: latitude.value,
    lng: longitude.value
  })

  router.push('/employee/success?type=wfa')
}

function goBack() {
  router.back()
}
</script>

<template>
  <div class="wrapper">

    <div class="header">
      <img src="/goBack.png" class="back" @click="goBack">
    </div>

    <LocationBanner :isInRadius="isInRadius" />

    <div class="map-box">
      <iframe
        v-if="latitude && longitude"
        :src="`https://www.google.com/maps?q=${latitude},${longitude}&z=17&output=embed`"
      ></iframe>
    </div>

    <div class="form">
      <h3>Catatan/Alasan</h3>

      <textarea
        v-model="note"
        minlength="5"
        placeholder="Jelaskan alasan anda bekerja di lokasi ini..."
      ></textarea>

      <p v-if="error" class="error">{{ error }}</p>

      <p class="counter">{{ note.length }}/200</p>
    </div>

    <button
      class="btn"
      @click="submitWFA"
      :disabled="note.trim().length < 1"
    >
      Kirim
    </button>

  </div>
</template>

<style scoped>
.wrapper {
    min-height: 100vh;
    background: #f3f4f6;
    padding-bottom: 40px;
}

.header {
    height: 70px;
    background: #4f46e5;
    display: flex;
    align-items: center;
    padding: 0 16px;
}

.back {
    color: white;
    font-size: 20px;
    cursor: pointer;
}

.map-box {
    position: relative;
    margin: 20px auto;
    width: 90%;
    border-radius: 16px;
    overflow: hidden;
}

iframe {
    width: 100%;
    height: 200px;
    border: none;
}

.form {
    padding: 0 20px;
}

textarea {
    position: relative;
    width: 100%;
    height: 140px;
    border-radius: 16px;
    border: none;
    padding: 16px;
    margin-top: 10px;
    background: #e5e7eb;
    resize: none;
}

.btn {
    width: 90%;
    margin: 20px auto;
    display: block;
    padding: 16px;
    border-radius: 14px;
    background: #4f46e5;
    color: white;
    border: none;
    font-weight: 600;
}

.error {
  color: #dc2626;
  font-size: 12px;
  margin-top: 6px;
}

.counter {
  font-size: 12px;
  text-align: right;
  margin-top: 6px;
  color: #6b7280;
}

.btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
}
</style>