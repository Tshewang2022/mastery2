-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email, hashed_password)
VALUES (gen_random_uuid(), NOW(), NOW(),$1, $2) RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users;

-- name: CreateChirpy :one
INSERT INTO chirps (id, created_at, updated_at, body, user_id)
VALUES(gen_random_uuid(), NOW(), NOW(), $1, $2) RETURNING *;

-- name: GetChirps :many
SELECT * FROM chirps ORDER BY created_at ASC;

-- name: GetChirpsByID :one
SELECT * FROM chirps WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: UpdateUser :one
UPDATE users
SET email = $2,
    hashed_password = $3,
    updated_at = NOW() AT TIME ZONE 'UTC'
WHERE id = $1
RETURNING *;

-- name: DeleteChirp :exec
DELETE FROM chirps
WHERE id = $1 AND user_id = $2;

-- name: UpgradeUserToChirpyRed :one
UPDATE users
SET is_chirpy_red = true,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: GetChirpsByAuthor :many
SELECT * FROM chirps
WHERE user_id = $1
ORDER BY created_at ASC;