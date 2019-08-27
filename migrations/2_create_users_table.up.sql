CREATE TABLE users
(
    id UUID NOT NULL DEFAULT uuid_generate_v1() ,
    email text UNIQUE ,
    username text ,
    password text ,
    token text,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT users_id PRIMARY KEY ( id )
);