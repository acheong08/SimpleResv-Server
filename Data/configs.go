package Data

const DBpath string = "Data/database.db"
const Schema string = `
  CREATE TABLE IF NOT EXISTS accounts (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
    email TEXT UNIQUE NOT NULL,
    hash TEXT NOT NULL
  );
  CREATE TABLE IF NOT EXISTS entries (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
    name TEXT UNIQUE NOT NULL,
    status BOOLEAN NOT NULL
  );
  INSERT INTO accounts (email, hash) VALUES ("admin@duti.tech", "83aa8d9ae9c7a057be1e839d27811e83b16e839fff72c9c3ab6d13ab1a7c57edcf8977cc1634c91a863c4eccd03760796e9e27d6e163151ba7ca7137ccf0ff79")`
