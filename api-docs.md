## API Documentations

**POST : api/v1/users/register**

**params :**

- name
- occupation
- email
- password

**response : **

```json
meta : {
  message: 'Your account has been created',
  code: 200,
  status: 'success'
},
data : {
  id: 1,
  name: 'Muhammad Arifin',
  occupation: 'software developer',
  email: 'arifin.muhammad@mail.com',
  token: 'tokenyangsangatpanjang'
}
```

**POST : api/v1/users/email_checkers**

**params :**

- email

**response : **

```json
meta : {
  message: 'Email address has been registered',
  code: 200,
  status: 'success'
},
data : {
  is_available: false
}
```
