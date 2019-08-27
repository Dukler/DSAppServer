CREATE TABLE app
(
    id UUID NOT NULL DEFAULT uuid_generate_v1() ,
    name text UNIQUE ,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    app_type_id integer,
    CONSTRAINT app_id PRIMARY KEY ( id ),
    foreign key (app_type_id) references app_type(id) on update cascade
);
