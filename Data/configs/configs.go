package configs

const DBpath string = "Data/database.db"
const Schema string = `
  CREATE TABLE IF NOT EXISTS accounts (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
    email TEXT UNIQUE NOT NULL,
    hash TEXT NOT NULL
  );
  CREATE TABLE IF NOT EXISTS items (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
    name TEXT UNIQUE NOT NULL,
    available BOOLEAN NOT NULL,
    status TEXT NOT NULL
  );`
const Port string = ":10000"
