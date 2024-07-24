 
-- name: EventRequestCreate :one
INSERT INTO event_requests(
    event_id,
user_name,
user_email,
user_phone,
city,
price,
is_shab,
discount )
    VALUES ($1, $2,$3, $4 , $5, $6,$7 , $8)
RETURNING
    *;
 
-- name: EventRequestUpdate :one
UPDATE
    event_requests
SET
    request_status_id = $2
WHERE
    event_request_id = $1
RETURNING
    *;

-- name: EventRequestDeleteRestore :exec
UPDATE
    event_requests
SET
    deleted_at = IIF(deleted_at IS NULL , now() , null)
WHERE
    event_id = ANY (sqlc.arg('event_request_ids')::int[]);
 
 -- name: EventRequestsList :many
 SELECT
        e.event_id,
        r.event_request_id,
        r.user_name,
        r.user_email,
        r.user_phone,
        r.city,
        r.is_shab,
        r.price,
        r.discount,
        e.event_name,
        e.constructor_name,
        e.event_image,
        c.category_id,
        c.category_name,
        s.request_status_id,
        s.request_status,
        r.created_at,
        r.updated_at,
        r.deleted_at,
        e.event_date
    FROM
        event_requests r
        JOIN request_statuses s ON r.request_status_id = s.request_status_id
        JOIN events e ON r.event_id = e.event_id
        JOIN categories c ON e.category_id = c.category_id;
 