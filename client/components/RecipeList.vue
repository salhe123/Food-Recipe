<!-- components/RecipeList.vue -->
<script setup>
// import { recipes } from '~/data/db.json';
</script>

<template>
  <div class="container mx-auto px-4 py-8">
    <h1 class="text-4xl font-bold text-center mb-8">Our Recipes</h1>
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
      <div v-for="recipe in recipes" :key="recipe.id" class="bg-white rounded-lg shadow-lg overflow-hidden">
        <img :src="recipe.images.thumbnail" :alt="recipe.title" class="w-full h-48 object-cover">
        <div class="p-6">
          <h2 class="text-2xl font-bold mb-2">{{ recipe.title }}</h2>
          <p class="text-gray-700 mb-4">{{ recipe.description }}</p>
          <nuxt-link :to="`/recipe/${recipe.id}`" class="text-blue-500 hover:underline">View Recipe</nuxt-link>
        </div>
      </div>
    </div>
  </div>
</template>


