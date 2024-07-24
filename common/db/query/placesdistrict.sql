-- name: DistrictsList :many
WITH neighbourhoodsCount AS (
    SELECT
        district_id,
        COUNT(neighbourhood_id) neighbourhoods_count
    FROM
        places_schema.neighbourhoods
    GROUP BY
        district_id
)
SELECT
    d.district_id,
    district_name,
    district_code,
    n.neighbourhoods_count,
    c.city_name,
    c.city_code,
    c.city_id,
    d.created_at,
    d.deleted_at
FROM
    places_schema.districts d
    JOIN places_schema.cities c ON d.city_id = c.city_id
    LEFT JOIN neighbourhoodsCount n ON n.district_id = d.district_id;

-- name: DistrictsInputList :many
SELECT
    district_id,
    district_name
FROM
    places_schema.districts
WHERE
    city_id = $1
    AND deleted_at IS NULL;

-- name: DistrictDeleteRestore :exec
UPDATE
    places_schema.districts
SET
    deleted_at = CASE WHEN deleted_at IS NULL THEN
        now()
    ELSE
        NULL
    END
WHERE
    district_id = ANY (sqlc.arg('district_ids')::int[]);

-- name: DistrictCreate :one
INSERT INTO places_schema.districts(city_id, district_name, district_code)
    VALUES ($1, $2, $3)
RETURNING
    *;

-- name: DistrictUpdate :one
UPDATE
    places_schema.districts
SET
    district_name = $2,
    district_code = $3,
    city_id = $4
WHERE
    district_id = $1
RETURNING
    *;

-- name: DistrictFindForUpdate :one
SELECT
    district_id,
    district_name,
    district_code,
    city_id
FROM
    places_schema.districts
WHERE
    district_id = $1;

