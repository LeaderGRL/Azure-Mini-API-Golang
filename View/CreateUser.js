const CreateUserForm = document.getElementById('CreateUserForm');
const CreateUserButton = document.getElementById('CreateUserButton');
const UserDiv = document.querySelector('#Users');

// function GetUser()
// {
//     fetch('/users')
//     .then(function(response) {
//         return response.json();
//     })
//     .then(function(users) {
//         users.forEach(function(user) {
//             const UserDiv = document.createElement('div');
//             const UserP = document.createElement('p');
//             UserP.textContent = user.username;
//             UserDiv.appendChild(UserP);
//             UserDivs.appendChild(UserDiv);
//         });
//     });
// }

CreateUserForm.addEventListener('submit', function(event) {
    console.log('Form submitted!');
    event.preventDefault();
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;
    const email = document.getElementById('email').value;
    const data = {username: username, password: password, email: email, created_at: new Date("2023, 02, 02"), updated_at: new Date("2023, 02, 02")};
    console.log("DATA : " + data.username + " " + data.password + " " + data.email + " " + data.created_at + " " + data.updated_at);
    fetch('/users/create', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data),
    }) 
    .then(function(response) {
        return response.json();
    })
    // .then(function(user) {
        // const UserDiv = document.createElement('div');
        // const UserP = document.createElement('p');
        // UserP.textContent = user.username;
        // UserDiv.appendChild(UserP);
        // UserDivs.appendChild(UserDiv);
    // })
    .catch(function(error) {
        console.log(error);
    });
});

//GetUser();
