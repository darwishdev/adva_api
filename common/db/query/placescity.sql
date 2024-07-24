-- name: CitiesList :many
WITH ditrictsCount AS (
    SELECT
        city_id,
        COUNT(district_id) districts_count
    FROM
        places_schema.districts
    GROUP BY
        city_id
)
SELECT
    c.city_id,
    city_name,
    city_code,
    d.districts_count,
    created_at,
    deleted_at
FROM
    places_schema.cities c
    LEFT JOIN ditrictsCount d ON c.city_id = d.city_id;

-- name: CitiesInputList :many
SELECT
    c.city_id,
    city_name
FROM
    places_schema.cities c
    JOIN places_schema.districts d ON d.city_id = c.city_id
WHERE
    c.deleted_at IS NULL
GROUP BY
    c.city_id,
    city_name;

-- name: CityDeleteRestore :exec
UPDATE
    places_schema.cities
SET
    deleted_at = CASE WHEN deleted_at IS NULL THEN
        now()
    ELSE
        NULL
    END
WHERE
    city_id = ANY (sqlc.arg('city_ids')::int[]);

-- name: CityCreate :one
INSERT INTO places_schema.cities(city_name, city_code)
    VALUES ($1, $2)
RETURNING
    *;

-- name: CityUpdate :one
UPDATE
    places_schema.cities
SET
    city_name = $2,
    city_code = $3
WHERE
    city_id = $1
RETURNING
    *;

-- name: CityFindForUpdate :one
SELECT
    city_id,
    city_name,
    city_code
FROM
    places_schema.cities
WHERE
    city_id = $1;

