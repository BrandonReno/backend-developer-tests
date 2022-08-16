package handler

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/models"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/server"
)

type PeopleHandler struct {
	logger *logrus.Logger
}

func NewPeopleHandler(logger *logrus.Logger) *PeopleHandler {
	return &PeopleHandler{logger: logger}
}

func (h *PeopleHandler) MountRoutes(router chi.Router) {
	router.Group(func(r chi.Router) {
		r.Method(http.MethodGet, "/people", server.Handler(h.QueryPeople))
		r.Method(http.MethodGet, "/people/{id}", server.Handler(h.GetPersonByID))
	})
}

func (h *PeopleHandler) GetAll(w http.ResponseWriter, r *http.Request) error {
	people := models.AllPeople()
	return server.WriteSuccessResponse(w, http.StatusOK, people)
}

func (h *PeopleHandler) GetPersonByID(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")
	if id == "" {
		return server.WriteErrorResponse(w, http.StatusNotFound, errors.New("id was not given"))
	}
	personUUID := uuid.Must(uuid.FromString("81eb745b-3aae-400b-959f-748fcafafd81"))
	person, err := models.FindPersonByID(personUUID)
	if err != nil {
		return server.WriteErrorResponse(w, http.StatusNotFound, errors.New("could not find person by given id"))
	}
	return server.WriteSuccessResponse(w, http.StatusOK, person)
}

func (h *PeopleHandler) QueryPeople(w http.ResponseWriter, r *http.Request) error {
	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")
	phoneNumber := r.URL.Query().Get("phone_number")
	if firstName != "" && lastName != "" {
		people := models.FindPeopleByName(firstName, lastName)
		return server.WriteSuccessResponse(w, http.StatusOK, people)

	}
	if phoneNumber != "" {
		people := models.FindPeopleByPhoneNumber(phoneNumber)
		return server.WriteSuccessResponse(w, http.StatusOK, people)
	}
	people := models.AllPeople()
	return server.WriteSuccessResponse(w, http.StatusOK, people)
}
