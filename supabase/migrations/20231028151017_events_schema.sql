
  
CREATE TABLE request_statuses(
    request_status_id serial PRIMARY KEY,
    request_status character varying(200) NOT NULL UNIQUE
);


CREATE TABLE events(
    event_id serial PRIMARY KEY,
    event_name character varying(200) NOT NULL UNIQUE,
    event_location character varying(200) NOT NULL,
    event_location_url character varying(200) NOT NULL,
    constructor_name character varying(200) NOT NULL,
    constructor_title character varying(200) NOT NULL,
    constructor_image character varying(200) NOT NULL,
    event_plan jsonb[]  NOT NULL,
    event_goals varchar[] NOT NULL,
    event_breif text,
    event_description text,
    event_video text,
    event_date date,
    event_start_time time,
    event_end_time time,
    price real not null,
    discount real,
    shab_discount real,
    tags varchar[],
     event_hours int,
    category_id int not null,
    event_image character varying(200),
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp,
    deleted_at timestamp
); 

CREATE TABLE event_requests(
    event_request_id serial PRIMARY KEY,
    request_status_id int not null DEFAULT 1,
    event_id int not null,
    user_name character varying(200) NOT NULL,
    user_email character varying(200) NOT NULL,
    user_phone character varying(200) NOT NULL, 
    city character varying(200), 
    job_title character varying(200),
    price real not null,
    discount real,
    is_shab boolean DEFAULT false,
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp,
    deleted_at timestamp
); 
-- Alter tables within the users schema
ALTER TABLE categories
    ADD FOREIGN KEY (category_type_id) REFERENCES category_types(category_type_id) ;
 

 -- Alter tables within the users schema
ALTER TABLE events
    ADD FOREIGN KEY (category_id) REFERENCES categories(category_id) ;
 
 -- Alter tables within the users schema
ALTER TABLE event_requests
    ADD FOREIGN KEY (request_status_id) REFERENCES request_statuses(request_status_id),
    ADD FOREIGN KEY (event_id) REFERENCES events(event_id);
 