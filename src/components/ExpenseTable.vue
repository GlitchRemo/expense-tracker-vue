<template>
  <div>
    <!-- Existing table structure -->
    <div class="overflow-x-auto">
      <table class="w-full text-left border-collapse">
        <thead>
          <tr class="text-gray-700 bg-gray-100">
            <th class="px-4 py-2 border-b min-w-[125px]">Date</th>
            <th class="px-4 py-2 border-b min-w-[200px]">Category</th>
            <th class="px-4 py-2 border-b min-w-[125px]">Amount</th>
            <th class="px-4 py-2 border-b min-w-[125px]">Action</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in expenses" :key="item.id" class="hover:bg-gray-50">
            <td class="px-4 py-2 border-b">{{ item.date }}</td>
            <td class="px-4 py-2 border-b">
              {{ item.category }}
              <div v-if="item.note" class="group relative inline-block ml-2">
                <i class="fas fa-info-circle text-gray-500 cursor-pointer"></i>
                <div
                  class="absolute z-10 hidden group-hover:block bg-gray-700 text-white text-sm px-3 py-2 rounded w-max max-w-xs whitespace-normal shadow-lg bottom-full left-1/2 -translate-x-1/2 mb-2">
                  {{ item.note }}
                </div>
              </div>
            </td>
            <td class="px-4 py-2 border-b">₹{{ item.amount }}</td>
            <td class="flex gap-4 px-4 py-2 border-b">
              <button class="text-blue-600 hover:underline" @click="editExpense(item)">
                <i class="fas fa-edit"></i>
              </button>
              <button class="text-red-600 hover:underline" @click="deleteExpense(item.id)">
                <i class="fas fa-trash"></i>
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <!-- AddExpense Modal -->
    <div v-if="editingExpense" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center">
      <div class="bg-white p-6 pt-12 rounded-lg shadow-lg w-full max-w-lg relative">
        <button @click="closeEditModal" class="absolute top-3 right-4 pt-1 text-gray-500 hover:text-gray-700">
          <i class="fas fa-times text-xl"></i>
        </button>
        <AddExpense :expense-to-edit="editingExpense" 
          :expenses="expenses"
          :update-expenses="updateExpenses" 
          @edit-expense="closeEditModal"
          @add-expense="closeEditModal" />
      </div>
    </div>
  </div>
</template>

<script>
import AddExpense from '../pages/AddExpense.vue';

export default {
  name: 'ExpenseTable',
  components: {
    AddExpense,
  },
  props: {
    expenses: {
      type: Array,
      required: true,
    },
    onDeleteExpense: {
      type: Function,
      required: false,
    },
    onEditExpense: {
      type: Function,
      required: false,
    },
    updateExpenses: {
      type: Function,
      required: false,
    },
  },
  data() {
    return {
      editingExpense: null, // Tracks the expense being edited
    };
  },
  methods: {
    deleteExpense(id) {
      if (this.onDeleteExpense) {
        this.onDeleteExpense(id);
      }
    },
    editExpense(item) {
      this.editingExpense = { ...item };
    },
    closeEditModal() {
      this.editingExpense = null;
    },
  },
};
</script>

<style scoped>
/* Add any component-specific styles here */
</style>
