 
-- name: TeamMemberCreate :one
INSERT INTO team_members(
team_member_name, 
job_title, 
team_member_image)
    VALUES ($1, $2,$3)
RETURNING
    *;
 
-- name: TeamMemberUpdate :one
UPDATE
    team_members
SET
    team_member_name = $2,
    team_member_image = $3,
    job_title = $4 
WHERE
    team_member_id = $1
RETURNING
    *;

-- name: TeamMemberDeleteRestore :exec
UPDATE
    team_members
SET
    deleted_at = IIF(deleted_at IS NULL , now() , null)
WHERE
    team_member_id = ANY (sqlc.arg('team_member_ids')::int[]);
 
 -- name: TeamMembersList :many
 SELECT
        t.team_member_id,
        t.team_member_name, 
        t.job_title, 
        t.team_member_image,
        t.created_at,
        t.updated_at,
        t.deleted_at
    FROM
        team_members t;
 -- name: TeamMemberFindForUpdate :one
 SELECT
        t.team_member_id,
        t.team_member_name, 
        t.job_title, 
        t.team_member_image
    FROM
        team_members t where t.team_member_id = $1;