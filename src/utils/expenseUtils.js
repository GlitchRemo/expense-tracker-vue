export async function deleteExpense(id, expenses, updateExpenses) {
    if (!confirm('Are you sure you want to delete this expense?')) {
        return;
    }

    try {
        const response = await fetch(`http://localhost:8080/api/expense?id=${id}`, {
            method: 'DELETE',
        });

        if (!response.ok) {
            throw new Error('Failed to delete expense');
        }

        const updatedExpenses = expenses.filter(item => item.id !== id);
        updateExpenses(updatedExpenses);
        alert('Expense deleted successfully');
    } catch (error) {
        console.error('Error deleting expense:', error);
        alert('An error occurred while deleting the expense');
    }
}
