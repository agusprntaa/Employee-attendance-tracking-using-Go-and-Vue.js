<script setup>
import { useRouter } from 'vue-router'
import { ref, onMounted } from 'vue'

const router = useRouter()

const time = ref('')
const date = ref('')

onMounted(() => {
  const now = new Date()

  time.value = now.toLocaleTimeString('id-ID', {
    hour: '2-digit',
    minute: '2-digit'
  })

  date.value = now.toLocaleDateString('id-ID', {
    weekday: 'long',
    day: 'numeric',
    month: 'long',
    year: 'numeric'
  })
})

function goBack() {
  router.push('/employee/dashboard')
}
</script>

<template>
  <div class="wrapper">

    <img
      src="/success.png"
      alt="success"
      class="image"
    />

    <h2 class="title">Absen Sukses</h2>

    <div class="card">
      <h1>{{ time }}</h1>
      <p>{{ date }}</p>
    </div>

    <button class="btn" @click="goBack">
      Done
    </button>

  </div>
</template>

<style scoped>
.wrapper {
  min-height: 100vh;
  background: #f3f4f6;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.image {
  width: 200px;
  margin-bottom: 16px;
}

.title {
  color: #4f46e5;
  margin-bottom: 20px;
}

.card {
  width: 80%;
  background: white;
  border-radius: 30px;
  padding: 30px;
  text-align: center;
  box-shadow: 0 15px 30px rgba(0,0,0,0.1);
  margin-bottom: 24px;
}

.card h1 {
  font-size: 36px;
  color: #4f46e5;
}

.card p {
  margin-top: 8px;
  font-size: 14px;
}

.btn {
  width: 80%;
  padding: 16px;
  border-radius: 14px;
  background: #4f46e5;
  color: white;
  border: none;
  font-weight: 600;
  cursor: pointer;
}
</style>