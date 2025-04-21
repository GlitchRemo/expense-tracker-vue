<template>
  <div class="view-expenses p-2 bg-gray-50 lg:p-6 mx-0 lg:mx-64 w-min md:w-auto">
  
  <!-- Navbar -->
  <Navbar/>
  <div class="add-expense max-w-xl mx-auto bg-white p-6 rounded-xl shadow-md border border-gray-200">
    <h1 class="text-2xl font-bold text-gray-800 mb-4">Add Expense</h1>
    <form class="space-y-5">
      <!-- Date Field -->
      <div>
        <label for="date" class="block text-sm font-medium text-gray-700 mb-1">Date</label>
        <input type="date" id="date"
          class="w-full border border-gray-300 rounded-md shadow-sm px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500" />
      </div>

      <!-- Category Field -->
      <div>
        <label for="category" class="block text-sm font-medium text-gray-700 mb-1">Category</label>
        <select id="category"
          class="w-full border border-gray-300 rounded-md shadow-sm px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500">
          <option disabled selected>Select a category</option>
          <option>Groceries</option>
          <option>Transport</option>
          <option>Utilities</option>
          <option>Entertainment</option>
        </select>
      </div>

      <!-- Amount Field -->
      <div>
        <label for="amount" class="block text-sm font-medium text-gray-700 mb-1">Amount</label>
        <input type="number" id="amount"
          class="w-full border border-gray-300 rounded-md shadow-sm px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500" />
      </div>

      <!-- Note Field -->
      <div>
        <label for="note" class="block text-sm font-medium text-gray-700 mb-1">Note (optional)</label>
        <textarea id="note" rows="3" placeholder="Add a description or note..."
          class="w-full border border-gray-300 rounded-md shadow-sm px-3 py-2 resize-none focus:outline-none focus:ring-2 focus:ring-blue-500"></textarea>
      </div>

      <!-- Submit Button -->
      <div class="text-right">
        <button type="submit"
          class="bg-blue-600 hover:bg-blue-700 text-white font-semibold px-6 py-2 rounded-md shadow">
          Add Expense
        </button>
      </div>
    </form>
  </div>
 </div>
</template>

<script>
import Navbar from '../components/Navbar.vue';

export default {
  components: {
    Navbar,
  },
  name: 'AddExpense',
};
</script>

<style scoped>
.add-expense {
  font-family: Arial, sans-serif;
}
</style>
