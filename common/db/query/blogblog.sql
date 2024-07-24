 
-- name: BlogCreate :one
INSERT INTO blogs(
blog_name, 
breif, 
blog_image,
content,
category_id,
user_id,
views,
links,
tags,
date_time
)
    VALUES ($1, $2,$3,$4, $5,$6,$7, $8,$9,$10)
RETURNING
    *;
 
-- name: BlogUpdate :one
UPDATE
    blogs
SET
    blog_name = $2, 
	breif = $3, 
	blog_image = $4,
	content = $5,
	category_id = $6,
	user_id = $7,
	views = $8,
	date_time = $9,
	links = $10,
	tags = $11
WHERE
    blog_id = $1
RETURNING
    *;

-- name: BlogDeleteRestore :exec
UPDATE
    blogs
SET
    deleted_at = IIF(deleted_at IS NULL , now() , null)
WHERE
    blog_id = ANY (sqlc.arg('blog_ids')::int[]);
 
 -- name: BlogsList :many
 SELECT
        b.blog_id,
		b.blog_name,
		b.blog_image,
		b.breif,
		b.content,
		b.category_id,
		b.user_id,
		b.views,
		b.date_time,
		b.links,
		b.tags,
		b.created_at,
		b.confirmed_at,
		b.updated_at,
		b.deleted_at,
		c.category_name
    FROM
        blogs b
		join categories c on b.category_id = c.category_id ;



 -- name: BlogFindForUpdate :one
 SELECT
        b.blog_id,
		b.blog_name,
		b.blog_image,
		b.breif,
		b.content,
		b.category_id,
		b.user_id,
		b.views,
		b.date_time,
		b.links,
		b.tags,
		b.created_at,
		b.confirmed_at,
		b.updated_at,
		b.deleted_at,
		c.category_name
    FROM
        blogs b
		join categories c on b.category_id = c.category_id
       where blog_id = $1 ;
 