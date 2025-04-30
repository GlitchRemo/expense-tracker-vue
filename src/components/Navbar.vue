<template>
    <nav class="bg-white shadow rounded-md p-4 mb-6 flex justify-between items-center">
        <span class="text-xl font-semibold text-gray-800" @click="navigateToHome">Expense Tracker</span>
        <button class="text-gray-600 hover:text-red-500" @click="logout">
            <i class="fas fa-sign-out-alt"></i>
        </button>
    </nav>
</template>

<script>
export default {
    name: 'Navbar',
    methods: {
        navigateToHome() {
            this.$router.push('/dashboard');
        },
        logout() {
            // Clear localStorage and redirect to login page
            localStorage.clear();
            this.$router.push('/login');
        },
    }
}
</script>

<style scoped>
/* Add any component-specific styles here */
span {
  cursor: pointer;
  transition: color 0.3s;
}
</style>