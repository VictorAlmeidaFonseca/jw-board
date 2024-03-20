-- Assignments
-- name: GetAssignmentByID :one
SELECT * FROM assignments WHERE id = ?;

-- name: GetAllAssignments :many
SELECT * FROM assignments;

-- name: InsertAssignment :exec
INSERT INTO assignments (id, role_id, person_id) VALUES (?, ?, ?);

-- name: UpdateAssignment :exec
UPDATE assignments SET role_id = ?, person_id = ? WHERE id = ?;

-- name: DeleteAssignment :exec
DELETE FROM assignments WHERE id = ?;