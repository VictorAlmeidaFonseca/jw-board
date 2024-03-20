-- Person PersonRole
-- name: GetPersonRoleByID :one
SELECT * FROM person_role WHERE id = ?;

-- name: InsertPersonRole :exec
INSERT INTO person_role (id, person_id, role_id) VALUES (?, ?, ?);

-- name: UpdatePersonRole :exec
UPDATE person_role SET person_id = ?, role_id = ? WHERE id = ?;

-- name: DeletePersonRole :exec
DELETE FROM person_role WHERE id = ?;

-- name: GetAllPersonRoles :many
SELECT * FROM person_role;