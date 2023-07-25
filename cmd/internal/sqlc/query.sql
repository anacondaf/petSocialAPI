-- name: GetPetById :one
SELECT * FROM pets
WHERE id = ? LIMIT 1;