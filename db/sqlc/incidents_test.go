package db

import (
	"context"
	"database/sql"
	"github.com/guregu/null"
	"github.com/stretchr/testify/require"
	"incidents_back/pkg/utils"
	"testing"
	"time"
)

func createRandomIncident(t *testing.T) *Incident {
	params := CreateIncidentParams{
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

	incident, err := testQueries.CreateIncident(context.Background(), params)

	require.NoError(t, err)
	require.NotEmpty(t, incident)

	return &incident
}

func TestQueries_CreateIncident(t *testing.T) {
	createRandomIncident(t)
}

func TestQueries_GetAllIncidents(t *testing.T) {
	createRandomIncident(t)

	incidents, err := testQueries.GetAllIncidents(context.Background())

	require.NoError(t, err)
	require.NotEmpty(t, incidents)
}

func TestQueries_GetIncidentById(t *testing.T) {
	incident := createRandomIncident(t)

	dbIncident, err := testQueries.GetIncidentById(context.Background(), incident.ID)

	require.NoError(t, err)
	require.NotEmpty(t, dbIncident)
	require.Equal(t, dbIncident.Summary.String, incident.Summary.String)
}

func TestQueries_UpdateIncident(t *testing.T) {
	incident := createRandomIncident(t)

	summary := utils.RandomString(10)

	params := UpdateIncidentParams{
		ID: incident.ID,
		Summary: null.String{
			NullString: sql.NullString{
				String: summary,
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

	dbIncident, err := testQueries.UpdateIncident(context.Background(), params)

	require.NoError(t, err)
	require.NotEmpty(t, dbIncident)
	require.Equal(t, dbIncident.Summary.String, summary)
}

func TestQueries_DeleteIncident(t *testing.T) {
	incident := createRandomIncident(t)

	_, err := testQueries.GetIncidentById(context.Background(), incident.ID)
	require.NoError(t, err)

	err = testQueries.DeleteIncident(context.Background(), incident.ID)
	require.NoError(t, err)

	_, err = testQueries.GetIncidentById(context.Background(), incident.ID)
	require.Error(t, err)
}
