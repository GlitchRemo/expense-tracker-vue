<template>
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
</template>

<script>
export default {
  name: 'ExpenseTable',
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
    },
    methods: {
        deleteExpense(id) {
            if (this.onDeleteExpense) {
                this.onDeleteExpense(id);
            }
        },
        editExpense(item) {
            const updatedData = { ...item };
            const newCategory = prompt('Edit Category:', item.category);
            const newAmount = prompt('Edit Amount:', item.amount);
            const newNote = prompt('Edit Note:', item.note || '');

            if (newCategory !== null) updatedData.category = newCategory;
            if (newAmount !== null) updatedData.amount = parseInt(newAmount, 10);
            if (newNote !== null) updatedData.note = newNote;

            this.$emit('edit-expense', item.id, updatedData);
        },
    },
};
</script>

<style scoped>
/* Add any component-specific styles here */
</style>
