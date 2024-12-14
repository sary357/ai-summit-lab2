CREATE DATABASE job_submit WITH 
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.utf8'
    LC_CTYPE = 'en_US.utf8'
    TABLESPACE = pg_default;

-- Path: sql/create_tables.sql
-- the following table must BE in pg_default tablespace and database "job_submit"
create table job_submit.public.jobs (
    id SERIAL,
    codes text not null,
    requirements_txt text,
    status int not null default 0,
    endpoint varchar(65535),
    lock_executor varchar(65535),
    running_executor varchar(65536),
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp,
    PRIMARY KEY (id)
)  TABLESPACE pg_default;  

create role readonly;
GRANT usage ON schema public TO readonly;
GRANT SELECT ON ALL TABLES IN schema public TO readonly;
GRANT SELECT ON ALL SEQUENCES IN schema public TO readonly;
