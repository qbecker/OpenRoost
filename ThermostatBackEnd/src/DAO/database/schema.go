package database

const Schema1 = `CREATE TABLE IF NOT EXISTS 'Settings' (
	'setTempHome'	INTEGER NOT NULL, 'setTempAway' INTEGER NOT NULL, 'zipCode' INTEGER NOT NULL, 'setCurrentTemp' INTEGER NOT NULL
); `

const Schema2 = ` CREATE TABLE IF NOT EXISTS 'DAILYDATA' ('Timestamp' TEXT, 'Reading' INTEGER); `
