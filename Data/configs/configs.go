package configs

//Database configs
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
    details TEXT,
    available BOOLEAN NOT NULL,
    status TEXT NOT NULL
  );`

//Server configs
const Port string = ":10000"

//Admin configs
const AdminEmail string = "admin@example.com"
const AdminPassword string = "ExamplePassword"
