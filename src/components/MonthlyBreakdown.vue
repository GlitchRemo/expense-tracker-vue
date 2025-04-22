<template>
    <div class="border border-gray-200 p-4 rounded-md">
        <h2 class="text-2xl font-semibold text-gray-800 mb-2">Monthly Breakdown</h2>
        <div class="flex flex-col lg:flex-row justify-between items-center gap-4">
            <!-- Text breakdown -->
            <div class="space-y-1 text-md">
                <p class="text-green-600">🟢 Income: ₹{{ totalIncome }}</p>
                <p class="text-red-500">🔴 Expenses: ₹{{ totalExpenses }}</p>
            </div>
            <!-- Pie chart -->
            <canvas ref="pieChart" style="width: 200px;"></canvas>
        </div>
    </div>
</template>

<script>
import { Chart } from 'chart.js';
export default {
    name: 'MonthlyBreakdown',
    props: {
        totalIncome: {
            type: Number,
            required: true,
        },
        totalExpenses: {
            type: Number,
            required: true,
        },
    },
    watch: {
        totalIncome: "drawPieChart",
        totalExpenses: "drawPieChart",
    },
    methods: {
        drawPieChart() {
            if (this.chartInstance) {
                this.chartInstance.destroy();
            }
            this.chartInstance = new Chart(this.$refs.pieChart, {
                type: 'pie',
                data: {
                    labels: ['Income', 'Expenses'],
                    datasets: [
                        {
                            data: [this.totalIncome, this.totalExpenses],
                            backgroundColor: ['#16a34a', '#dc2626'],
                            borderWidth: 1,
                            hoverOffset: 10,
                        },
                    ],
                },
                options: {
                    responsive: false,
                    maintainAspectRatio: false,
                    plugins: {
                        legend: {
                            display: false,
                        },
                    },
                },
            });
        },
    },
    mounted() {
        this.drawPieChart();
    },
}
</script>

<style scoped>
/* Add any component-specific styles here */
</style>
