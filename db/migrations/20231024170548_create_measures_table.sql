-- migrate:up
CREATE TABLE measures (
  id INTEGER NOT NULL PRIMARY KEY,
  value REAL NOT NULL,
  created_at DATETIME NOT NULL DEFAULT  CURRENT_TIMESTAMP,
  sensor_id INTEGER NOT NULL,
  FOREIGN KEY (sensor_id) REFERENCES sensors(id)
);

-- migrate:down
drop table measures;
