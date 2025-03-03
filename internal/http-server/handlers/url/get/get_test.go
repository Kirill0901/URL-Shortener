package get_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"

	"url-shortener/internal/http-server/handlers/url/get"
	"url-shortener/internal/http-server/handlers/url/get/mocks"
)

func TestGetHandler(t *testing.T) {
	cases := []struct {
		name       string
		httpMethod string
		shortURL   string
		longURL    string
		statusCode int
		respError  string
		mockError  error
	}{
		{
			name:       "Success",
			httpMethod: http.MethodGet,
			shortURL:   "shorturl",
			longURL:    "https://example.com",
			statusCode: http.StatusOK,
		},
		{
			name:       "Empty Short URL",
			httpMethod: http.MethodGet,
			shortURL:   "",
			statusCode: http.StatusBadRequest,
			respError:  "Bad request",
		},
		{
			name:       "GetURL Error",
			httpMethod: http.MethodGet,
			shortURL:   "shorturl",
			statusCode: http.StatusInternalServerError,
			respError:  "Internal server error",
			mockError:  errors.New("unexpected error"),
		},
		{
			name:       "Not GET Method",
			httpMethod: http.MethodPost,
			shortURL:   "shorturl",
			statusCode: http.StatusBadRequest,
			respError:  "Gets only GET request",
		},
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			urlGetterMock := mocks.NewURLGetter(t)

			urlGetterMock.On("GetURL", tc.shortURL).
				Return(tc.longURL, tc.mockError).Maybe()

			handler := get.New(urlGetterMock)

			req := httptest.NewRequest(tc.httpMethod, "/"+tc.shortURL, nil)
			rec := httptest.NewRecorder()
			c := echo.New().NewContext(req, rec)

			err := handler(c)
			require.NoError(t, err)

			require.Equal(t, tc.statusCode, rec.Code)
			body := rec.Body.String()
			var resp get.Response
			require.NoError(t, json.Unmarshal([]byte(body), &resp))

			if tc.respError != "" {
				require.Equal(t, tc.respError, resp.Message)
			} else {
				require.Equal(t, tc.longURL, resp.LongURL)
				require.Equal(t, tc.statusCode, rec.Code)
			}
		})
	}
}
