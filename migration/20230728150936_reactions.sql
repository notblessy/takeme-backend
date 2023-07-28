-- migrate:up
  CREATE TABLE IF NOT EXISTS reactions (
    id VARCHAR(128) NOT NULL,
    user_by VARCHAR(128),
    user_to VARCHAR(128),
    type VARCHAR(35),

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,
    
    CONSTRAINT reactions_id_pkey PRIMARY KEY (id),
    FOREIGN KEY (user_by) REFERENCES users(id),
    FOREIGN KEY (user_to) REFERENCES users(id)
  );

-- migrate:down
  DROP TABLE IF EXISTS reactions;