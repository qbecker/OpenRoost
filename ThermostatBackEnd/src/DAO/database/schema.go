package database

const Schema1 = `CREATE TABLE IF NOT EXISTS 'Settings' (
	'setTemp'	INTEGER NOT NULL 
); `

const Schema2 = ` CREATE TABLE IF NOT EXISTS 'DAILYDATA' ('Timestamp' TEXT, 'Reading' INTEGER); `
