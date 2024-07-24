 
-- name: ProjectCreate :one
INSERT INTO projects(
project_name, 
category_id, 
project_image)
    VALUES ($1, $2,$3)
RETURNING
    *;
 
-- name: ProjectUpdate :one
UPDATE
    projects
SET
    project_name = $2,
    category_id = $3,
    project_image = $4 
WHERE
    project_id = $1
RETURNING
    *;

-- name: ProjectDeleteRestore :exec
UPDATE
    projects
SET
    deleted_at = IIF(deleted_at IS NULL , now() , null)
WHERE
    project_id = ANY (sqlc.arg('project_ids')::int[]);
 
 -- name: ProjectsList :many
 SELECT
        p.project_id,
        p.project_name,
        c.category_name,
        p.category_id,
        p.project_image,
        p.created_at,
        p.updated_at,
        p.deleted_at
    FROM
        projects p
        JOIN categories c ON p.category_id = c.category_id;
 



 -- name: ProjectFindForUpdate :one
 SELECT
        p.project_id,
        p."project_name", 
        p.category_id,
        p.project_image
    FROM
        projects p where p.project_id = $1 ;