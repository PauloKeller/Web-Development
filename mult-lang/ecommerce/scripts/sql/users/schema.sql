CREATE EXTENSION pgcrypto;

CREATE TABLE users (
  id    UUID PRIMARY KEY  DEFAULT gen_random_uuid(),
  first_name VARCHAR(100)      NOT NULL,
  last_name  VARCHAR(100)      NOT NULL,
  username   VARCHAR(50)       NOT NULL UNIQUE,
  email      VARCHAR(50)       NOT NULL UNIQUE,
  birthdate TIMESTAMP,
  created_at  TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'utc'),
  updated_at  TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'utc'),
  balance    NUMERIC(15,6)     DEFAULT 0.00
);