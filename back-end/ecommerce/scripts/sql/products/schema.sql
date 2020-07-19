CREATE EXTENSION pgcrypto;

CREATE TABLE products (
  id    UUID PRIMARY KEY  DEFAULT gen_random_uuid(),
  name VARCHAR(100)      NOT NULL UNIQUE,
  description  VARCHAR(100)      NOT NULL,
  quantity   int    DEFAULT 0,
  price      NUMERIC(15,6)     DEFAULT 0.00,
  created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'utc'),
  updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT (NOW() AT TIME ZONE 'utc')
);