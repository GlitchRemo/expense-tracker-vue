import { createRouter, createWebHistory } from 'vue-router';
import Dashboard from '../pages/Dashboard.vue';
import AddExpense from '../pages/AddExpense.vue';
import ViewExpenses from '../pages/ViewExpenses.vue';

const routes = [
  { path: '/', name: 'Dashboard', component: Dashboard },
  { path: '/add-expense', name: 'AddExpense', component: AddExpense },
  { path: '/expenses', name: 'ViewExpenses', component: ViewExpenses },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
