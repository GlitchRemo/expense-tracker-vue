<script>
import { Chart, PieController, ArcElement, Tooltip, Legend } from 'chart.js';
import ExpenseTable from '../components/ExpenseTable.vue';
import Navbar from '../components/Navbar.vue';
import { deleteExpense, editExpense } from '../utils/expenseUtils';

Chart.register(PieController, ArcElement, Tooltip, Legend);

export default {
  components: {
    Navbar,
    ExpenseTable,
  },
  name: 'ExpenseList',
  data() {
    return {
      id: localStorage.getItem('userId') || null,
      expenses: [],
      categoryColors: {
        Groceries: '#f87171', // Red
        Transport: '#60a5fa', // Blue
        Utilities: '#34d399', // Green
        Dining: '#fbbf24',    // Yellow
      },
    };
  },

  methods: {
    async fetchExpenses() {
      try {
        const response = await fetch(`http://56.228.42.208:8080/api/dashboard?userID=${this.id}`);
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        const data = await response.json();
        this.expenses = data.monthlyBreakdown;

        // Render the pie chart after data is fetched
        this.renderPieChart();
      } catch (error) {
        console.error('Error fetching expenses:', error);
      }
    },
    
    renderPieChart() {
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

    async deleteExpenseHandler(id) {
      await deleteExpense(id, this.expenses, updatedExpenses => {
        this.expenses = updatedExpenses;
      });
    },

    async editExpenseHandler(id, updatedData) {
      await editExpense(id, updatedData, this.expenses, updatedExpenses => {
        this.expenses = updatedExpenses;
      });
    },

    updateExpensesHandler(updatedExpenses) {
      this.expenses = updatedExpenses;
    },
  },

  created() {
    this.fetchExpenses();
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
    <ExpenseTable 
      :expenses="expenses" 
      :onDeleteExpense="deleteExpenseHandler" 
      :onEditExpense="editExpenseHandler" 
      :updateExpenses="updateExpensesHandler" 
    />
  </div>
</template>
