package database

const Schema1 = `CREATE TABLE IF NOT EXISTS 'Settings' (
	'setTempHome'	INTEGER, 'setTempAway' INTEGER, 'zipCode' INTEGER, 'setCurrentTemp' INTEGER, 'APPNAME' TEXT NOT NULL PRIMARY KEY
); `

const Schema2 = ` CREATE TABLE IF NOT EXISTS 'DAILYDATA' ('Timestamp' TEXT, 'Reading' INTEGER); `

const Schema3 = ` INSERT INTO Settings(APPNAME) VALUES('ROOST');`
