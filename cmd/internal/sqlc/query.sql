-- name: GetPetById :one
SELECT * FROM pets
WHERE id = ? LIMIT 1;

-- name: CreatePetAndReturnId :execresult
INSERT INTO pets (id, name)
VALUES (?, ?);