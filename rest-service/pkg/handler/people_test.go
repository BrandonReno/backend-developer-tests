package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stackpath/backend-developer-tests/rest-service/pkg/config"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/mocks"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/models"
	"github.com/stretchr/testify/require"
)

func TestPeople_GetAll(t *testing.T) {
	t.Run("GetAll -- Success", func(t *testing.T) {
		SetupEndpoints(t, func(fixture *mocks.BackendMock) {
			resp, err := fixture.MakeRequest(http.MethodGet, fixture.MakeURL("/people"), nil)
			require.NoError(t, err)
			require.Equal(t, resp.StatusCode, http.StatusOK)
			var people []*models.Person
			err = fixture.UnmarshallResponseData(resp, &people)
			require.NoError(t, err)
			require.NotNil(t, people)
		})
	})
}

func TestPeople_GetByID(t *testing.T) {
	t.Run("GetAll -- Success", func(t *testing.T) {
		SetupEndpoints(t, func(fixture *mocks.BackendMock) {
			resp, err := fixture.MakeRequest(http.MethodGet, fixture.MakeURL("/people/3"), nil)
			require.NoError(t, err)
			require.Equal(t, resp.StatusCode, http.StatusOK)
			var person models.Person
			err = fixture.UnmarshallResponseData(resp, &person)
			require.NoError(t, err)
			require.NotNil(t, person)
		})
	})
}

func TestPeople_GetByFirstLastName(t *testing.T) {
	t.Run("GetAll -- Success", func(t *testing.T) {
		firstName := "John"
		lastName := "Doe"
		route := fmt.Sprintf("/people?first_name=%s&last_name=%s", firstName, lastName)
		SetupEndpoints(t, func(fixture *mocks.BackendMock) {
			resp, err := fixture.MakeRequest(http.MethodGet, fixture.MakeURL(route), nil)
			require.NoError(t, err)
			require.Equal(t, resp.StatusCode, http.StatusOK)
			var people []*models.Person
			err = fixture.UnmarshallResponseData(resp, &people)
			require.NoError(t, err)
			require.NotNil(t, people)
			for _, person := range people {
				require.Equal(t, person.FirstName, firstName)
				require.Equal(t, person.LastName, lastName)
			}
		})
	})
}

func TestPeople_GetByPhone(t *testing.T) {
	t.Run("GetAll -- Success", func(t *testing.T) {
		phone := "+447700900077"
		route := fmt.Sprintf("/people?phone_number=%s", phone)
		SetupEndpoints(t, func(fixture *mocks.BackendMock) {
			resp, err := fixture.MakeRequest(http.MethodGet, fixture.MakeURL(route), nil)
			require.NoError(t, err)
			require.Equal(t, resp.StatusCode, http.StatusOK)
			var people []*models.Person
			err = fixture.UnmarshallResponseData(resp, &people)
			require.NoError(t, err)
			require.NotNil(t, people)
			for _, person := range people {
				require.Equal(t, person.PhoneNumber, phone)
			}
		})
	})
}

func SetupEndpoints(t *testing.T, testBody func(fixture *mocks.BackendMock)) {
	fixture := mocks.NewBackendMock()
	logger := config.NewLogger()
	handler := NewPeopleHandler(logger)
	handler.MountRoutes(fixture.Router)
	fixture.TestServer = httptest.NewServer(fixture.Router)
	defer fixture.TestServer.Close()
	testBody(fixture)
}
