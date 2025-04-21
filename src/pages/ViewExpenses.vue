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
  name: 'ExpenseList',
  data() {
    return {
      expenses: [
        { date: '2025-04-18', category: 'Groceries', amount: 1200, note: 'This is a note for the expense' },
        { date: '2025-04-19', category: 'Transport', amount: 400, note: 'This is a note for the expense' },
        { date: '2025-04-20', category: 'Utilities', amount: 800, note: 'This is a note for the expense' },
        { date: '2025-04-21', category: 'Dining', amount: 500, note: 'This is a note for the expense' },
        { date: '2025-04-20', category: 'Dining', amount: 1000, note: 'This is a note for the expense' },
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
    <Navbar/>

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

    <h1 class="text-2xl font-bold my-8">Expense List</h1>
    <ExpenseTable :expenses="expenses" />
  </div>
</template>
