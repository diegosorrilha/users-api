users-api
=============

## Running
```bash
make run
```
```bash
curl http://localhost:8080/users
```

## Deployment

1 - Copy config file and set real values:
```bash
cp contrib/config.toml.sample config.toml
```

```bash
make build
```

## Changelog

[CHANGELOG.md](CHANGELOG.md)

