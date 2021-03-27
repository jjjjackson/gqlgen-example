-- +migrate Up
CREATE TABLE users(
  uid CHAR(36) NOT NULL,
  email VARCHAR(255) NOT NULL,
  password VARCHAR(60) NOT NULL,
  salt CHAR(6) NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  PRIMARY KEY(uid),
  UNIQUE(email)
);


-- +migrate Down
DROP TABLE users;
