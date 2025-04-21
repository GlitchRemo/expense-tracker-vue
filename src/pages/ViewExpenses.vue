<script>
import { Chart, PieController, ArcElement, Tooltip, Legend } from 'chart.js';

Chart.register(PieController, ArcElement, Tooltip, Legend);

export default {
  name: 'ExpenseList',
  data() {
    return {
      expenses: [
        { date: '2025-04-18', category: 'Groceries', amount: 1200, note: 'This is a note for the expense' },
        { date: '2025-04-19', category: 'Transport', amount: 400, note: 'This is a note for the expense' },
        { date: '2025-04-20', category: 'Utilities', amount: 800, note: 'This is a note for the expense' },
        { date: '2025-04-21', category: 'Dining', amount: 500, note: 'This is a note for the expense' },
      ],
      categoryColors: {
        Groceries: '#f87171', // Red
        Transport: '#60a5fa', // Blue
        Utilities: '#34d399', // Green
        Dining: '#fbbf24',    // Yellow
      },
    };
  },

  mounted() {
    const categorySums = {};
    this.expenses.forEach(exp => {
      categorySums[exp.category] = (categorySums[exp.category] || 0) + exp.amount;
    });

    const categories = Object.keys(categorySums);
    const values = Object.values(categorySums);
    const colors = categories.map(cat => this.categoryColors[cat] || '#ccc');

    new Chart(this.$refs.expenseChart, {
      type: 'pie',
      data: {
        labels: categories,
        datasets: [
          {
            data: values,
            backgroundColor: colors,
            borderWidth: 1,
            hoverOffset: 10,
          },
        ],
      },
      options: {
        responsive: false,
        plugins: {
          legend: {
            display: false, // disable built-in legend
          },
        },
      },
    });
  },
};
</script>

<template>
  <div class="view-expenses p-2 bg-gray-50 lg:p-6 mx-0 lg:mx-64 w-min md:w-auto">
    <!-- Navbar -->
    <nav class="bg-white shadow rounded-md p-4 mb-6 flex justify-between items-center">
      <span class="text-xl font-semibold text-gray-800">Expense Tracker</span>
      <button class="text-gray-600 hover:text-red-500">Log out</button>
    </nav>

    <!-- Pie Chart -->
    <div class="flex flex-col items-center gap-8 bg-white p-6 rounded-xl shadow-sm border border-gray-200">
      <h2 class="text-lg font-semibold mb-2">Expenses by Category</h2>
      <canvas ref="expenseChart" width="400" height="400" class="mb-4"></canvas>

      <!-- Custom Legend -->
      <div class="flex flex-wrap gap-4">
        <div v-for="(color, category) in categoryColors" :key="category" class="flex items-center gap-2">
          <span :style="{ backgroundColor: color }" class="w-4 h-4 rounded-full inline-block"></span>
          <span class="text-sm">{{ category }}</span>
        </div>
      </div>
    </div>

    <h1 class="text-2xl font-bold">Expense List</h1>
    <table class="table-auto text-left w-full mt-4 border-collapse border border-gray-300">
      <thead>
        <tr class="bg-gray-200">
          <th class="border-b border-gray-300 px-4 py-2 min-w-[125px]">Date</th>
          <th class="border-b border-gray-300 px-4 py-2 min-w-[200px]">Category</th>
          <th class="border-b border-gray-300 px-4 py-2 min-w-[125px]">Amount</th>
          <th class="border-b border-gray-300 px-4 py-2 min-w-[125px]">Action</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="expense in expenses" :key="expense.id">
          <td class="border-b border-gray-300 px-4 py-2">{{ expense.date }}</td>
          <td class="border-b border-gray-300 px-4 py-2">{{ expense.category }}
            <div v-if="expense.note" class="group relative inline-block ml-2">
              <i class="fas fa-info-circle text-gray-500 cursor-pointer"></i>
              <div
                class="absolute z-10 hidden group-hover:block bg-gray-700 text-white text-sm px-3 py-2 rounded w-max max-w-xs whitespace-normal shadow-lg bottom-full left-1/2 -translate-x-1/2 mb-2">
                {{ expense.note }}
              </div>
            </div>
          </td>
          <td class="border-b border-gray-300 px-4 py-2">₹{{ expense.amount }}</td>
          <td class="flex gap-6 px-4 py-2 border-b">
            <button class="text-blue-600 hover:underline">
              <i class="fas fa-edit"></i>
            </button>
            <button class="text-red-600 hover:underline">
              <i class="fas fa-trash"></i>
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
