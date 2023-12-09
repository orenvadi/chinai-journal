run:
	go run cmd/chinai-journal/main.go                                                              ─╯

build:
	go build cmd/chinai-journal/main.go
	mv main backend-app

build-win:
	GOOS=windows GOARCH=amd64 go build ./cmd/chinai-journal/main.go
	mv main.exe backend-app.exe

build-mac:
	GOOS=darwin GOARCH=amd64 go build ./cmd/chinai-journal/main.go
	mv main backend-app

build-arm:
	GOOS=linux GOARCH=arm go build ./cmd/chinai-journal/main.go
	mv main backend-app
