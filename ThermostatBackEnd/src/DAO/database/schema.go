package database

const Schema1 = `CREATE TABLE IF NOT EXISTS 'Settings' (
	'setTempHome'	INTEGER, 'setTempAway' INTEGER, 'zipCode' INTEGER, 'setCurrentTemp' INTEGER
); `

const Schema2 = ` CREATE TABLE IF NOT EXISTS 'DAILYDATA' ('Timestamp' TEXT, 'Reading' INTEGER); `
