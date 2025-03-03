package shortener

import (
	"errors"
	"testing"

	"url-shortener/internal/shortener/mocks"

	"github.com/stretchr/testify/assert"
)

func TestCountInit_ValidCountGetter(t *testing.T) {
	mockCountGetter := mocks.NewCountGetter(t)
	mockCountGetter.On("GetCount").Return(int64(5), nil)

	err := CountInit(mockCountGetter)
	assert.NoError(t, err)
	assert.Equal(t, int64(5), count)
}

func TestCountInit_ErrorFromCountGetter(t *testing.T) {
	mockCountGetter := mocks.NewCountGetter(t)
	mockCountGetter.On("GetCount").Return(int64(0), errors.New("some error"))

	err := CountInit(mockCountGetter)
	assert.Error(t, err)
	assert.Equal(t, "some error", err.Error())
	assert.Equal(t, int64(0), count)
}

func TestMakeShorter(t *testing.T) {
	mockCountGetter := mocks.NewCountGetter(t)
	mockCountGetter.On("GetCount").Return(int64(0), nil)

	_ = CountInit(mockCountGetter)
	shortURL, err := MakeShorter("http://example.com")
	assert.NoError(t, err)
	assert.Equal(t, "aaaaaaaaaa", shortURL)
}

func TestMakeShorter_AfterCountInit(t *testing.T) {
	mockCountGetter := mocks.NewCountGetter(t)
	mockCountGetter.On("GetCount").Return(int64(1), nil)

	_ = CountInit(mockCountGetter)
	shortURL1, _ := MakeShorter("http://example.com")
	shortURL2, _ := MakeShorter("http://example.com")
	assert.Equal(t, "baaaaaaaaa", shortURL1)
	assert.Equal(t, "caaaaaaaaa", shortURL2)
}
