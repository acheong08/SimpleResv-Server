package database

import (
	"database/sql"
	_ "fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
  configs "github.com/acheong08/SimpleResv/Data"
	"crypto/sha512"
	"encoding/hex"
	"os"
)

func ResetDB() {
	const schema string = `
    CREATE TABLE IF NOT EXISTS accounts (
      id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
      email TEXT NOT NULL,
      hash TEXT NOT NULL
    );
    CREATE TABLE IF NOT EXISTS entries (
      id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
      name TEXT NOT NULL,
      status BOOLEAN NOT NULL
    );
		INSERT INTO accounts (email, hash) VALUES ("admin@duti.tech", "83aa8d9ae9c7a057be1e839d27811e83b16e839fff72c9c3ab6d13ab1a7c57edcf8977cc1634c91a863c4eccd03760796e9e27d6e163151ba7ca7137ccf0ff79")`
		os.Remove(configs.DBpath)
		db, err := sql.Open("sqlite3", configs.DBpath)
		if err != nil {
			log.Fatal(err)
		}
		if _, err := db.Exec(schema); err != nil {
			log.Fatal(err)
		}
		db.Close()

}
func AddUser(username string, password string) {
	// Hash password and insert into database
	pwdhash := sha512.Sum512([]byte(password))
	var strhash string = hex.EncodeToString(pwdhash[:])
	db, err := sql.Open("sqlite3", configs.DBpath)
	if err != nil {
		log.Fatal(err)
	}
	sqlcmd := `INSERT INTO accounts(email, hash) VALUES (?,?)`
	statement, err := db.Prepare(sqlcmd)
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec(username, strhash)
	db.Close()
}
