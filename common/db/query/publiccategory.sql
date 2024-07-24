 
-- name: CategoryCreate :one
INSERT INTO categories(
category_name, 
category_type_id )
    VALUES ($1, $2)
RETURNING
    *;
 
-- name: CategoryUpdate :one
UPDATE
    categories
SET
    category_name = $2,

    category_type_id = $3 
WHERE
    category_id = $1
RETURNING
    *;

-- name: CategoryDeleteRestore :exec
UPDATE
    categories
SET
    deleted_at = IIF(deleted_at IS NULL , now() , null)
WHERE
    category_id = ANY (sqlc.arg('category_ids')::int[]);

-- name: CategoryFind :one
SELECT
    c.category_id,
    c.category_name, 
    c.category_type_id,
    ct.category_type,
    c.created_at,
    c.updated_at,
    c.deleted_at
FROM
    categories c
    JOIN category_types ct ON c.category_type_id = ct.category_type_id
WHERE
    c.category_id = $1 ;
 
 -- name: CategoriesList :many
 SELECT
    c.category_id,
    c.category_name, 
    c.category_type_id,
    ct.category_type,
    c.created_at,
    c.updated_at,
    c.deleted_at
    FROM
        categories c
    JOIN category_types ct ON c.category_type_id = ct.category_type_id
    where c.category_type_id = IsZero($1, c.category_type_id);

-- name: CategoriesInputList :many
SELECT
    category_id,
    category_name
FROM
    categories
WHERE
    deleted_at IS NULL and category_type_id = IsZero($1, category_type_id);



-- name: CategoryFindForUpdate :one
SELECT
    category_id,
    category_name,
    category_type_id
FROM
    categories
WHERE
    category_id = $1;
