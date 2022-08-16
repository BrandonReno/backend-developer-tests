package server

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type JSONResponse struct {
	Status int         `json:"status,omitempty"`
	Data   interface{} `json:"data,omitempty"`
	Error  interface{} `json:"error,omitempty"`
}

func WriteSuccessResponse(w http.ResponseWriter, status int, data interface{}) error {
	return WriteJSON(w, status, JSONResponse{Status: status, Data: data})
}

func WriteErrorResponse(w http.ResponseWriter, status int, err interface{}) error {
	return WriteJSON(w, status, JSONResponse{Status: status, Error: err})
}

func WriteJSON(w http.ResponseWriter, status int, body interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		return err
	}
	return nil
}

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		var data interface{}
		status := http.StatusInternalServerError
		data = err
		if err := WriteErrorResponse(w, status, data); err != nil {
			logrus.Errorf("unable to write the json response: %v", err)
		}
	}
}
