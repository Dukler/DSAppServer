CREATE TABLE users_app
(
    user_id UUID NOT NULL,
    app_id UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT users_app_ids PRIMARY KEY (user_id, app_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE,
    FOREIGN KEY (app_id) REFERENCES app(id) ON UPDATE CASCADE
);