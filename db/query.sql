-- name: CreateGuild :execresult
INSERT INTO guilds (name, discord_id)
VALUES ($1, $2);
-- name: FindByIDGuild :one
SELECT *
FROM guilds
WHERE id = $1
LIMIT 1;
-- name: FindByDiscordIDGuild :one
SELECT *
FROM guilds
WHERE discord_id = $1
LIMIT 1;
-- name: DeleteGuild :exec
DELETE FROM guilds
WHERE id = $1;
-- name: CreateToken :execresult
INSERT INTO tokens (
    system_user_id,
    access_token,
    token_type,
    refresh_token,
    expiry
  )
VALUES ($1, $2, $3, $4, $5);
-- name: FindByIDToken :one
SELECT *
FROM tokens
WHERE id = $1
LIMIT 1;
-- name: FindByUserIDToken :one
SELECT *
FROM tokens
WHERE system_user_id = $1
LIMIT 1;
-- name: DeleteToken :exec
DELETE FROM tokens
WHERE id = $1;
-- name: CreateSystemUser :execresult
INSERT INTO system_users (discord_id)
VALUES ($1);
-- name: FindByIDSystemUser :one
SELECT *
FROM system_users
WHERE id = $1
LIMIT 1;
-- name: FindByDiscordIDSystemUser :one
SELECT *
FROM system_users
WHERE discord_id = $1
LIMIT 1;
-- name: DeleteSystemUser :exec
DELETE FROM system_users
WHERE id = $1;
-- name: CreateSystemUserGuild :execresult
INSERT INTO system_user_guild (system_user_id, guild_id)
VALUES ($1, $2);
-- name: FindByIDSystemUserGuild :one
SELECT *
FROM system_user_guild
WHERE guild_id = $1
  AND system_user_id = $2;
-- name: FindByGuildIDSystemUserGuild :many
SELECT *
FROM system_user_guild
WHERE guild_id = $1;
-- name: FindBySystemUserIDSystemUserGuild :many
SELECT *
FROM system_user_guild
WHERE system_user_id = $1;
-- name: DeleteSystemUserGuild :exec
DELETE FROM system_user_guild
WHERE system_user_id = $1
  AND guild_id = $2;
