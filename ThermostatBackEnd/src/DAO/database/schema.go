package database

const Schema1 = `CREATE TABLE IF NOT EXISTS 'Settings' (
	'setTempHome'	INTEGER, 'setTempAway' INTEGER, 'zipCode' INTEGER, 'setCurrentTemp' INTEGER, 'APPNAME' TEXT NOT NULL PRIMARY KEY,
		'isHome' INTEGER, 'currentTemp' INTERGER
);
CREATE TABLE IF NOT EXISTS 'DAILYDATA' ('Timestamp' TEXT, 'Reading' INTEGER);
 INSERT OR IGNORE INTO Settings(APPNAME) VALUES('ROOST');
 CREATE TRIGGER IF NOT EXISTS update_current_temp AFTER INSERT ON DAILYDATA
BEGIN
UPDATE Settings SET currentTemp = (SELECT Reading FROM DAILYDATA
Order By Timestamp desc limit 1) WHERE APPNAME = 'ROOST';
END;`
