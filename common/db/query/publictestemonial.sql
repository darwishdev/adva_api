 
-- name: TestemonialCreate :one
INSERT INTO testemonials(
testemonial_name, 
testemonial_title, 
breif, 
testemonial_image)
    VALUES ($1, $2,$3,$4)
RETURNING
    *;
 
-- name: TestemonialUpdate :one
UPDATE
    testemonials
SET
    testemonial_name = $2,
    breif = $3,
    testemonial_image = $4,
    testemonial_title = $5 
WHERE
    testemonial_id = $1
RETURNING
    *;

-- name: TestemonialDeleteRestore :exec
UPDATE
    testemonials
SET
    deleted_at = IIF(deleted_at IS NULL , now() , null)
WHERE
    testemonial_id = ANY (sqlc.arg('testemonial_ids')::int[]);
 
 -- name: TestemonialsList :many
 SELECT
        t.testemonial_id,
        t.testemonial_name, 
        t.testemonial_title, 
        t.breif,
        t.testemonial_image,
        t.created_at,
        t.updated_at,
        t.deleted_at
    FROM
        testemonials t;
 -- name: TestemonialFindForUpdate :one
 SELECT
        t.testemonial_id,
        t.testemonial_name, 
        t.testemonial_title, 
        t.breif,
        t.testemonial_image
    FROM
        testemonials t where t.testemonial_id = $1;


