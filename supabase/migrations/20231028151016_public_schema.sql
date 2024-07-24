
CREATE TABLE tags(
    tag varchar NOT NULL UNIQUE
);

CREATE TABLE category_types(
    category_type_id serial PRIMARY KEY,
    category_type character varying(200) NOT NULL UNIQUE
);


CREATE TABLE categories(
    category_id serial PRIMARY KEY,
    category_type_id int not null,
    category_name character varying(200) NOT NULL,
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp,
    deleted_at timestamp
); 
CREATE TABLE setting_types(
    setting_type_id serial PRIMARY KEY,
    setting_type character varying(20) NOT NULL UNIQUE
);

CREATE TABLE settings(
    setting_id serial PRIMARY KEY,
    setting_type_id int NOT NULL,
    setting_key character varying(100) NOT NULL UNIQUE,
    setting_value text NOT NULL,
    updated_at timestamp

);
 
CREATE TABLE services(
    service_id serial PRIMARY KEY,
    "service_name" character varying(100) NOT NULL UNIQUE,
    service_image character varying(200) NOT NULL,
    breif   text,
     created_at timestamp NOT NULL DEFAULT NOW(),
    tags varchar[],
    updated_at timestamp,
    deleted_at timestamp
);
 
CREATE TABLE programs(
    program_id serial PRIMARY KEY,
    "program_name" character varying(100) NOT NULL UNIQUE,
    program_image character varying(200) NOT NULL,
    breif   text,
     created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp,
    deleted_at timestamp
);
 
CREATE TABLE projects(
    project_id serial PRIMARY KEY,
    "project_name" character varying(100) NOT NULL UNIQUE,
    project_image character varying(200) NOT NULL,
    category_id int not null,
     created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp,
    deleted_at timestamp
);

CREATE TABLE team_members(
    team_member_id serial PRIMARY KEY,
    "team_member_name" character varying(100) NOT NULL UNIQUE,
    team_member_image character varying(200) NOT NULL,
    job_title character varying(200) NOT NULL,
     created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp,
    deleted_at timestamp
);
 
 CREATE TABLE testemonials(
    testemonial_id serial PRIMARY KEY,
    testemonial_name character varying(100) NOT NULL UNIQUE,
    testemonial_image character varying(200),
    testemonial_title character varying(200) NOT NULL,
    breif text NOT NULL,
     created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp,
    deleted_at timestamp
);


-- Alter tables within the "users" schema
ALTER TABLE settings
    ADD FOREIGN KEY (setting_type_id) REFERENCES setting_types(setting_type_id);
 
