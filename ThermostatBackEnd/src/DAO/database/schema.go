package database

const Schema = `CREATE TABLE IF NOT EXISTS 'Settings' (
	'setTemp'	INTEGER NOT NULL 
); CREATE TABLE IF NOT EXISTS 'DAILYDATA' ('Timestamp' TEXT, 'Reading' INTEGER);`
