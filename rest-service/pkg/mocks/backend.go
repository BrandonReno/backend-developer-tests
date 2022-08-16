package mocks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/go-chi/chi"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/server"
)

type BackendMock struct {
	Router     chi.Router
	TestServer *httptest.Server
}

func NewBackendMock() *BackendMock {
	return &BackendMock{
		Router: chi.NewRouter(),
	}
}

func (fx *BackendMock) MakeRequest(method, url string, body interface{}) (*http.Response, error) {
	marshaled, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, url, bytes.NewReader(marshaled))
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(req)
}

func (fx *BackendMock) UnmarshallResponseData(resp *http.Response, target interface{}) error {
	jsonResp := new(server.JSONResponse)
	if err := json.NewDecoder(resp.Body).Decode(&jsonResp); err != nil {
		return err
	}
	b, err := json.Marshal(jsonResp.Data)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, target)
}

func (fx *BackendMock) MakeURL(path string, parameters ...interface{}) string {
	return MakeURLFromServer(fx.TestServer, path, parameters...)
}

func MakeURLFromServer(server *httptest.Server, path string, parameters ...interface{}) string {
	return fmt.Sprintf("%s%s", server.URL, fmt.Sprintf(path, parameters...))
}
