-- Person
-- name: GetPersonByID :one
SELECT * FROM person WHERE id = ?;

-- name: GetAllPersons :many
SELECT * FROM person;

-- name: InsertPerson :exec
INSERT INTO person (id, name, role_id) VALUES (?, ?, ?);

-- name: UpdatePerson :exec
UPDATE person SET name = ?, role_id = ? WHERE id = ?;

-- name: DeletePerson :exec
DELETE FROM person WHERE id = ?;