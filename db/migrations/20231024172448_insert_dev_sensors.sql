-- ESP1 - 54E2053402
-- ESP2 - 54E2053366

-- migrate:up
ALTER TABLE sensors
ADD serial_code VARCHAR(10);

INSERT INTO sensors (serial, serial_code, label)
VALUES (2053402, "54E2053402", "PoC sensor 1"),
(2053366, "54E2053366", "PoC sensor 2");

-- migrate:down
ALTER TABLE sensors
DROP COLUMN serial_code;
