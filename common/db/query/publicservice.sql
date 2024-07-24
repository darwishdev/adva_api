 
-- name: ServiceCreate :one
INSERT INTO services(
"service_name", 
breif, 
tags,
service_image)
    VALUES ($1, $2,$3 ,$4)
RETURNING
    *;
 
-- name: ServiceUpdate :one
UPDATE
    services
SET
    "service_name" = $2,
    breif = $3,
    tags = sqlc.arg('tags')::varchar[],
    service_image = $4 
WHERE
    service_id = $1
RETURNING
    *;

-- name: ServiceDeleteRestore :exec
UPDATE
    services
SET
    deleted_at = IIF(deleted_at IS NULL , now() , null)
WHERE
    service_id = ANY (sqlc.arg('service_ids')::int[]);
 
 -- name: ServicesList :many
 SELECT
        s.service_id,
        s."service_name", 
        s.breif,
        s.service_image,
        s.tags,
        s.created_at,
        s.updated_at,
        s.deleted_at
    FROM
        services s ;



 -- name: ServiceFindForUpdate :one
 SELECT
        s.service_id,
        s."service_name", 
        s.breif,
        s.tags,
        s.service_image
    FROM
        services s where service_id = $1 ;
 