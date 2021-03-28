# gqlgen-example

## Start 🌟
Make sure you got docker installed
#### Run Docker 🐳
```
  docker-compose up --build
```

#### Get Database ready 🚐
```
  docker-compose exec gqlgen-example sql-migrate up
```

#### Set Up Seed 🌱
```
  docker-compose exec gqlgen-example go run database/seed/main.go
```

#### Try 🚀
Access `localhost:3333` to play ⚾

## Development 🛠

#### Before Development 🥚

###### Installation
```sh
# code quality
brew install pre-commit golangci-lint
brew upgrade golangci-lint
pre-commit autoupdate
```

#### File Structure 📁

#### How to Develop 🐣

1. create `schema/XXXX.graphqls`
2. run `gqlgen` from console
3. (optional) create `model` and link the model by `gqlgen.yml`
4. build the repository and service to link the resolver
5. try it on `localhost:3333` 🐔

## Todo 🗒
- [x] Docker Compose
  - [x] app server
  - [x] postgres server
- [x] pre-commit to ensure the quality
- [x] migration
- [x] seed
- [ ] app
  - [ ] login + jwt
  - [ ] mutation sample
  - [ ] query sample
- [ ] test
- [ ] CI
