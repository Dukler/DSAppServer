CREATE TABLE domains
(
    id UUID NOT NULL DEFAULT uuid_generate_v1() ,
    name text UNIQUE ,
    app_id UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (app_id) REFERENCES app(id) ON UPDATE CASCADE,
    CONSTRAINT domain_id PRIMARY KEY ( id )
);