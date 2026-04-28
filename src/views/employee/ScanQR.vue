<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { BrowserMultiFormatReader } from '@zxing/browser';
import { useLocation } from '@/composables/useLocation'
import { checkInAPI} from '@/services/attendance'

const router = useRouter()

//state
const loading = ref(false)
const error = ref('')
const scanned = ref(false)

const { isInRadius, getCurrentLocation } = useLocation()

let codeReader = null

//init
onMounted(async () => {
    const ok = await getCurrentLocation()
    if(!ok) {
        error.value = 'Gagal mengambil lokasi'
        return
    }
    startScanner()
})

onUnmounted(() => {
    stopScanner()
})

//start scan
async function startScanner() {
    codeReader = new BrowserMultiFormatReader()

    try {
        await codeReader.decodeFromVideoDevice(
            null,
            document.getElementById('video'),
            (result, err) => {
                if (result) {
                    handleScan(result.getText())
                }
            }
        )
    } catch (err) {
        console.error (err) 
        error.value = 'Gagal mengakses kamera'
    }
}

//stop scan
function stopScanner() {
    if(codeReader) {
        codeReader.reset()
    }
}

//handle scan
async function handleScan(decodedText) {
    if (scanned.value) return
    scanned.value = true
    loading.value = true

    stopScanner()

    try {
        if (!isInRadius.value) {
            alert('Anda diluar radius kantor')
            return router.push('/employee/dashboard')
        }
        await checkInAPI( {
            qr_code: decodedText
        }) 
        router.push('/employee/success')
    } catch (err) {
        console.error(err)
        alert('Check In gagal')
        router.push('/employee/dashboard')
    } finally {
        loading.value = false
    }
}
//back
function goBack() {
    router.back()
}
</script>

<template>
    <div class="wrapper">

        <div class="header">
        <img src="/public/goBack.png"
        class="back"
        @click="goBack">
        </div>
        
        <div v-if="error" class="error">
            {{ error }}
        </div>

        <div class="scan-box">
            <video id="video" autoplay muted playsinlin></video>

            <div class="frame"></div>
        </div>

        <div class="scan-status">
            <span class="spinner">⟳</span>
            <span>{{ loading? 'memproses...' :'scanning...' }}</span>
        </div>
    </div>

</template>

<style scoped>
.wrapper {
    min-height: 100vh;
    background: #f3f4f6;
    display: flex;
    flex-direction: column;
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

.error {
    background: #fee2e2;
    color: #b91c1c;
    padding: 10px;
    margin: 10px;
    border-radius: 10px;
    text-align: center;
}

.scan-box {
    position: relative;
    margin: 40px auto;
    width: 90%;
    max-width: 420px;
    height: 500px;
    border-radius: 20px;
    overflow: hidden;
    background: black;
}

video {
    width: 100%;
    height: 100%;
    object-fit: cover;
}

.frame {
    position: absolute;
    width: 220px;
    height: 220px;
    border: 3px solid white;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    border-radius: 12px;
}

.scan-status {
    width: 70%;
    margin: auto;
    margin-top: 40px;
    padding: 14px;
    background: white;
    border-radius: 12px;
    display: flex;
    justify-content: center;
    gap: 10px;
    box-shadow: 0 10px 20px rgba(0,0,0,0.1);
    font-size: 14px;
}

.spinner {
    animation: spin 1s linear infinite;
}

@keyframes spin {
    100% {
        transform: rotate(360deg);
    }
}
</style>