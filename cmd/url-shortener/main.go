package main

import (
	"log"
	"url-shortener/internal/http-server/handlers/url/get"
	"url-shortener/internal/http-server/handlers/url/save"
	"url-shortener/internal/shortener"
	"url-shortener/internal/storage"

	"github.com/labstack/echo/v4"
)

func main() {

	storage, err := storage.GetStorage()
	if err != nil {
		log.Fatal(err)
		return
	}

	defer storage.Close()

	err = shortener.CountInit(storage)
	if err != nil {
		log.Fatal(err)
		return
	}

	e := echo.New()

	e.POST("/", save.New(storage))
	e.GET("/*", get.New(storage))

	err = e.Start(":12000")
	if err != nil {
		log.Fatal(err)
	}
}
