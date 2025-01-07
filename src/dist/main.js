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

async function handleSubmit(event) {
    event.preventDefault();
    const form = event.target;
    const formData = new FormData(form);
    const response = await fetch(form.action, {
        method: form.method,
        body: formData
    });
    const result = await response.text();
    if (result.includes('success')) {
        document.getElementById('successMessage').innerText = 'Form submitted successfully!';
        document.getElementById('successMessage').style.display = 'block';
    }
}