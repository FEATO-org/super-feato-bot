-- name: CreateGuild :execresult
INSERT INTO guilds (name, discord_id)
VALUES (?, ?);
-- name: FindByIDGuild :one
SELECT *
FROM guilds
WHERE id = ?
LIMIT 1;
-- name: FindByDiscordIDGuild :one
SELECT *
FROM guilds
WHERE discord_id = $1
LIMIT 1;
-- name: DeleteGuild :exec
DELETE FROM guilds
WHERE id = ?;
-- name: CreateToken :execresult
INSERT INTO tokens (
    system_user_id,
    access_token,
    token_type,
    refresh_token,
    expiry
  )
VALUES (?, ?, ?, ?, ?);
-- name: FindByIDToken :one
SELECT *
FROM tokens
WHERE id = ?
LIMIT 1;
-- name: FindByUserIDToken :one
SELECT *
FROM tokens
WHERE system_user_id = ?
LIMIT 1;
-- name: DeleteToken :exec
DELETE FROM tokens
WHERE id = ?;
-- name: CreateSystemUser :execresult
INSERT INTO system_users (discord_id)
VALUES (?);
-- name: FindByIDSystemUser :one
SELECT *
FROM system_users
WHERE id = ?
LIMIT 1;
-- name: FindByDiscordIDSystemUser :one
SELECT *
FROM system_users
WHERE discord_id = ?
LIMIT 1;
-- name: DeleteSystemUser :exec
DELETE FROM system_users
WHERE id = ?;
-- name: CreateSystemUserGuild :execresult
INSERT INTO system_user_guild (system_user_id, guild_id)
VALUES (?, ?);
-- name: FindByIDSystemUserGuild :one
SELECT *
FROM system_user_guild
WHERE guild_id = ?
  AND system_user_id = ?;
-- name: FindByGuildIDSystemUserGuild :many
SELECT *
FROM system_user_guild
WHERE guild_id = ?;
-- name: FindBySystemUserIDSystemUserGuild :many
SELECT *
FROM system_user_guild
WHERE system_user_id = ?;
-- name: DeleteSystemUserGuild :exec
DELETE FROM system_user_guild
WHERE system_user_id = ?
  AND guild_id = ?;
