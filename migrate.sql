CREATE TABLE if not EXISTS Waypoints ('id' INTEGER PRIMARY KEY, 'name' TEXT, 'longitude' DOUBLE, 'latitude' DOUBLE, 'altitude' DOUBLE);
CREATE TABLE IF NOT EXISTS `aeac_routes` (`id` INTEGER PRIMARY KEY, `number` INT NOT NULL, `start_waypoint` TEXT NOT NULL, `end_waypoint` TEXT NOT NULL, `passengers` FLOAT NOT NULL, `max_weight` FLOAT NOT NULL, `value` INT NOT NULL, `remarks` TEXT, `odr` INT);
CREATE TABLE IF NOT EXISTS `restrictions` (`id` INTEGER PRIMARY KEY, `bound_1` INT, `bound_2` INT, `bound_3` INT, `bound_4` INT, `rejoin` INT)
