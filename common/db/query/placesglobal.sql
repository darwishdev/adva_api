-- name: PlacesList :many
SELECT
    Row_to_json(place_record) record
FROM (
    SELECT
        *
    FROM
        places_schema.places_view) place_record;

