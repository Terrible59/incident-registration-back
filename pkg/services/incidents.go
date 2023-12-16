package services

import (
	"context"
	"database/sql"
	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	db "incidents_back/db/sqlc"
	"time"
)

func GetIncidents(repo *db.Repo) ([]db.Incident, error) {
	return repo.GetAllIncidents(context.Background())
}

func GetIncident(id uuid.UUID, repo *db.Repo) (db.Incident, error) {
	return repo.GetIncidentById(context.Background(), id)
}

func CreateIncident(summary, incidentType string, repo *db.Repo) (db.Incident, error) {
	return repo.CreateIncident(context.Background(), db.CreateIncidentParams{
		Summary: null.String{
			NullString: sql.NullString{
				String: summary,
				Valid:  true,
			},
		},
		IncidentType: null.String{
			NullString: sql.NullString{
				String: incidentType,
				Valid:  true,
			},
		},
		RegistrationDate: time.Now(),
	})
}

func UpdateIncident(id uuid.UUID, summary, incidentType string, repo *db.Repo) (db.Incident, error) {
	return repo.UpdateIncident(context.Background(), db.UpdateIncidentParams{
		ID: id,
		Summary: null.String{
			NullString: sql.NullString{
				String: summary,
				Valid:  true,
			},
		},
		IncidentType: null.String{
			NullString: sql.NullString{
				String: incidentType,
				Valid:  true,
			},
		},
	})
}

func DeleteIncident(id uuid.UUID, repo *db.Repo) error {
	return repo.DeleteIncident(context.Background(), id)
}
