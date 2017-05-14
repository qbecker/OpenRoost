package database

const Schema1 = `CREATE TABLE IF NOT EXISTS 'Settings' (
	'setTempHome'	INTEGER, 'setTempAway' INTEGER, 'zipCode' INTEGER, 'setCurrentTemp' INTEGER, 'APPNAME' TEXT NOT NULL PRIMARY KEY,
		'isHome' INTEGER, 'currentTemp' INTERGER
); `

const Schema2 = ` CREATE TABLE IF NOT EXISTS 'DAILYDATA' ('Timestamp' TEXT, 'Reading' INTEGER); `

const Schema3 = ` INSERT INTO Settings(APPNAME) VALUES('ROOST');`

const Schema4 = `CREATE TRIGGER update_current_temp AFTER INSERT ON DAILYDATA
BEGIN
UPDATE Settings SET currentTemp = (SELECT Reading FROM DAILYDATA
Order By Timestamp desc limit 1) WHERE APPNAME = 'ROOST';
END`
