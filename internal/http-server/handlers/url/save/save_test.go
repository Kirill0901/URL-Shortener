package save_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"

	"url-shortener/internal/http-server/handlers/url/save"
	"url-shortener/internal/http-server/handlers/url/save/mocks"
)

func TestSaveHandler(t *testing.T) {
	cases := []struct {
		name               string
		httpMethod         string
		longURL            string
		shortURL           string
		existing_short_url string
		statusCode         int
		respError          string
		mockError          error
	}{
		{
			name:       "Success",
			httpMethod: http.MethodPost,
			longURL:    "https://google.com",
			shortURL:   "aaaaaaaaaa",
			statusCode: http.StatusOK,
		},
		{
			name:               "Success",
			httpMethod:         http.MethodPost,
			longURL:            "https://google.com",
			shortURL:           "baaaaaaaaa",
			existing_short_url: "aaaaaaaaaa",
			statusCode:         http.StatusOK,
		},
		{
			name:       "Empty URL",
			httpMethod: http.MethodPost,
			longURL:    "",
			statusCode: http.StatusBadRequest,
			respError:  "Bad request",
		},
		{
			name:       "SaveURL Error",
			httpMethod: http.MethodPost,
			longURL:    "https://google.com",
			shortURL:   "caaaaaaaaa",
			statusCode: http.StatusInternalServerError,
			respError:  "Internal server error",
			mockError:  errors.New("unexpected error"),
		},
		{
			name:       "Not POST method",
			httpMethod: http.MethodGet,
			longURL:    "https://ya.ru",
			statusCode: http.StatusBadRequest,
			respError:  "Gets only POST request",
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			urlSaverMock := mocks.NewURLSaver(t)

			urlSaverMock.On("SaveURL", tc.longURL, tc.shortURL).
				Return(tc.existing_short_url, tc.mockError).Maybe()

			handler := save.New(urlSaverMock)

			input := fmt.Sprintf(`{"long_url": "%s"}`, tc.longURL)

			e := echo.New()
			req := httptest.NewRequest(tc.httpMethod, "/", strings.NewReader(input))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := handler(c)
			require.NoError(t, err)

			require.Equal(t, tc.statusCode, rec.Code)
			body := rec.Body.String()
			var resp save.Response
			require.NoError(t, json.Unmarshal([]byte(body), &resp))

			if tc.respError != "" {
				require.Equal(t, tc.respError, resp.Message)
			} else {
				require.Equal(t, "Success", resp.Status)
			}
		})
	}
}
