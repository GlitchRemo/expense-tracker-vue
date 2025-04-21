<template>
  <div class="min-h-screen bg-gray-50 p-6 font-sans sm:mx-0 lg:mx-64">
    <!-- Navbar -->
    <nav class="bg-white shadow rounded-md p-4 mb-6 flex justify-between items-center">
      <span class="text-xl font-semibold text-gray-800">Expense Tracker</span>
      <button class="text-gray-600 hover:text-red-500">Log out</button>
    </nav>

    <!-- Greeting and Add Button -->
    <div class="flex justify-between items-start mb-4">
      <div>
        <h1 class="text-2xl font-bold text-gray-800">Hi, Riya 👋</h1>
        <p class="text-gray-600">Welcome back: Here’s your expense summary:</p>
      </div>
      <button @click="navigateToAddExpense"
        class="bg-blue-600 text-white px-4 py-2 rounded-md shadow hover:bg-blue-700 transition">
        Add
      </button>
    </div>

    <!-- Summary Cards Grouped Box -->
    <div class="flex flex-wrap gap-12 justify-between border border-gray-200 bg-white p-4 rounded-xl shadow-sm mb-6">
      <!-- Balance Summary -->
      <div class="border border-gray-200 p-4 rounded-md lg:pr-32">
        <h2 class="text-2xl font-semibold text-gray-800">Balance Summary</h2>
        <p class="mt-4 mb-8 text-gray-700">Total:<br><span class="text-4xl font-semi-bold">₹10,000</span></p>
        <p class="text-gray-700">Spent:<br><span class="text-4xl font-semi-bold">₹5,000</span></p>
      </div>

      <!-- Monthly Breakdown with Pie Chart -->
      <div class="border border-gray-200 p-4 rounded-md">
        <h2 class="text-2xl font-semibold text-gray-800 mb-2">Monthly Breakdown</h2>
        <div class="flex flex-col lg:flex-row justify-between items-center gap-4">
          <!-- Text breakdown -->
          <div class="space-y-1 text-md">
            <p class="text-green-600">🟢 Income: ₹10,000</p>
            <p class="text-red-500">🔴 Expenses: ₹5,000</p>
          </div>
          <!-- Pie chart -->
          <canvas ref="pieChart" style="width: 200px;"></canvas>
        </div>
      </div>
    </div>

    <!-- View All Expenses Link -->
    <a href="/expenses" class="text-blue-600 mt-4 block hover:underline">
      View All Expenses
    </a>

    <!-- Expense Table -->
    <div class="bg-white mt-6 p-4 rounded-md shadow border border-gray-200">
      <h2 class="text-lg font-semibold text-gray-800 mb-2">Expense List</h2>
      <div class="overflow-x-auto">
        <table class="w-full text-left border-collapse">
          <thead>
            <tr class="text-gray-700 bg-gray-100">
              <th class="px-4 py-2 border-b">Date</th>
              <th class="px-4 py-2 border-b">Category</th>
              <th class="px-4 py-2 border-b">Amount</th>
              <th class="px-4 py-2 border-b">Action</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="item in expenses" :key="item.id" class="hover:bg-gray-50">
              <td class="px-4 py-2 border-b">{{ item.date }}</td>
              <td class="px-4 py-2 border-b">{{ item.category }}</td>
              <td class="px-4 py-2 border-b">₹{{ item.amount }}</td>
              <td class="px-4 py-2 border-b">
                <button class="text-blue-600 hover:underline mr-4 lg:mr-8">
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
    </div>
  </div>
</template>

<script>
import { Chart, PieController, ArcElement, Tooltip, Legend } from 'chart.js';

Chart.register(PieController, ArcElement, Tooltip, Legend);

export default {
  name: 'Dashboard',
  data() {
    return {
      expenses: [
        { id: 1, date: '2025-04-21', category: 'Groceries', amount: 1200 },
        { id: 2, date: '2025-04-20', category: 'Transport', amount: 300 },
        { id: 3, date: '2025-04-18', category: 'Utilities', amount: 800 },
      ],
    };
  },
  methods: {
    navigateToAddExpense() {
      this.$router.push('/add-expense');
    },
    drawPieChart() {
      new Chart(this.$refs.pieChart, {
        type: 'pie',
        data: {
          labels: ['Income', 'Expenses'],
          datasets: [
            {
              data: [10000, 5000],
              backgroundColor: ['#16a34a', '#dc2626'],
              borderWidth: 0,
            },
          ],
        },
        options: {
          responsive: false,
          maintainAspectRatio: false,
          plugins: {
            legend: {
              display: false,
            },
          },
        },
      });
    },
  },
  mounted() {
    this.drawPieChart();
  },
};
</script>
