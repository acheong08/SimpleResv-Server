package database

import (
	"crypto/sha512"
	"database/sql"
	"encoding/hex"
	_ "fmt"
	configs "github.com/acheong08/SimpleResv-Server/Data"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

/////////////////////////////////////// User functions /////////////////////////////////////////////////
//Private functions
func userexists(email string) bool {
	// Open database
	db, err := sql.Open("sqlite3", configs.DBpath)
	if err != nil {
		log.Fatal(err)
	}
	// Read password from database with same email
	sqlcmd := `SELECT hash FROM accounts WHERE email = ?`
	statement, err := db.Prepare(sqlcmd)
	if err != nil {
		log.Fatal(err)
	}
	//Get result
	var result string
	err = statement.QueryRow(email).Scan(&result)
	//See if result is blank
	if result == "" {
		return false
	} else {
		return true
	}
}

//Public functions
func ResetDB() {
	//Delete database file
	os.Remove(configs.DBpath)
	//Create database and configure based on Schema
	db, err := sql.Open("sqlite3", configs.DBpath)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := db.Exec(configs.Schema); err != nil {
		log.Fatal(err)
	}
	//Close and unlock database
	db.Close()

}
func AddUser(email string, password string) bool {
	//Check if email already used
	if userexists(email) {
		return false
	}
	// Get SHA512 hash of password
	pwdhash := sha512.Sum512([]byte(password))
	var strhash string = hex.EncodeToString(pwdhash[:])
	// Open database
	db, err := sql.Open("sqlite3", configs.DBpath)
	if err != nil {
		log.Fatal(err)
	}
	// Put email and password into database
	sqlcmd := `INSERT INTO accounts(email, hash) VALUES (?,?)`
	statement, err := db.Prepare(sqlcmd)
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec(email, strhash)
	if err != nil {
		log.Fatal(err)
	}
	// Close database
	db.Close()
	return true
}
func AuthUser(email string, password string) bool {
	// Hash password
	pwdhash := sha512.Sum512([]byte(password))
	var strhash string = hex.EncodeToString(pwdhash[:])
	// Open database
	db, err := sql.Open("sqlite3", configs.DBpath)
	if err != nil {
		log.Fatal(err)
	}
	// Read password from database with same email
	sqlcmd := `SELECT hash FROM accounts WHERE email = ?`
	statement, err := db.Prepare(sqlcmd)
	if err != nil {
		log.Fatal(err)
	}
	var result string
	err = statement.QueryRow(email).Scan(&result)
	// Check whether password matches
	if strhash == result {
		return true
	} else {
		return false
	}
}
