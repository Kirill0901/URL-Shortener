package save

import (
	"net/http"
	"url-shortener/internal/shortener"

	"github.com/labstack/echo/v4"
)

type URLSaver interface {
	SaveURL(long_url, short_url string) error
}

type Request struct {
	LongURL string `json:"long_url"`
}

type Response struct {
	Status   string `json:"status"`
	Message  string `json:"message"`
	ShortURL string `json:"short_url"`
}

func New(urlSaver URLSaver) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Method != http.MethodPost {
			return c.JSON(http.StatusBadRequest, Response{
				Status:   "Error",
				Message:  "Gets only POST request",
				ShortURL: "",
			})
		}

		var request Request
		if err := c.Bind(&request); err != nil || request.LongURL == "" {
			return c.JSON(http.StatusBadRequest, Response{
				Status:   "Error",
				Message:  "Bad request",
				ShortURL: "",
			})
		}

		short_url, err := shortener.MakeShorter(request.LongURL)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Response{
				Status:   "Error",
				Message:  "Internal server error",
				ShortURL: "",
			})
		}

		err = urlSaver.SaveURL(request.LongURL, short_url)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Response{
				Status:   "Error",
				Message:  "Internal server error",
				ShortURL: "",
			})
		}

		return c.JSON(http.StatusOK, Response{
			Status:   "Success",
			Message:  "ShortURL for " + request.LongURL + " created succesfully",
			ShortURL: short_url,
		})
	}
}
