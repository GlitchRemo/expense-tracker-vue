<template>
  <div class="add-expense p-6">
    <h1 class="text-2xl font-bold">Add Expense</h1>
    <form class="mt-4 space-y-4">
      <div>
        <label for="date" class="block text-sm font-medium">Date</label>
        <input type="date" id="date" class="mt-1 block w-full border-gray-300 rounded-md shadow-sm" />
      </div>
      <div>
        <label for="category" class="block text-sm font-medium">Category</label>
        <select id="category" class="mt-1 block w-full border-gray-300 rounded-md shadow-sm">
          <option>Groceries</option>
          <option>Transport</option>
          <option>Utilities</option>
          <option>Entertainment</option>
        </select>
      </div>
      <div>
        <label for="amount" class="block text-sm font-medium">Amount</label>
        <input type="number" id="amount" class="mt-1 block w-full border-gray-300 rounded-md shadow-sm" />
      </div>
      <button type="submit" class="bg-blue-500 text-white px-4 py-2 rounded">Add Expense</button>
    </form>
  </div>
</template>

<script>
export default {
  name: 'AddExpense',
};
</script>

<style scoped>
.add-expense {
  font-family: Arial, sans-serif;
}
</style>
