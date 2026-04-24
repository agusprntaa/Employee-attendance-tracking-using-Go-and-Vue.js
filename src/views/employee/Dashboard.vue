<script setup>
import { ref, onMounted } from 'vue'
import { useLocation } from '@/composables/useLocation'

const {
  isInRadius,
  getCurrentLocation
} = useLocation()

// status banner
const locationStatus = ref('') 
// 'inside' | 'outside'

// jalan saat halaman dibuka
onMounted(async () => {
  const ok = await getCurrentLocation()
  if (!ok) return

  locationStatus.value = isInRadius.value ? 'inside' : 'outside'
})
</script>

<template>
  <div class="container">

    <div
      v-if="locationStatus"
      class="banner"
      :class="locationStatus"
    >
      {{
        locationStatus === 'inside'
          ? 'Anda berada dalam radius kantor'
          : 'Anda berada di luar radius kantor'
      }}
    </div>

  </div>
</template>

<style scoped>
.container {
  padding: 20px;
}
.banner {
  width: 100%;
  max-width: 400px;
  margin: auto;
  padding: 16px;
  border-radius: 20px;
  text-align: center;
  font-size: 14px;
}
.banner.outside {
  background: #f3ead7;
  color: #7c6f57;
}
.banner.inside {
  background: #e6f4ea;
  color: #2e7d32;
}
</style>