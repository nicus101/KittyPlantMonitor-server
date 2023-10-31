-- ESP1 - 54E2053402
-- ESP2 - 54E2053366

-- migrate:up
INSERT INTO sensors (serial_code, label)
VALUES ("54E2053402", "PoC sensor 1"),
("54E2053366", "PoC sensor 2");

-- migrate:down
DELETE FROM sensors WHERE
serial_code = "54E2053402" OR
serial_code = "54E2053366";
