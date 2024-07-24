 
-- name: ProgramCreate :one
INSERT INTO programs(
"program_name", 
breif, 
program_image)
    VALUES ($1, $2,$3)
RETURNING
    *;
 
-- name: ProgramUpdate :one
UPDATE
    programs
SET
    "program_name" = $2,
    breif = $3,
    program_image = $4 
WHERE
    program_id = $1
RETURNING
    *;

-- name: ProgramDeleteRestore :exec
UPDATE
    programs
SET
    deleted_at = IIF(deleted_at IS NULL , now() , null)
WHERE
    program_id = ANY (sqlc.arg('program_ids')::int[]);
 
 -- name: ProgramsList :many
 SELECT
        p.program_id,
        p."program_name", 
        p.breif,
        p.program_image,
        p.created_at,
        p.updated_at,
        p.deleted_at
    FROM
        programs p ;



 -- name: ProgramFindForUpdate :one
 SELECT
        p.program_id,
        p."program_name", 
        p.breif,
        p.program_image
    FROM
        programs p where program_id = $1 ;
 