-- migrate:up
  CREATE TABLE IF NOT EXISTS notifications (
    id VARCHAR(128) NOT NULL,
    user_by VARCHAR(128),
    user_to VARCHAR(128),
    message TEXT,
    is_read boolean,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,
    
    CONSTRAINT notifications_id_pkey PRIMARY KEY (id),
    FOREIGN KEY (user_by) REFERENCES users(id),
    FOREIGN KEY (user_to) REFERENCES users(id)
  );

-- migrate:down
  DROP TABLE IF EXISTS notifications;