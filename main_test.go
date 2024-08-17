package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainHandlerBodyOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	city := req.URL.Query().Get("city")
	countStr := req.URL.Query().Get("count")
	count, err := strconv.Atoi(countStr)
	if err != nil {
		fmt.Println("conversion error")
		return
	}

	body := responseRecorder.Body
	stBody := body.String()

	list := cafeList[city]
	list = list[:count]
	stList := strings.Join(list, ",")

	assert.NotEmpty(t, body)
	assert.Equal(t, stBody, stList)
}

func TestMainHandlerWhenOk(t *testing.T) {
	req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code

	assert.Equal(t, status, http.StatusOK)

}

func TestMainHandlerCityOk(t *testing.T) {
	errorMessage := "wrong city value"

	req := httptest.NewRequest("GET", "/cafe?count=4&city=che", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	status := responseRecorder.Code
	resBody := responseRecorder.Body

	assert.Equal(t, status, http.StatusBadRequest)
	assert.Equal(t, resBody.String(), errorMessage)

}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4

	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()
	list := strings.Split(body, ",")

	assert.Len(t, list, totalCount)

}

