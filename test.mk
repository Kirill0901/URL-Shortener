run:
	go run ./cmd/url-shortener/main.go -host='172.31.204.125' -port=5432 -user=postgres -password=qwerty -dbname=gopgtest

test_postgres:
	go test -coverprofile=coverage.out  ./internal/storage/postgresql/ -host='172.31.204.125' -port=5432 -user=postgres -password=qwerty -dbname=gopgtest
	go tool cover -html=coverage.out

test_cache:
	go test -coverprofile=coverage.out  ./internal/storage/cache/
	go tool cover -html=coverage.out