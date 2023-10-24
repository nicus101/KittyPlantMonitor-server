CREATE TABLE IF NOT EXISTS "schema_migrations" (version varchar(128) primary key);
CREATE TABLE sensors (
  id INTEGER NOT NULL PRIMARY KEY,
  serial INTEGER NOT NULL UNIQUE,
  label VARCHAR(1024)
);
-- Dbmate schema migrations
INSERT INTO "schema_migrations" (version) VALUES
  ('20231024162705');
