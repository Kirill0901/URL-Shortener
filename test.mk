run:
	go run ./cmd/url-shortener/main.go -host='172.31.204.125' -port=5432 -user=postgres -password=qwerty -dbname=gopgtest

test:
	go test -coverprofile=coverage.out  ./internal/storage/postgresql/ -host='172.31.204.125' -port=5432 -user=postgres -password=qwerty -dbname=gopgtest
	go tool cover -html=coverage.out