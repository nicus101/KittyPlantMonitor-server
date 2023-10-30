CREATE TABLE IF NOT EXISTS "schema_migrations" (version varchar(128) primary key);
CREATE TABLE sensors (
  id INTEGER NOT NULL PRIMARY KEY,
  serial INTEGER NOT NULL UNIQUE,
  label VARCHAR(1024)
);
CREATE TABLE measures (
  id INTEGER NOT NULL PRIMARY KEY,
  value REAL NOT NULL,
  created_at DATETIME NOT NULL DEFAULT  CURRENT_TIMESTAMP,
  sensor_id INTEGER NOT NULL,
  FOREIGN KEY (sensor_id) REFERENCES sensors(id)
);
-- Dbmate schema migrations
INSERT INTO "schema_migrations" (version) VALUES
  ('20231024162705'),
  ('20231024170548');
