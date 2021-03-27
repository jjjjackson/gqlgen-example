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

## Todo
- [ ] Docker Compose
  - [ ] app server
  - [ ] postgres server
- [ ] migration
- [ ] sample app  
...

