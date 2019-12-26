package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	expTime       = 60
	longURL       = "https://www.example.com"
	shortLink     = "IFHzao"
	shortLinkInfo = `{"url":"https://www.example.com","created_at":"2017-06-19 15:55:55"}`
)

type storageMock struct {
	mock.Mock
}

var app App
var mockR *storageMock

func (s *storageMock) Shorten(url string, exp int64) (string, error) {
	args := s.Called(url, exp)
	return args.String(0), args.Error(1)
}

func (s *storageMock) Unshorten(eid string) (string, error) {
	args := s.Called(eid)
	return args.String(0), args.Error(1)
}
func (s *storageMock) ShortlinkInfo(eid string) (interface{}, error) {
	args := s.Called(eid)
	return args.String(0), args.Error(1)
}

func init() {
	app = App{}
	mockR = new(storageMock)
	app.Initialize(&Env{S: mockR})
}

func TestCreateShortlink(t *testing.T) {
	var jsonStr = []byte(`{

		"url":"https://www.example.com",
		"expiration_in_minutes":60
`)

	req, err := http.NewRequest("POST", "/api/shorten",
		bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal("创建request失败", err)

	}

	req.Header.Set("Content-Type", "application/json")
	mockR.On("Shorten", longURL, int64(expTime)).Return(shortLink, nil).Once()
	rw := httptest.NewRecorder()
	app.Router.ServeHTTP(rw, req)

	if rw.Code != http.StatusCreated {
		t.Fatalf("expect receive %d. Got %d", http.StatusCreated, rw.Code)
	}
	resp := struct {
		Shortlink string `json:"shortlink"`
	}{}
	if err := json.NewDecoder(rw.Body).Decode(&resp); err != nil {
		t.Fatalf("should decode the response")
	}
	if resp.Shortlink != shortLink {
		t.Fatalf("expect receive %s. Got %s", shortLink, resp.Shortlink)

	}

}

func TestRedirect(t *testing.T) {
	r := fmt.Sprintf("/%s", shortLink)
	req, err := http.NewRequest("GET", r, nil)
	if err != nil {
		t.Fatal("Should be able to create a request.", err)
	}

	mockR.On("Unshorten", shortLink).Return(longURL, nil).Once()
	rw := httptest.NewRecorder()
	app.Router.ServeHTTP(rw, req)

	if rw.Code != http.StatusTemporaryRedirect {
		t.Fatalf("Excepted receice %d. Got %d", http.StatusTemporaryRedirect, rw.Code)
	}
}
