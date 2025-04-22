<template>
  <div class="view-expenses p-2 bg-gray-50 lg:p-6 mx-0 lg:mx-64 w-min md:w-auto">
  
  <!-- Navbar -->
  <Navbar/>
  <div class="add-expense max-w-xl mx-auto bg-white p-6 rounded-xl shadow-md border border-gray-200">
    <h1 class="text-2xl font-bold text-gray-800 mb-4">Add Expense</h1>
    <form class="space-y-5" @submit.prevent="submitForm">
      <!-- Date Field -->
      <div>
        <label for="date" class="block text-sm font-medium text-gray-700 mb-1">Date</label>
        <input type="date" id="date" v-model="formData.create_date"
          class="w-full border border-gray-300 rounded-md shadow-sm px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500" />
        <p v-if="formErrors.create_date" class="text-red-500 text-sm">{{ formErrors.create_date }}</p>
      </div>

      <!-- Category Field -->
      <div>
        <label for="category" class="block text-sm font-medium text-gray-700 mb-1">Category</label>
        <select id="category" v-model="formData.category"
          class="w-full border border-gray-300 rounded-md shadow-sm px-3 py-2 focus:outline-none focus:ring-2 focus:ring-blue-500">
          <option disabled selected>Select a category</option>
          <option>Groceries</option>
          <option>Transport</option>
          <option>Utilities</option>
          <option>Dining</option>
        </select>
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
  data() {
        return {
            formData: {
                create_date: '',
                category: '',
                amount: null,
                note: ''
            },
            formErrors: {
                create_date: '',
                category: '',
                amount: ''
            }
        };
    },
    methods: {
        async submitForm() {
            this.formErrors = { create_date: '', category: '', amount: '' };

            if (!this.formData.create_date) {
                this.formErrors.create_date = 'Date is required';
            }
            if (!this.formData.category) {
                this.formErrors.category = 'Category is required';
            }
            if (this.formData.amount <= 0) {
              this.formErrors.amount = 'Amount must be greater than zero';
            }
            if (!this.formData.amount) {
                this.formErrors.amount = 'Amount is required';
            }

            if (this.formErrors.create_date || this.formErrors.category || this.formErrors.amount) {
                return;
            }

            try {
                const response = await fetch('http://localhost:8080/api/expense', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(this.formData)
                });

                if (!response.ok) {
                    throw new Error('Failed to add expense');
                }

                alert('Expense added successfully');
                this.formData = { create_date: '', category: '', amount: null, note: '' };
            } catch (error) {
                console.error(error);
                alert('An error occurred while adding the expense');
            }
        }
    }
};
</script>

<style scoped>
.add-expense {
  font-family: Arial, sans-serif;
}
</style>
