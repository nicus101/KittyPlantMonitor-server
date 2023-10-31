-- migrate:up
create table sensors (
  id INTEGER NOT NULL PRIMARY KEY,
  serial_code VARCHAR(32) NOT NULL UNIQUE,
  label VARCHAR(1024)
);

-- migrate:down
drop table sensors;
