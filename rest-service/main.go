package main

import (
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/config"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/handler"
)

// ********* TO RUN USE 'go run *.go' *********
func main() {
	router := NewRouter()
	logger := config.NewLogger()
	peopleHandler := handler.NewPeopleHandler(logger)
	Params := NewParams(router, peopleHandler)
	if err := MountRoutes(Params); err != nil {
		panic("error mounting routes")
	}
	if err := RunService(router); err != nil {
		panic("error starting service")
	}
}
