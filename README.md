URL shortener on Golang.

# Docker run

## postgresql
```sh
docker run --publish 12000:12000 --name url-shortener --rm url-shortener -storage=postgresql -host='172.17.0.2' -port=5432 -user=postgres -password=qwerty -dbname=gopgtest
```

## cache
```sh
docker run --publish 12000:12000 --name url-shortener --rm url-shortener -storage=cache
```
