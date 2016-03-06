-- Database: tododb

-- DROP DATABASE tododb;

CREATE DATABASE tododb
  WITH OWNER = postgres
       ENCODING = 'UTF8'
       TABLESPACE = pg_default
       LC_COLLATE = 'en_US.UTF-8'
       LC_CTYPE = 'en_US.UTF-8'
       CONNECTION LIMIT = -1;

-- Schema: public

-- DROP SCHEMA public;

CREATE SCHEMA public
  AUTHORIZATION postgres;

GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO public;
COMMENT ON SCHEMA public
  IS 'standard public schema';

-- Table: todo_webapi_todoitem

-- DROP TABLE todo_webapi_todoitem;

CREATE TABLE todo_webapi_todoitem
(
  id serial NOT NULL,
  content character varying(1024) NOT NULL,
  created_time timestamp with time zone NOT NULL,
  finished_time timestamp with time zone NOT NULL,
  finished boolean NOT NULL,
  CONSTRAINT todo_webapi_todoitem_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE todo_webapi_todoitem
  OWNER TO postgres;
