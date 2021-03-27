# gqlgen-example

## Start 🌟

#### Run Docker 🚀
```
  docker-compose up --build
```

#### Get Database ready 🚀
```
  docker-compose exec gqlgen-example sql-migrate up
```

#### Set Up Seed 🚀
```
  docker-compose exec gqlgen-example go run database/seed/main.go
```


## Development 🛠

#### Installation
```sh
# code quality
brew install pre-commit golangci-lint
brew upgrade golangci-lint
pre-commit autoupdate
```
## Todo
- [ ] Docker Compose
  - [ ] app server
  - [ ] postgres server
- [ ] migration
- [ ] sample app
...
