-- name: GetAllIncidents :many
SELECT * FROM incidents;

-- name: GetIncidentById :one
SELECT * FROM incidents
WHERE id=$1;

-- name: CreateIncident :one
INSERT INTO incidents (
    registration_date, summary, incident_type
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: UpdateIncident :one
UPDATE incidents
SET summary=$2, incident_type=$3
WHERE id=$1
RETURNING *;

-- name: DeleteIncident :exec
DELETE FROM incidents
WHERE id=$1;