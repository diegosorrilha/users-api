users-api
=============

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
#### Get all users:
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

## Deployment

```bash
make build
```

## Changelog

[CHANGELOG.md](CHANGELOG.md)

