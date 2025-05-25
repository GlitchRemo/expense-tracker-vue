import { createRouter, createWebHistory } from 'vue-router';
import Dashboard from '../pages/Dashboard.vue';
import AddExpense from '../pages/AddExpense.vue';
import ViewExpenses from '../pages/ViewExpenses.vue';
import LoginRegister from '../pages/LoginRegister.vue';
import SelectCategory from '../pages/SelectCategory.vue';

const routes = [
  { path: '/', name: 'LoginRegister', component: LoginRegister},
  { path: '/dashboard', name: 'Dashboard', component: Dashboard },
  { path: '/add-expense', name: 'AddExpense', component: AddExpense },
  { path: '/expenses', name: 'ViewExpenses', component: ViewExpenses },
  { path: '/select-category', component: SelectCategory },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
