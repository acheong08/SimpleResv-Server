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
    details TEXT,
    available BOOLEAN NOT NULL,
    status TEXT NOT NULL
  );
  INSERT INTO accounts(email, hash) VALUES ("admin@example.com", "3958e5ecbc5f55cd623e9986f2ea94f8704f00f17b5a04ed491baefaca2d042c26c7d6b26c0a09508944e4107ee2ec9d427eb8fe84eb73558bcf634fa90ac17e")`
const Port string = ":10000"
