<script setup>
import { ref } from 'vue';
import gql from 'graphql-tag';
import client from '@/apolloClient'; // Adjust the path based on your project structure

const email = ref('');
const username = ref('');
const password = ref('');

const SIGNUP_MUTATION = gql`
  mutation Signup($email: String!, $username: String!, $password: String!) {
    insert_users(objects: { email: $email, username: $username, password: $password }) {
      affected_rows
      returning {
        id
        email
        username
        password
        created_at
        updated_at
      }
    }
  }
`;

const handleSignup = async () => {
  try {
    const response = await apolloClient.mutate({
      mutation: SIGNUP_MUTATION,
      variables: {
        email: email.value,
        username: username.value,
        password: password.value,
      },
    });

    console.log('User signed up:', response.data.insert_users.returning[0]);
    // Handle success (e.g., redirect, show success message)
  } catch (error) {
    console.error('Error during signup:', error);
     // Handle error (e.g., show error message)
  }
};
</script>

<template>
  <div class="flex items-center justify-center min-h-screen bg-gray-100">
    <div class="w-full max-w-md p-8 space-y-6 bg-white shadow-lg rounded-lg">
      <h2 class="text-2xl font-bold text-center">Sign Up</h2>
      <form @submit.prevent="handleSignup">
        <div class="space-y-4">
          <div>
            <label for="username" class="block text-sm font-medium text-gray-700">Username</label>
            <input
              type="text"
              id="username"
              v-model="username"
              class="mt-1 block w-full h-10 border border-gray-300 rounded-lg shadow-sm focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
              required
            />
          </div>
          <div>
            <label for="email" class="block text-sm font-medium text-gray-700">Email</label>
            <input
              type="email"
              id="email"
              v-model="email"
              class="mt-1 block w-full h-10 border border-gray-300 rounded-lg shadow-sm focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
              required
            />
          </div>
          <div>
            <label for="password" class="block text-sm font-medium text-gray-700">Password</label>
            <input
              type="password"
              id="password"
              v-model="password"
              class="mt-1 block w-full h-10 border border-gray-300 rounded-lg shadow-sm focus:ring-blue-500 focus:border-blue-500 sm:text-sm"
              required
            />
          </div>
          <button
            type="submit"
            class="w-full py-2 px-4 bg-blue-600 text-white font-semibold rounded-lg shadow-md hover:bg-blue-700"
          >
            Sign Up
          </button>
        </div>
      </form>
      <p class="text-center text-sm text-gray-600">
        Already have an account? <nuxt-link to="/login" class="text-blue-600 hover:underline">Log in</nuxt-link>
      </p>
    </div>
  </div>
</template>
