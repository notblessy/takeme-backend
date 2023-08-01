-- migrate:up
  CREATE TABLE IF NOT EXISTS subscriptions (
    id BIGINT AUTO_INCREMENT NOT NULL,
    user_id VARCHAR(128),
    subscription_plan_id INT,
    is_active BOOLEAN,
    expired_at TIMESTAMP NULL,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NULL,
    
    CONSTRAINT subscriptions_id_pkey PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (subscription_plan_id) REFERENCES subscription_plans(id)
  );

-- migrate:down
  DROP TABLE IF EXISTS subscriptions;