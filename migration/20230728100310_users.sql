-- migrate:up
  CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(128) NOT NULL,
    email VARCHAR(255),
    password VARCHAR(255),
    name VARCHAR(255),
    description TEXT,
    gender INT,
    age INT,
    is_premium TEXT,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,
    
    CONSTRAINT users_id_pkey PRIMARY KEY (id)
  );

-- migrate:down
  DROP TABLE IF EXISTS users;