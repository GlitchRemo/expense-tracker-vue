export async function deleteExpense(id, expenses, updateExpenses) {
    if (!confirm('Are you sure you want to delete this expense?')) {
        return;
    }

    try {
        const response = await fetch(`http://13.48.29.46:8080/api/expense?id=${id}`, {
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

export async function addExpense(newExpense) {
    const userId = parseInt(localStorage.getItem('userId'));
    try {
        const response = await fetch(`http://13.48.29.46:8080/api/expense`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ ...newExpense, userId }),
        });

        if (!response.ok) {
            throw new Error('Failed to add expense');
        }

        alert('Expense added successfully');
    } catch (error) {
        console.error('Error adding expense:', error);
        alert('An error occurred while adding the expense');
    }
}

export async function editExpense(id, updatedData, expenses) {
    try {
        const response = await fetch(`http://13.48.29.46:8080/api/expense?id=${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(updatedData),
        });

        if (!response.ok) {
            throw new Error('Failed to edit expense');
        }

        const updatedExpenses = expenses.map(item =>
            item.id === id ? { ...item, ...updatedData } : item
        );
        alert('Expense updated successfully');

        return updatedExpenses
    } catch (error) {
        console.error('Error editing expense:', error);
        alert('An error occurred while editing the expense');
    }
}

