 
CREATE TABLE blogs (
  blog_id serial PRIMARY KEY ,
  blog_name VARCHAR(255) NOT NULL,
  blog_image VARCHAR(255),
  breif TEXT NOT NULL,
  content TEXT NOT NULL,
  category_id INT NOT NULL,
  user_id INT NOT NULL,
  views INT DEFAULT 0,
  date_time timestamp NOT NULL,
  links JSON,
  tags varchar[],
  created_at timestamp NOT NULL DEFAULT NOW(),
  confirmed_at timestamp,
  updated_at timestamp,
  deleted_at timestamp
);