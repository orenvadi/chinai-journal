
# INAI attendance journal


## How to run

To run the project you must have a [SurrealDB](https://surrealdb.com/docs/surrealdb/installation/) database, configs are must be located in .env

**Create the database**

you must be in a root of project

```sh
surreal start --user root --pass root --bind 0.0.0.0:8080 file:storage
```


**Migrate the database**

```sh
make migrate
```

if it does not work, you have to adjust the main Makefile


**Run the server on local machine**

```sh
make run_local
```


## How to change and regenerate protos

### Requirements

- [buf cli](https://buf.build/docs/installation) 


At first you should change directory

```sh
cd protos
```

Then

```sh
buf mod update
```

```sh
buf generate
```


## Important note!!!

I could not solve issue with generated `sso.pb.go`
so every time you regenerate protos

**Replace**

```go
	_ "buf/validate"
```

**With this**
```go
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
```





## TODO 

- [x] Finish MVP

- [ ] Connect to ebilim through web scraping

- [ ] Reset password

- [x] Get user data

- [ ] Refresh token

- [ ] Write more functional migrator

- [ ] Refactor code, cause there are a lot of boilerplate code repetition, especially in domain/models and internal/storage
