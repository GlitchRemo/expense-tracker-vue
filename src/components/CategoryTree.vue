<script setup lang="ts">
import { defineProps, ref, watch, computed } from 'vue';

interface CategoryNode {
  id: number;
  name: string;
  parentId: number | null;
  children: CategoryNode[];
}

const props = defineProps<{
  categories: CategoryNode[];
  selectedCategoryId: number | null;
}>();

const emit = defineEmits<{
  (e: 'select', categoryId: number): void;
}>();

const isExpanded = ref<Record<number, boolean>>({});

function toggleExpand(id: number) {
  isExpanded.value[id] = !isExpanded.value[id];
}

function selectCategory(id: number) {
  emit('select', id);
}

const isSelected = (id: number) => props.selectedCategoryId === id;
</script>

<template>
  <ul>
    <li v-for="category in categories" :key="category.id" class="category-node">
      <div
          class="flex items-center space-x-2 cursor-pointer"
          @click="selectCategory(category.id)"
          :class="{ 'font-bold text-blue-600': isSelected(category.id) }"
      >
        <button
            v-if="category.children.length"
            @click.stop="toggleExpand(category.id)"
            class="w-5 h-5 flex items-center justify-center border rounded"
            aria-label="Toggle expand"
        >
          <span v-if="isExpanded[category.id]">-</span>
          <span v-else>+</span>
        </button>
        <span v-else class="w-5"></span>
        <span>{{ category.name }}</span>
      </div>

      <CategoryTree
          v-if="category.children.length && isExpanded[category.id]"
          :categories="category.children"
          :selectedCategoryId="selectedCategoryId"
          @select="selectCategory"
      />
    </li>
  </ul>
</template>

<style scoped>
.category-node {
  margin-left: 1rem;
  list-style: none;
}
</style>
