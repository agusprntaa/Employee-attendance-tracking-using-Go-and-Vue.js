import { ref } from 'vue'
import { getDistance } from '@/utils/geo'

export function useLocation() {
  const latitude = ref(null)
  const longitude = ref(null)
  const distance = ref(0)
  const isInRadius = ref(false)
  const error = ref('')
  const loading = ref(false)

  // default (nanti bisa dari backend)
  const officeLat = -8.65
  const officeLng = 115.216
  const radius = 100

  function getCurrentLocation() {
    loading.value = true
    error.value = ''

    return new Promise((resolve) => {
      navigator.geolocation.getCurrentPosition(
        (pos) => {
          latitude.value = pos.coords.latitude
          longitude.value = pos.coords.longitude

          distance.value = getDistance(
            officeLat,
            officeLng,
            latitude.value,
            longitude.value
          )

          isInRadius.value = distance.value <= radius

          loading.value = false
          resolve(true)
        },
        () => {
          error.value = 'Gagal mengambil lokasi'
          loading.value = false
          resolve(false)
        },
        { timeout: 5000 }
      )
    })
  }

  return {
    latitude,
    longitude,
    distance,
    isInRadius,
    error,
    loading,
    getCurrentLocation
  }
}