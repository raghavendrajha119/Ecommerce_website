const userListContainer = document.querySelector('.user-list');
function fetchUserList(){
fetch('/admin/users', {
    method: 'GET',
})
.then(response => response.json())
.then(users => {
    users.forEach(user => {
        const userElement = document.createElement('div');
        userElement.className = 'user';

        userElement.innerHTML = `
        <div class="details">
            <p>Name: ${user.Name}</p>
            <p>Email: ${user.Email}</p>
        </div>
        <div class="buttons">
            ${user.Role == 'user' ? '<button class="make-admin" data-id="' + user.ID + '">Make Admin</button>': ''}
            ${user.Role === 'admin' ? '<button class="remove-admin" data-id="' + user.ID + '">Remove Admin</button>' : ''}
        </div>
        `;
        userListContainer.appendChild(userElement);
    });

    userListContainer.addEventListener('click', event => {
        const button = event.target;
        const userId = button.getAttribute('data-id');

        if (button.classList.contains('make-admin')) {
            RoleUser(userId, 'admin');
        } else if (button.classList.contains('remove-admin')) {
            RoleAdmin(userId, 'user');
        }
    });
})
.catch(error => {
    console.error('Error fetching user data:', error);
});
}
fetchUserList();
function RoleUser(userId, newRole) {
    fetch(`/admin/make-admin/${userId}`, {
        method: 'POST',
        headers: {
            'Content-Type':'application/json',
        },
        body: JSON.stringify({ role: newRole }),
    })
    .then(response => response.text())
    .then(message => {
        console.log(message);
        fetchUserList();
    })
    .catch(error => {
        console.error('Error updating user role:', error);
    });
}
function RoleAdmin(userId, newRole) {
    fetch(`/admin/remove-admin/${userId}`, {
        method: 'POST',
        headers: {
            'Content-Type':'application/json',
        },
        body: JSON.stringify({ role: newRole }),
    })
    .then(response => response.text())
    .then(message => {
        console.log(message);
        fetchUserList();
    })
    .catch(error => {
        console.error('Error updating user role:', error);
    });
}