-- migrate:up
create table sensors (
  id INTEGER NOT NULL PRIMARY KEY,
  serial INTEGER NOT NULL UNIQUE,
  label VARCHAR(1024)
);

-- migrate:down
drop table sensors;
