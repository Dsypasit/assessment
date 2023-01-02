it_test_start:
	docker compose -f docker-compose.test.yml up --build --abort-on-container-exit --exit-code-from it_tests
it_test_down:
	docker-compose -f docker-compose.test.yml down
start:
	DATABASE_URL=postgres://root:1234@localhost/kbgt?sslmode=disable PORT=:2565 go run server.go
coverage:
	go test -v ./... -tags=unit -coverprofile=coverage.out && go tool cover -html=coverage.out
