CREATE EXTENSION pgcrypto;

CREATE TABLE users (
  user_id    UUID PRIMARY KEY  DEFAULT gen_random_uuid(),
  full_name  VARCHAR(100)      NOT NULL UNIQUE,
  nick_name  VARCHAR(50)       NOT NULL,
  email      VARCHAR(50)       NOT NULL UNIQUE,
  balance    NUMERIC(15,6)     DEFAULT 10000.00
);

CREATE TABLE items (
  item_id     UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name        VARCHAR(80)      NOT NULL,
  description TEXT
);

CREATE TABLE auctions (
  auction_id     UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  minimum_bid NUMERIC(15,6)     NOT NULL,
  buy_out     NUMERIC(15,6)     NOT NULL,
  user_id     UUID REFERENCES  users(user_id),
  item_id     UUID REFERENCES  items(item_id)
);