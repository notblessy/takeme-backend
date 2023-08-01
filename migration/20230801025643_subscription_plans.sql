-- migrate:up
  CREATE TABLE IF NOT EXISTS subscription_plans (
    id INT AUTO_INCREMENT NOT NULL,
    name VARCHAR(128),
    price INT,
    feature JSON,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,
    
    CONSTRAINT subscription_plans_id_pkey PRIMARY KEY (id)
  );

-- migrate:down
  DROP TABLE IF EXISTS subscription_plans;