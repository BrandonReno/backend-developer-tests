package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/handler"
)

type (
	Params struct {
		Router        chi.Router
		PeopleHandler *handler.PeopleHandler
	}
)

func NewParams(router chi.Router, peopleHandler *handler.PeopleHandler) Params {
	return Params{Router: router, PeopleHandler: peopleHandler}
}

func MountRoutes(params Params) error {
	params.PeopleHandler.MountRoutes(params.Router)
	return nil
}

func NewRouter() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	return router
}

func RunService(router chi.Router) error {
	srv := http.Server{}
	srv.Addr = fmt.Sprintf(":%d", 8080)
	srv.Handler = router
	fmt.Println("serving")
	if err := srv.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
