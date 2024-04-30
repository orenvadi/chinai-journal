run_local:
	go run cmd/sso/main.go

migrate:
	go run ./cmd/migrator --host=ws://localhost:8080/rpc --dbname=test --dbnamespace=test --migr-path=./migrations/surreal/1_init.up.surql
