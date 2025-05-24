const API_URL = import.meta.env.VITE_API_URL;

export async function deleteExpense(id, expenses, updateExpenses) {
    if (!confirm('Are you sure you want to delete this expense?')) {
        return;
    }

    try {
        const response = await fetch(`${API_URL}/api/expense?id=${id}`, {
            method: 'DELETE',
        });

        if (!response.ok) {
            throw new Error('Failed to delete expense');
        }

        const updatedExpenses = expenses.filter(item => item.id !== id);
        updateExpenses(updatedExpenses);
    } catch (error) {
        console.error('Error deleting expense:', error);
        alert('An error occurred while deleting the expense');
    }
}

export async function addExpense(newExpense) {
    const userId = parseInt(localStorage.getItem('userId'));
    try {
        const response = await fetch(`${API_URL}/api/expense`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ ...newExpense, userId }),
        });

        if (!response.ok) {
            throw new Error('Failed to add expense');
        }
    } catch (error) {
        console.error('Error adding expense:', error);
        alert('An error occurred while adding the expense');
    }
}

export async function editExpense(id, updatedData, expenses) {
    try {
        const response = await fetch(`${API_URL}/api/expense?id=${id}`, {
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

        return updatedExpenses
    } catch (error) {
        console.error('Error editing expense:', error);
        alert('An error occurred while editing the expense');
    }
}

