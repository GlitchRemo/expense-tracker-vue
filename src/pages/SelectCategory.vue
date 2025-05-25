<script setup lang="ts">
import { ref, onMounted } from 'vue';
import CategoryTree from '../components/CategoryTree.vue';
import { useRouter } from 'vue-router';

const API_URL = import.meta.env.VITE_API_URL;

interface CategoryNode {
  id: number;
  name: string;
  parentId: number | null;
  children: CategoryNode[];
}

const router = useRouter();
const categories = ref<CategoryNode[]>([]);
const selectedCategoryId = ref<number | null>(null);

async function fetchCategories() {
  const res = await fetch(`${API_URL}/api/categories`);
  categories.value = await res.json();
}

function onCategorySelect(id: number) {
  selectedCategoryId.value = id;
  router.push({ name: 'AddExpense', query: { categoryId: id } });
}

onMounted(fetchCategories);
</script>


<template>
  <div>
    <h2>Select a Category to Add Expense</h2>
    <CategoryTree
        :categories="categories"
        :selectedCategoryId="selectedCategoryId"
        @select="onCategorySelect"
    />
  </div>
</template>
