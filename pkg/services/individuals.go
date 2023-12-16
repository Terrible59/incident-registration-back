package services

import (
	"context"
	"database/sql"
	"github.com/gofrs/uuid"
	"github.com/guregu/null"
	db "incidents_back/db/sqlc"
)

func GetIndividuals(repo *db.Repo) ([]db.Individual, error) {
	return repo.GetAllIndividuals(context.Background())
}

func GetIndividual(id uuid.UUID, repo *db.Repo) (db.Individual, error) {
	return repo.GetIndividualById(context.Background(), id)
}

func CreateIndividual(firstName, lastName, middleName, address string, recordsCount int64, repo *db.Repo) (db.Individual, error) {
	return repo.CreateIndividual(context.Background(), db.CreateIndividualParams{
		FirstName: null.String{
			NullString: sql.NullString{
				String: firstName,
				Valid:  true,
			},
		},
		LastName: null.String{
			NullString: sql.NullString{
				String: lastName,
				Valid:  true,
			},
		},
		MiddleName: null.String{
			NullString: sql.NullString{
				String: middleName,
				Valid:  true,
			},
		},
		Address: null.String{
			NullString: sql.NullString{
				String: address,
				Valid:  true,
			},
		},
		CriminalRecordsCount: null.Int{
			NullInt64: sql.NullInt64{
				Int64: recordsCount,
				Valid: true,
			},
		},
	})
}

func UpdateIndividual(id uuid.UUID, firstName, lastName, middleName, address string, recordsCount int64, repo *db.Repo) (db.Individual, error) {
	return repo.UpdateIndividual(context.Background(), db.UpdateIndividualParams{
		ID: id,
		FirstName: null.String{
			NullString: sql.NullString{
				String: firstName,
				Valid:  true,
			},
		},
		LastName: null.String{
			NullString: sql.NullString{
				String: lastName,
				Valid:  true,
			},
		},
		MiddleName: null.String{
			NullString: sql.NullString{
				String: middleName,
				Valid:  true,
			},
		},
		Address: null.String{
			NullString: sql.NullString{
				String: address,
				Valid:  true,
			},
		},
		CriminalRecordsCount: null.Int{
			NullInt64: sql.NullInt64{
				Int64: recordsCount,
				Valid: true,
			},
		},
	})
}

func DeleteIndividual(id uuid.UUID, repo *db.Repo) error {
	return repo.DeleteIndividual(context.Background(), id)
}
