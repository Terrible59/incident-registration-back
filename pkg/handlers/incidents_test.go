package handlers

import (
	"context"
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/guregu/null"
	"github.com/stretchr/testify/require"
	db "incidents_back/db/sqlc"
	"incidents_back/pkg/services"
	"incidents_back/pkg/utils"
	"net/http"
	"testing"
	"time"
)

func createRandomIncident(t *testing.T) *db.Incident {
	params := db.CreateIncidentParams{
		RegistrationDate: time.Now(),
		Summary: null.String{
			NullString: sql.NullString{
				String: utils.RandomString(10),
				Valid:  true,
			},
		},
		IncidentType: null.String{
			NullString: sql.NullString{
				String: utils.RandomString(10),
				Valid:  true,
			},
		},
	}

	incident, err := repo.CreateIncident(context.Background(), params)

	require.NoError(t, err)
	require.NotEmpty(t, incident)

	return &incident
}

func TestGetIncidentsRoute(t *testing.T) {
	user := createRandomUser()
	auth, _ := services.Login(user.Email, "test123", repo, maker)
	tests := []HTTPTest{
		{
			route:  "/api/v1/incidents",
			method: http.MethodGet,
			requestHeaders: map[string]string{
				"Authorization": "Bearer " + auth.AccessToken,
			},
			expectedStatusCode: fiber.StatusOK,
		},
		{
			route:  "/api/v1/incidents",
			method: http.MethodGet,
			requestHeaders: map[string]string{
				"Authorization": "Bearer " + auth.AccessToken + "sdfs",
			},
			expectedStatusCode: fiber.StatusUnauthorized,
		},
	}

	HttpTest(&tests, t)
}

func TestGetIncidentRoute(t *testing.T) {
	user := createRandomUser()
	auth, _ := services.Login(user.Email, "test123", repo, maker)

	incident := createRandomIncident(t)

	tests := []HTTPTest{
		{
			route:  "/api/v1/incidents/" + incident.ID.String(),
			method: http.MethodGet,
			requestHeaders: map[string]string{
				"Authorization": "Bearer " + auth.AccessToken,
			},
			expectedStatusCode: fiber.StatusOK,
			expectedBodyKeys:   []string{"id", "summary"},
		},
		{
			route:  "/api/v1/incidents/" + incident.ID.String(),
			method: http.MethodGet,
			requestHeaders: map[string]string{
				"Authorization": "Bearer " + auth.AccessToken + "sdfs",
			},
			expectedStatusCode: fiber.StatusUnauthorized,
		},
	}

	HttpTest(&tests, t)
}
