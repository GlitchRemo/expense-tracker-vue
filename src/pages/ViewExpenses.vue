<template>
  <div class="view-expenses p-6">
    <h1 class="text-2xl font-bold">View All Expenses</h1>
    <table class="table-auto w-full mt-4 border-collapse border border-gray-300">
      <thead>
        <tr class="bg-gray-200">
          <th class="border border-gray-300 px-4 py-2">Date</th>
          <th class="border border-gray-300 px-4 py-2">Category</th>
          <th class="border border-gray-300 px-4 py-2">Amount</th>
          <th class="border border-gray-300 px-4 py-2">Action</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td class="border border-gray-300 px-4 py-2">2025-04-21</td>
          <td class="border border-gray-300 px-4 py-2">Groceries</td>
          <td class="border border-gray-300 px-4 py-2">₹1,200</td>
          <td class="border border-gray-300 px-4 py-2"><button class="text-blue-500">Edit</button></td>
        </tr>
        <tr>
          <td class="border border-gray-300 px-4 py-2">2025-04-20</td>
          <td class="border border-gray-300 px-4 py-2">Transport</td>
          <td class="border border-gray-300 px-4 py-2">₹300</td>
          <td class="border border-gray-300 px-4 py-2"><button class="text-blue-500">Edit</button></td>
        </tr>
        <tr>
          <td class="border border-gray-300 px-4 py-2">2025-04-18</td>
          <td class="border border-gray-300 px-4 py-2">Utilities</td>
          <td class="border border-gray-300 px-4 py-2">₹800</td>
          <td class="border border-gray-300 px-4 py-2"><button class="text-blue-500">Edit</button></td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  name: 'ViewExpenses',
};
</script>

<style scoped>
.view-expenses {
  font-family: Arial, sans-serif;
}
</style>
