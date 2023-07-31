-- migrate:up
  CREATE TABLE IF NOT EXISTS photos (
    id BIGINT AUTO_INCREMENT NOT NULL,
    user_id VARCHAR(128),
    url TEXT,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,
    
    CONSTRAINT photos_id_pkey PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id)
  );

-- migrate:down
  DROP TABLE IF EXISTS photos;