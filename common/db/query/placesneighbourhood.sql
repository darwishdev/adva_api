-- name: NeighbourhoodsList :many
SELECT
    n.neighbourhood_id,
    neighbourhood_name,
    neighbourhood_code,
    c.city_name,
    c.city_code,
    c.city_id,
    d.district_name,
    d.district_code,
    d.district_id,
    n.created_at,
    n.deleted_at
FROM
    places_schema.neighbourhoods n
    JOIN places_schema.districts d ON d.district_id = n.district_id
    JOIN places_schema.cities c ON d.city_id = c.city_id;

-- name: NeighbourhoodsInputList :many
SELECT
    neighbourhood_id,
    neighbourhood_name
FROM
    places_schema.neighbourhoods
WHERE
    district_id = $1
    AND deleted_at IS NULL;

-- name: NeighbourhoodDeleteRestore :exec
UPDATE
    places_schema.neighbourhoods
SET
    deleted_at = CASE WHEN deleted_at IS NULL THEN
        now()
    ELSE
        NULL
    END
WHERE
    neighbourhood_id = ANY (sqlc.arg('neighbourhood_ids')::int[]);

-- name: NeighbourhoodCreate :one
INSERT INTO places_schema.neighbourhoods(district_id, neighbourhood_name, neighbourhood_code)
    VALUES ($1, $2, $3)
RETURNING
    *;

-- name: NeighbourhoodUpdate :one
UPDATE
    places_schema.neighbourhoods
SET
    neighbourhood_name = $2,
    neighbourhood_code = $3,
    district_id = $4
WHERE
    neighbourhood_id = $1
RETURNING
    *;

-- name: NeighbourhoodFindForUpdate :one
SELECT
    neighbourhood_id,
    neighbourhood_name,
    neighbourhood_code,
    district_id
FROM
    places_schema.neighbourhoods
WHERE
    neighbourhood_id = $1;

