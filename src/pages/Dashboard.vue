<script>
import ExpenseTable from '../components/ExpenseTable.vue';
import Navbar from '../components/Navbar.vue';
import BalanceSummary from '../components/BalanceSummary.vue';
import MonthlyBreakdown from '../components/MonthlyBreakdown.vue';
import { deleteExpense, editExpense } from '../utils/expenseUtils';

const API_URL = import.meta.env.VITE_API_URL;

export default {
  components: {
    Navbar,
    ExpenseTable,
    BalanceSummary,
    MonthlyBreakdown,
  },
  name: 'Dashboard',
  data() {
    return {
      id: localStorage.getItem('userId') || null,
      username: localStorage.getItem('username') || '',
      expenses: [],
      totalIncome: 0,
      totalExpenses: 0,
    };
  },
  methods: {
    navigateToSelectCategory() {
      this.$router.push('/select-category');
    },
    navigateToViewExpenses() {
      this.$router.push('/expenses');
    },
    async fetchDashboardData() {
      try {
        const response = await fetch(`${API_URL}/api/dashboard?userID=${this.id}`);
        const data = await response.json();
        this.expenses = data.monthlyBreakdown;
        this.totalIncome = data.totalIncome;
        this.totalExpenses = data.totalExpenses;
      } catch (error) {
        console.error('Error fetching dashboard data:', error);
      }
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
      this.fetchDashboardData();
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
        <h1 class="text-2xl font-bold text-gray-800">Hello, {{username}}! 👋</h1>
        <p class="text-gray-600">Welcome back: Here’s your expense summary:</p>
      </div>
      <button @click="navigateToSelectCategory"
        class="bg-blue-600 text-white px-4 py-2 rounded-md shadow hover:bg-blue-700 transition">
        Add
      </button>
    </div>

    <!-- Summary Cards Grouped Box -->
    <div class="flex flex-wrap gap-12 justify-between border border-gray-200 bg-white p-4 rounded-xl shadow-sm mb-6">
      <!-- Balance Summary -->
      <BalanceSummary :totalIncome="totalIncome" :totalExpenses="totalExpenses" />

      <!-- Monthly Breakdown with Pie Chart -->
      <MonthlyBreakdown :totalIncome="totalIncome" :totalExpenses="totalExpenses" />
    </div>

    <!-- View All Expenses Link -->
    <a class="text-blue-600 mt-4 block hover:underline" @click="navigateToViewExpenses">
      View All Expenses
    </a>

    <!-- Expense List -->
    <div class="mt-8 border border-gray-200 bg-white p-6 rounded-xl shadow-sm">
      <h1 class="text-2xl font-bold mb-4">Expense List</h1>
      <ExpenseTable 
        :expenses="expenses" 
        :onDeleteExpense="deleteExpenseHandler" 
        :onEditExpense="editExpenseHandler" 
        :updateExpenses="updateExpensesHandler" 
      />
    </div>
  </div>
</template>
