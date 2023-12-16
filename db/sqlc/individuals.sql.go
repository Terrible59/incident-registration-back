// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: individuals.sql

package db

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/guregu/null"
)

const createIndividual = `-- name: CreateIndividual :one
INSERT INTO individuals (
    first_name, last_name, middle_name, address, criminal_records_count
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING id, first_name, last_name, middle_name, address, criminal_records_count
`

type CreateIndividualParams struct {
	FirstName            null.String `json:"first_name"`
	LastName             null.String `json:"last_name"`
	MiddleName           null.String `json:"middle_name"`
	Address              null.String `json:"address"`
	CriminalRecordsCount null.Int    `json:"criminal_records_count"`
}

func (q *Queries) CreateIndividual(ctx context.Context, arg CreateIndividualParams) (Individual, error) {
	row := q.db.QueryRowContext(ctx, createIndividual,
		arg.FirstName,
		arg.LastName,
		arg.MiddleName,
		arg.Address,
		arg.CriminalRecordsCount,
	)
	var i Individual
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.MiddleName,
		&i.Address,
		&i.CriminalRecordsCount,
	)
	return i, err
}

const deleteIndividual = `-- name: DeleteIndividual :exec
DELETE FROM individuals
WHERE id=$1
`

func (q *Queries) DeleteIndividual(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteIndividual, id)
	return err
}

const getAllIndividuals = `-- name: GetAllIndividuals :many
SELECT id, first_name, last_name, middle_name, address, criminal_records_count FROM individuals
`

func (q *Queries) GetAllIndividuals(ctx context.Context) ([]Individual, error) {
	rows, err := q.db.QueryContext(ctx, getAllIndividuals)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Individual
	for rows.Next() {
		var i Individual
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.MiddleName,
			&i.Address,
			&i.CriminalRecordsCount,
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

const getIndividualById = `-- name: GetIndividualById :one
SELECT id, first_name, last_name, middle_name, address, criminal_records_count FROM individuals
WHERE id=$1
`

func (q *Queries) GetIndividualById(ctx context.Context, id uuid.UUID) (Individual, error) {
	row := q.db.QueryRowContext(ctx, getIndividualById, id)
	var i Individual
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.MiddleName,
		&i.Address,
		&i.CriminalRecordsCount,
	)
	return i, err
}

const updateIndividual = `-- name: UpdateIndividual :one
UPDATE individuals
SET first_name=$2, last_name=$3, middle_name=$4, address=$5, criminal_records_count=$6
WHERE id=$1
RETURNING id, first_name, last_name, middle_name, address, criminal_records_count
`

type UpdateIndividualParams struct {
	ID                   uuid.UUID   `json:"id"`
	FirstName            null.String `json:"first_name"`
	LastName             null.String `json:"last_name"`
	MiddleName           null.String `json:"middle_name"`
	Address              null.String `json:"address"`
	CriminalRecordsCount null.Int    `json:"criminal_records_count"`
}

func (q *Queries) UpdateIndividual(ctx context.Context, arg UpdateIndividualParams) (Individual, error) {
	row := q.db.QueryRowContext(ctx, updateIndividual,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.MiddleName,
		arg.Address,
		arg.CriminalRecordsCount,
	)
	var i Individual
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.MiddleName,
		&i.Address,
		&i.CriminalRecordsCount,
	)
	return i, err
}
