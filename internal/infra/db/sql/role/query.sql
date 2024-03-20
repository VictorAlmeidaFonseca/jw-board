-- Roles
-- name: GetRoleByID :one
SELECT * FROM roles WHERE id = ?;

-- name: InsertRole :exec
INSERT INTO roles (id, description) VALUES (?, ?);

-- name: UpdateRoleDescription :exec
UPDATE roles SET description = ? WHERE id = ?;

-- name: DeleteRole :exec
DELETE FROM roles WHERE id = ?;

-- name: GetAllRoles :many
SELECT * FROM roles;