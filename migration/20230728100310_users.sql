-- migrate:up
  CREATE TABLE IF NOT EXISTS users (
    id VARCHAR(128) NOT NULL,
    organization_id VARCHAR(128) NOT NULL,
    name VARCHAR(255),
    role VARCHAR(64),
    email VARCHAR(255),
    password VARCHAR(255),
    address TEXT,
    photo TEXT,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,
    
    CONSTRAINT users_id_pkey PRIMARY KEY (id),
    FOREIGN KEY (organization_id) REFERENCES organizations(id)
  );

-- migrate:down
  DROP TABLE IF EXISTS users;