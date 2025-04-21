<template>
  <div class="border border-gray-200 p-4 rounded-md lg:pr-32">
    <h2 class="text-2xl font-semibold text-gray-800">Balance Summary</h2>
    <p class="mt-4 mb-8 text-gray-700">Total:<br><span class="text-4xl font-semi-bold">₹{{ totalIncome }}</span></p>
    <p class="text-gray-700">Spent:<br><span class="text-4xl font-semi-bold">₹{{ totalExpenses }}</span></p>
  </div>
</template>

<script>
export default {
  name: 'BalanceSummary',
  props: {
    totalIncome: {
      type: Number,
      required: true,
    },
    totalExpenses: {
      type: Number,
      required: true,
    },
  },
};
</script>

<style scoped>
/* Add any component-specific styles here */
</style>
