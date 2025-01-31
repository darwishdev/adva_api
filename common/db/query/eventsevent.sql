 
-- name: EventCreate :one
INSERT INTO events(
event_name,
event_location,
event_location_url,
constructor_name,
constructor_title,
constructor_image,
event_plan,
event_goals,
event_breif,
event_description,
event_video,
event_date,
event_start_time,
event_end_time,
 event_hours,
category_id,
shab_discount,
price,
discount,
event_image,
tags)
    VALUES ($1, $2,$3, $4 , $5, $6,$7, $8 , $9, $10,$11, $12,$13, $14,$15, $16 , $17, $18 , $19 , $20 , sqlc.arg('tags')::varchar[])
RETURNING
    *;
 
-- name: EventUpdate :one
UPDATE
    events
SET
    event_name = $2,
    event_location = $3,
    event_location_url = $4,
    constructor_name = $5,
    constructor_title = $6,
    constructor_image = $7,
    event_plan = $8,
    event_goals = $9,
    event_breif = $10,
    event_description = $11,
    event_video = $12,
    event_date = $13,
    event_start_time = $14,
    event_end_time = $15,
    event_hours = $16,
    category_id = $17,
    event_image = $18,
    price = $19,
    discount = $20,
    shab_discount = $21,
    tags = sqlc.arg('tags')::varchar[]
WHERE
    event_id = $1
RETURNING
    *;

-- name: EventDeleteRestore :exec
UPDATE
    events
SET
    deleted_at = IIF(deleted_at IS NULL , now() , null)
WHERE
    event_id = ANY (sqlc.arg('event_ids')::int[]);

-- name: EventFind :one
    SELECT
        e.event_id,
        e.event_name,
        e.event_location,
        e.event_location_url,
        e.constructor_name,
        e.tags,
        e.constructor_title,
        e.constructor_image,
        e.event_plan,
        e.event_goals,
        c.category_name,
        e.event_breif,
        e.event_description,
        e.event_video,
        e.event_date,
        e.price,
        e.shab_discount,
        e.discount,
        e.event_start_time,
        e.event_end_time,
        e.event_hours,
        e.category_id,
        e.event_image,
        e.created_at,
        e.updated_at,
        e.deleted_at
    FROM
        events e
        JOIN categories c ON e.category_id = c.category_id
    WHERE
        e.event_id = $1 ;
 
 -- name: EventsList :many
 SELECT
        e.event_id,
        e.event_name,
        e.event_location,
        e.event_location_url,
        e.constructor_name,
        e.tags,
        e.constructor_title,
        e.constructor_image,
        e.event_plan,
        e.event_goals,
        e.shab_discount,
        c.category_name,
        e.event_breif,
        e.event_description,
        e.event_video,
        e.event_date,
        e.price,
        e.discount,
        e.event_start_time,
        e.event_end_time,
        e.event_hours,
        e.category_id,
        e.event_image,
        e.created_at,
        e.updated_at,
        e.deleted_at
    FROM
        events e
        JOIN categories c ON e.category_id = c.category_id;

-- name: EventsInputList :many
SELECT
    event_id,
    event_name
FROM
    events
WHERE
    deleted_at IS NULL;






 -- name: EventFindForUpdate :one
 SELECT
 event_id,
 event_name,
        event_location,
        price,
        tags,
        discount,
        shab_discount,
        event_location_url,
        constructor_name,
        constructor_title,
        constructor_image,
        event_plan,
        event_goals,
        event_breif,
        event_description,
        event_video,
        event_date,
        event_start_time,
        event_end_time,
         event_hours,
        category_id,
        event_image
    FROM
        events where event_id = $1;