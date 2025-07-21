document.querySelector('form').addEventListener('submit', function(e) {
    e.preventDefault(); // Ngăn form submit mặc định

    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;
    const authMethod = document.getElementById('authMethod').value;

    fetch('/secure/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            username: username,
            password: password,
            auth_method: authMethod
        })
    })
    .then(response => response.json())
    .then(data => {
        alert('Server response:', data);
    })
    .catch(error => {
        alert('Error:', error);
    });
});