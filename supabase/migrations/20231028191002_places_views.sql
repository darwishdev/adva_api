-- view
CREATE OR REPLACE VIEW places_schema.places_view AS
WITH neighbourhoodsRecords AS (
    SELECT
        n.district_id,
(
            SELECT
                Jsonb_agg(nested_neighbourhoods) neighbourhoods
            FROM (
                SELECT
                    neighbourhood_id,
                    neighbourhood_name,
                    neighbourhood_code
                FROM
                    places_schema.neighbourhoods nn
                WHERE
                    nn.district_id = n.district_id) nested_neighbourhoods) neighbourhoods
        FROM
            places_schema.neighbourhoods n
        GROUP BY
            n.district_id
        ORDER BY
            n.district_id
),
districtsRecords AS (
    SELECT
        d.city_id,
(
            SELECT
                Jsonb_agg(nested_districts)
            FROM (
                SELECT
                    nd.district_id,
                    district_name,
                    district_code,
                    nr.neighbourhoods
                FROM
                    places_schema.districts nd
                    JOIN neighbourhoodsRecords nr ON nd.district_id = nr.district_id
                WHERE
                    nd.city_id = d.city_id
                GROUP BY
                    nd.district_id,
                    district_name,
                    district_code,
                    nr.neighbourhoods) nested_districts) districts
        FROM
            places_schema.districts d
        GROUP BY
            d.city_id
)
SELECT
    c.city_id,
    city_name,
    city_code,
    d.districts,
    created_at,
    deleted_at
FROM
    places_schema.cities c
    JOIN districtsRecords d ON c.city_id = d.city_id
