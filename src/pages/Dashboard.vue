<script>
import { Chart, PieController, ArcElement, Tooltip, Legend } from 'chart.js';
import ExpenseTable from '../components/ExpenseTable.vue';
import Navbar from '../components/Navbar.vue';

Chart.register(PieController, ArcElement, Tooltip, Legend);

export default {
  components: {
    Navbar,
    ExpenseTable,
  },
  name: 'Dashboard',
  data() {
    return {
      expenses: [
        { date: '2025-04-18', category: 'Groceries', amount: 1200, note: 'This is a note for the expense' },
        { date: '2025-04-19', category: 'Transport', amount: 400, note: 'This is a note for the expense' },
        { date: '2025-04-20', category: 'Utilities', amount: 800, note: 'This is a note for the expense' },
        { date: '2025-04-21', category: 'Dining', amount: 500, note: 'This is a note for the expense' },
        { date: '2025-04-20', category: 'Dining', amount: 1000, note: 'This is a note for the expense' },
      ],
      totalIncome: 10000,
      totalExpenses: 5000,
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
              data: [this.totalIncome, this.totalExpenses],
              backgroundColor: ['#16a34a', '#dc2626'],
              borderWidth: 1,
              hoverOffset: 10,
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

<template>
  <div class="min-h-screen bg-gray-50 p-6 font-sans mx-0 lg:mx-64">
    <!-- Navbar -->
    <Navbar />

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
            <p class="text-green-600">🟢 Income: ₹{{ this.totalIncome }}</p>
            <p class="text-red-500">🔴 Expenses: ₹{{ this.totalExpenses }}</p>
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

    <!-- Expense List -->
    <div class="mt-8 border border-gray-200 bg-white p-6 rounded-xl shadow-sm">
      <h1 class="text-2xl font-bold mb-4">Expense List</h1>
      <ExpenseTable :expenses="expenses" />
    </div>
  </div>
</template>
