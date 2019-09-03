CREATE TABLE app
(
    id UUID NOT NULL DEFAULT uuid_generate_v1() ,
    name text UNIQUE ,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT app_id PRIMARY KEY ( id )
);
