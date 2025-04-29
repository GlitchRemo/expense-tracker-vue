<template>
      <div class="min-h-screen bg-gradient-to-br from-pink-100 via-blue-100 to-purple-100 flex items-center justify-center">
    <div class="bg-white/70 backdrop-blur-lg rounded-2xl shadow-2xl p-8 w-full max-w-md">
      <div class="flex justify-between mb-6">
        <button
          @click="mode = 'login'"
          :class="['w-1/2 py-2 text-lg font-medium rounded-l-xl transition', mode === 'login' ? 'bg-purple-600 text-white' : 'bg-purple-200 text-purple-700 hover:bg-purple-300']"
        >
          Login
        </button>
        <button
          @click="mode = 'register'"
          :class="['w-1/2 py-2 text-lg font-medium rounded-r-xl transition', mode === 'register' ? 'bg-purple-600 text-white' : 'bg-purple-200 text-purple-700 hover:bg-purple-300']"
        >
          Register
        </button>
      </div>
  
        <form @submit.prevent="handleSubmit">
          <div class="space-y-4">
            <div>
              <label for="username" class="block mb-1 text-purple-700 font-medium">Username</label>
              <input type="text" id="username" v-model="username" required
                class="w-full px-4 py-2 rounded-lg border border-purple-300 focus:ring-2 focus:ring-purple-400 focus:outline-none bg-white/90 placeholder-purple-400" />
            </div>
            <div>
              <label for="password" class="block mb-1 text-purple-700 font-medium">Password</label>
              <input type="password" id="password" v-model="password" required
                class="w-full px-4 py-2 rounded-lg border border-purple-300 focus:ring-2 focus:ring-purple-400 focus:outline-none bg-white/90 placeholder-purple-400" />
            </div>
            <div v-if="mode === 'register'">
            <label for="income" class="block mb-1 text-purple-700 font-medium">Total Income</label>
            <input type="number" id="income" v-model="totalIncome" required
              class="w-full px-4 py-2 rounded-lg border border-purple-300 focus:ring-2 focus:ring-purple-400 focus:outline-none bg-white/90 placeholder-purple-400" />
          </div>
            <button type="submit"
              class="w-full py-2 bg-purple-500 hover:bg-purple-600 text-white font-semibold rounded-lg transition duration-300">
              {{ mode === 'login' ? 'Login' : 'Register' }}
            </button>
          </div>
          <p v-if="errorMessage" class="text-red-600 text-sm mt-3 text-center">{{ errorMessage }}</p>
        </form>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        username: '',
        password: '',
        errorMessage: '',
        mode: 'login' // or 'register'
      };
    },
    methods: {
      async handleSubmit() {
        const endpoint = this.mode === 'login'
          ? 'http://56.228.42.208:8080/api/auth/login'
          : 'http://56.228.42.208:8080/api/auth/register';
  
        try {
          const response = await fetch(endpoint, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ 
              username: this.username, 
              password: this.password, 
              ...(this.mode === 'register' && { totalIncome: this.totalIncome }) 
            })
          });
  
          if (!response.ok) throw await response.json();
  
          const { userId, username } = await response.json();
  
          localStorage.setItem('userId', userId);
          localStorage.setItem('username', username);
          this.$router.push('/');
        } catch (error) {
          this.errorMessage = error.message || 'Something went wrong';
        }
      }
    }
  };
  </script>
  
  <style>
  /* Pastel vibe fonts */
  body {
    font-family: 'Inter', 'Segoe UI', sans-serif;
  }
  </style>
