function showInputField() {
    const selectedType = document.getElementById('inputType').value;
    document.getElementById('numberField').style.display = 'none';
    document.getElementById('mailField').style.display = 'none';
    document.getElementById('phoneField').style.display = 'none';

    if (selectedType === 'number') {
        document.getElementById('numberField').style.display = 'block';
    } else if (selectedType === 'mail') {
        document.getElementById('mailField').style.display = 'block';
    } else if (selectedType === 'phone') {
        document.getElementById('phoneField').style.display = 'block';
    }
}

// Listen to the form and display the response message in the same page
// after submitting the form it must view responseMessage block and hide the form block
document.getElementById('searchForm').addEventListener('submit', async (event) => {
    event.preventDefault(); // Prevent default form submission

    const form = event.target;
    const formData = new FormData(form);
    const params = new URLSearchParams(formData);

    // Disable the submit button
    const submitButton = form.querySelector('button[type="submit"]');
    submitButton.disabled = true;
    submitButton.classList.add('bg-gray-400', 'cursor-not-allowed');
    submitButton.classList.remove('bg-indigo-600', 'hover:bg-indigo-700');

    try {
        const response = await fetch(`${form.action}?${params.toString()}`, {
            method: 'GET', // No body allowed for GET
        });
        const result = await response.text();
        
        const messageElement = document.getElementById('responseMessage');
        messageElement.innerText = result;
        messageElement.style.display = 'block';
        document.getElementById('searchForm').style.display = 'none'; // Hide the form
    } catch (error) {
        console.error('Error:', error);
    }
});
