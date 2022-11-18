CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE Table if NOT exists admin(
  id serial PRIMARY key,
  username TEXT,
  password TEXT
);
CREATE Table if not exists moderator(
  id serial PRIMARY key,
  username TEXT,
  password TEXT
);

INSERT INTO admin(username,password) VALUES('asliddin','compos1995');
INSERT INTO moderator(username,password) VALUES('asliddin','compos1995');
