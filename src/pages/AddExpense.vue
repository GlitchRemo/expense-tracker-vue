<template>
  <div :class="expenseToEdit ? 'p-0' : 'p-2 bg-gray-50 lg:p-6 mx-0 lg:mx-64 md:w-auto'">
    <!-- Navbar -->
    <Navbar v-if="!expenseToEdit" />
    <div class="add-expense max-w-xl mx-auto bg-white p-6 rounded-xl shadow-md border border-gray-200">
      <h1 class="text-2xl font-bold text-gray-800 mb-4">{{ expenseToEdit ? 'Edit Expense' : 'Add Expense' }}</h1>
      <form class="space-y-5" @submit.prevent="submitForm">
        <!-- Date Field -->
        <div>
          <label for="date" class="block text-sm font-medium text-gray-700 mb-1">Date</label>
          <input type="date" id="date" v-model="formData.date"
                 class="w-full border border-gray-300 rounded-md shadow-sm px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500" />
          <p v-if="formErrors.date" class="text-red-500 text-sm">{{ formErrors.date }}</p>
        </div>

        <!-- Category Display -->
        <div>
          <label class="block text-sm font-medium text-gray-700 mb-1">Category</label>
          <div class="p-2 border border-gray-300 rounded-md bg-gray-100">
            {{ formData.categoryName || 'No category selected' }}
          </div>
          <p v-if="formErrors.category" class="text-red-500 text-sm">{{ formErrors.category }}</p>
        </div>

        <!-- Amount Field -->
        <div>
          <label for="amount" class="block text-sm font-medium text-gray-700 mb-1">Amount</label>
          <input type="number" id="amount" v-model="formData.amount"
                 class="w-full border border-gray-300 rounded-md shadow-sm px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500" />
          <p v-if="formErrors.amount" class="text-red-500 text-sm">{{ formErrors.amount }}</p>
        </div>

        <!-- Note Field -->
        <div>
          <label for="note" class="block text-sm font-medium text-gray-700 mb-1">Note (optional)</label>
          <textarea id="note" v-model="formData.note" rows="3" placeholder="Add a description or note..."
                    class="w-full border border-gray-300 rounded-md shadow-sm px-3 py-2 resize-none focus:outline-none focus:ring-2 focus:ring-blue-500"></textarea>
        </div>

        <!-- Submit Button -->
        <div class="text-right">
          <button type="submit"
                  class="bg-blue-600 hover:bg-blue-700 text-white font-semibold px-6 py-2 rounded-md shadow">
            {{ expenseToEdit ? 'Edit Expense' : 'Add Expense' }}
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import Navbar from '../components/Navbar.vue';
import { addExpense, editExpense } from '../utils/expenseUtils';

const API_URL = import.meta.env.VITE_API_URL;

export default {
  components: { Navbar },
  name: 'AddExpense',
  props: {
    expenseToEdit: {
      type: Object,
      default: null,
    },
    expenses: {
      type: Array,
      required: false,
    },
    updateExpenses: {
      type: Function,
      required: false,
    },
  },
  data() {
    return {
      categories: [],
      formData: {
        date: this.getTodayDate(),
        categoryId: null,
        categoryName: '',
        amount: null,
        note: '',
        userId: localStorage.getItem('userId') || null,
      },
      formErrors: {
        date: '',
        category: '',
        amount: ''
      },
    };
  },
  watch: {
    expenseToEdit: {
      immediate: true,
      handler(newExpense) {
        if (newExpense) {
          this.formData = {
            ...newExpense,
            categoryId: newExpense.categoryId,
            categoryName: '', // will be set after categories load
            date: newExpense.date || this.getTodayDate(),
          };
          this.setCategoryNameById(newExpense.categoryId);
        }
      },
    },
  },
  mounted() {
    this.fetchCategories();
  },
  methods: {
    getTodayDate() {
      const today = new Date();
      return today.toISOString().slice(0, 10);
    },

    async fetchCategories() {
      try {
        // Replace the URL below with your real API endpoint
        const res = await fetch(`${API_URL}/api/categories`);
        if (!res.ok) throw new Error('Failed to fetch categories');
        this.categories = await res.json();

        // If categoryId is present in route query, set it
        const catId = Number(this.$route.query.categoryId);
        if (catId) {
          this.formData.categoryId = catId;
          this.setCategoryNameById(catId);
        }

        // If editing expense and categoryName not set yet, set it
        if (this.expenseToEdit && this.formData.categoryId) {
          this.setCategoryNameById(this.formData.categoryId);
        }

      } catch (error) {
        console.error('Error fetching categories:', error);
      }
    },

    findCategoryById(categories, id) {
      for (const category of categories) {
        if (category.id === id) {
          return category;
        }
        if (category.children && category.children.length > 0) {
          const found = this.findCategoryById(category.children, id);
          if (found) return found;
        }
      }
      return null;
    },

    setCategoryNameById(id) {
      const cat = this.findCategoryById(this.categories, id);
      if (cat) {
        this.formData.categoryName = cat.name;
        this.formData.categoryId = cat.id;
        this.formErrors.category = '';
      } else {
        this.formData.categoryName = '';
        this.formErrors.category = 'Category not found';
      }
    },

    async submitForm() {
      this.formErrors = { date: '', category: '', amount: '' };

      if (!this.formData.date) {
        this.formErrors.date = 'Date is required';
      }
      if (!this.formData.categoryId) {
        this.formErrors.category = 'Category is required';
      }
      if (!this.formData.amount || this.formData.amount <= 0) {
        this.formErrors.amount = 'Amount must be greater than zero';
      }

      if (this.formErrors.date || this.formErrors.category || this.formErrors.amount) {
        return;
      }

      const payload = {
        date: this.formData.date,
        categoryId: this.formData.categoryId,
        amount: this.formData.amount,
        note: this.formData.note,
        userId: this.formData.userId,
      };

      if (this.expenseToEdit) {
        const updatedExpenses = await editExpense(this.expenseToEdit.id, payload, this.expenses);
        if (this.updateExpenses) {
          this.updateExpenses(updatedExpenses);
        }
        this.$emit('edit-expense', payload);
      } else {
        await addExpense(payload);
        this.formData = {
          date: this.getTodayDate(),
          categoryId: null,
          categoryName: '',
          amount: null,
          note: '',
          userId: this.formData.userId,
        };
        this.formErrors = { date: '', category: '', amount: '' };
        this.$router.push('/expenses'); // Redirect to view expenses page
      }
    },
  },
};
</script>

<style scoped>
.add-expense {
  font-family: Arial, sans-serif;
}
</style>
