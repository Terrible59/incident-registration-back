-- name: GetAllIndividuals :many
SELECT * FROM individuals;

-- name: GetIndividualById :one
SELECT * FROM individuals
WHERE id=$1;

-- name: CreateIndividual :one
INSERT INTO individuals (
    first_name, last_name, middle_name, address, criminal_records_count
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: UpdateIndividual :one
UPDATE individuals
SET first_name=$2, last_name=$3, middle_name=$4, address=$5, criminal_records_count=$6
WHERE id=$1
RETURNING *;

-- name: DeleteIndividual :exec
DELETE FROM individuals
WHERE id=$1;