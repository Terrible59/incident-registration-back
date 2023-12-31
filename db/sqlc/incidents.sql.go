// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: incidents.sql

package db

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
)

const createIncident = `-- name: CreateIncident :one
INSERT INTO incidents (
    registration_date, summary, incident_type
) VALUES (
    $1, $2, $3
) RETURNING id, registration_date, summary, incident_type
`

type CreateIncidentParams struct {
	RegistrationDate time.Time   `json:"registration_date"`
	Summary          null.String `json:"summary"`
	IncidentType     null.String `json:"incident_type"`
}

func (q *Queries) CreateIncident(ctx context.Context, arg CreateIncidentParams) (Incident, error) {
	row := q.db.QueryRowContext(ctx, createIncident, arg.RegistrationDate, arg.Summary, arg.IncidentType)
	var i Incident
	err := row.Scan(
		&i.ID,
		&i.RegistrationDate,
		&i.Summary,
		&i.IncidentType,
	)
	return i, err
}

const deleteIncident = `-- name: DeleteIncident :exec
DELETE FROM incidents
WHERE id=$1
`

func (q *Queries) DeleteIncident(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteIncident, id)
	return err
}

const getAllIncidents = `-- name: GetAllIncidents :many
SELECT id, registration_date, summary, incident_type FROM incidents
`

func (q *Queries) GetAllIncidents(ctx context.Context) ([]Incident, error) {
	rows, err := q.db.QueryContext(ctx, getAllIncidents)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Incident
	for rows.Next() {
		var i Incident
		if err := rows.Scan(
			&i.ID,
			&i.RegistrationDate,
			&i.Summary,
			&i.IncidentType,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getIncidentById = `-- name: GetIncidentById :one
SELECT id, registration_date, summary, incident_type FROM incidents
WHERE id=$1
`

func (q *Queries) GetIncidentById(ctx context.Context, id uuid.UUID) (Incident, error) {
	row := q.db.QueryRowContext(ctx, getIncidentById, id)
	var i Incident
	err := row.Scan(
		&i.ID,
		&i.RegistrationDate,
		&i.Summary,
		&i.IncidentType,
	)
	return i, err
}

const updateIncident = `-- name: UpdateIncident :one
UPDATE incidents
SET summary=$2, incident_type=$3
WHERE id=$1
RETURNING id, registration_date, summary, incident_type
`

type UpdateIncidentParams struct {
	ID           uuid.UUID   `json:"id"`
	Summary      null.String `json:"summary"`
	IncidentType null.String `json:"incident_type"`
}

func (q *Queries) UpdateIncident(ctx context.Context, arg UpdateIncidentParams) (Incident, error) {
	row := q.db.QueryRowContext(ctx, updateIncident, arg.ID, arg.Summary, arg.IncidentType)
	var i Incident
	err := row.Scan(
		&i.ID,
		&i.RegistrationDate,
		&i.Summary,
		&i.IncidentType,
	)
	return i, err
}
