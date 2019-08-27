CREATE TABLE app_types
(
    id SERIAL ,
    name text UNIQUE ,
    CONSTRAINT app_types_id PRIMARY KEY ( id )
);
