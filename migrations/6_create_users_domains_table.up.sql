CREATE TABLE users_domains
(
    user_id UUID NOT NULL,
    domain_id UUID NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT users_domain_ids PRIMARY KEY (user_id, domain_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE,
    FOREIGN KEY (domain_id) REFERENCES domains(id) ON UPDATE CASCADE
);