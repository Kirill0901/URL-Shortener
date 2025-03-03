package get

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type URLGetter interface {
	GetURL(short_url string) (string, error)
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	LongURL string `json:"long_url"`
}

func New(urlGetter URLGetter) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Method != http.MethodGet {
			return c.JSON(http.StatusBadRequest, Response{
				Status:  "Error",
				Message: "Gets only GET request",
			})
		}

		short_url := strings.Replace(c.Request().URL.Path, "/", "", 1)

		if short_url == "" {
			return c.JSON(http.StatusBadRequest, Response{
				Status:  "Error",
				Message: "Bad request",
			})
		}

		long_url, err := urlGetter.GetURL(short_url)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, Response{
				Status:  "Error",
				Message: "Internal server error",
			})
		}

		return c.JSON(http.StatusOK, Response{
			Status:  "Success",
			Message: "Long URL is found",
			LongURL: long_url,
		})
	}
}
