run_postgres:
	go run ./cmd/url-shortener/main.go -host='172.31.204.125' -port=5432 -user=postgres -password=qwerty -dbname=gopgtest

run_cache:
	go run ./cmd/url-shortener/main.go -storage=cache

generate_mocks:
	go run github.com/vektra/mockery/v2 --name=URLSaver --dir=./internal/http-server/handlers/url/save --output=./internal/http-server/handlers/url/save/mocks
	go run github.com/vektra/mockery/v2 --name=URLGetter --dir=./internal/http-server/handlers/url/get --output=./internal/http-server/handlers/url/get/mocks
	go run github.com/vektra/mockery/v2 --name=CountGetter --dir=./internal/shortener/ --output=./internal/shortener/mocks

test_postgres:
	go test -coverprofile=coverage.out  ./internal/storage/postgresql/ -host='172.31.204.125' -port=5432 -user=postgres -password=qwerty -dbname=gopgtest
	go tool cover -html=coverage.out

test_cache:
	go test -coverprofile=coverage.out  ./internal/storage/cache/
	go tool cover -html=coverage.out

test_save:
	go test -coverprofile=coverage.out  ./internal/http-server/handlers/url/save
	go tool cover -html=coverage.out

test_get:
	go test -coverprofile=coverage.out  ./internal/http-server/handlers/url/get
	go tool cover -html=coverage.out

test_shortener:
	go test -coverprofile=coverage.out  ./internal/shortener/
	go tool cover -html=coverage.out

test_all: test_postgres test_cache test_get test_save test_shortener