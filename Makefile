it_test:
	go clean --cache && DATABASE_URL=postgres://root:1234@localhost/kbgt?sslmode=disable PORT=:2565 go test --tags=integration -v ./...
start:
	DATABASE_URL=postgres://root:1234@localhost/kbgt?sslmode=disable PORT=:2565 go run server.go
