users-api
=============

| METHOD | ENDPOINT    | DESCRIPTION                     |
|--------|-------------|---------------------------------|
| GET    | /users      | Get all users                   |
| POST   | /users      | Create a new user               |
| GET    | /users/{id} | Get info for a specific user    |
| PUT    | /users/{id} | Update info for a specific user |
| DELETE | /users/{id} | Delete a specific user          |


## Project structure
```bash
users-api/
├── configs -> Package responsible by load all configuration necessary to run the project
├── contrib -> Directory with project support files
├── crypt -> Package that contains encryption functions.
├── db -> Package that contains dabatase functions.
├── handlers -> Package that contains handlers functions.
├── models -> Package that contains data models functions.
├── repositories -> Package that contains repositories functions.
├── responses ->  Package that contains response functions.
└── routers -> Package that contains routes configurations.

```

## Running
### 1 - Copy config file and set real values:
```bash
cp contrib/config.toml.sample config.toml
```

### 2 - Run it:
```bash
make run
```

### 3 - Hit in the API:
#### Get all users
```bash
curl http://localhost:8000/users
```

#### Get data from a specific user
```bash
curl http://localhost:8000/users/1
```

#### Create a user
```bash
curl --location --request POST 'http://localhost:8000/users' -H 'Content-Type: application/json' --data '{
    "name": "John",
    "age": 42,
    "email": "john@gmail.com",
    "password": "4242",
    "address": "Johns street"
}'

```

#### Delete a user
```bash
curl --location --request DELETE 'http://localhost:8000/users/1'

```

#### Update a user
```bash
curl --location --request PUT 'http://localhost:8000/users/16' -H 'Content-Type: application/json' --data '{
    "name": "John4",
    "age": 42,
    "email": "john3@gmail.com",
    "password": "424242",
    "address": "Johns street"
}'
```
## Running tests

```bash
make test
```

## Running Documentation

```bash
make docs
```

## Deployment

```bash
make build
```

## Changelog

[CHANGELOG.md](CHANGELOG.md)

