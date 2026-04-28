<script setup>
defineProps({
  isInRadius: Boolean
})
</script>

<template>
  <div
    class="banner"
    :class="isInRadius ? 'inside' : 'outside'"
  >
    {{
      isInRadius
        ? 'Anda berada dalam radius kantor'
        : 'Anda berada di luar radius kantor'
    }}
  </div>
</template>

<style scoped>
.banner {
  margin: 20px;
  padding: 16px;
  border-radius: 16px;
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