<template>
  <div class="relative min-h-screen bg-cover bg-center flex items-center justify-center">
    <!-- Background Image -->
    <img src="/assets/css/image/image.png" alt="Background" class="absolute inset-0 w-full h-full object-cover z-0" />

    <!-- Content Box -->
    <div class="relative z-10 bg-white bg-opacity-80 p-8 rounded-lg shadow-lg text-center max-w-lg">
      <h1 class="text-5xl font-bold text-gray-800 mb-4">Welcome to Our Food Recipe Website</h1>
      <p class="text-xl text-gray-600 mb-6">Your journey to great experiences starts here.</p>
      <nuxt-link to="/login">
        <button class="bg-blue-500 text-white px-6 py-3 rounded-md hover:bg-blue-600 transition duration-300">
          Get Started
        </button>
      </nuxt-link>
    </div>
  </div>
</template>

<style scoped>
/* Additional styles if needed */
</style>
