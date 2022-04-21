package database

//Native
import (
	"crypto/sha512"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"log"
	"os"
)
//Local
import (
	configs "github.com/acheong08/SimpleResv-Server/Data/configs"
	data "github.com/acheong08/SimpleResv-Server/Data"
)
//Drivers
import (
	_ "github.com/mattn/go-sqlite3"
)

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
func DeleteUser(email string) bool {
	//Check if email already used
	if !userexists(email) {
		return false
	}
	// Open database
	db, err := sql.Open("sqlite3", configs.DBpath)
	if err != nil {
		log.Fatal(err)
	}
	// Put email and password into database
	sqlcmd := `DELETE FROM accounts WHERE email = ?`
	statement, err := db.Prepare(sqlcmd)
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec(email)
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

///////////////////////////////////////// Booking system ////////////////////////////////////////////////
func AddItem(name string) {
	// Open database
	db, err := sql.Open("sqlite3", configs.DBpath)
	if err != nil {
		log.Fatal(err)
	}
	// Insert new entry
	sqlcmd := `INSERT INTO items(name, available, status) VALUES (?, ?, ?)`
	statement, err := db.Prepare(sqlcmd)
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec(name, true, "Available")
	if err != nil {
		log.Fatal(err)
	}
}
func DeleteItem(id int) {
	// Open database
	db, err := sql.Open("sqlite3", configs.DBpath)
	if err != nil {
		log.Fatal(err)
	}
	// Delete by ID
	sqlcmd := `DELETE FROM items WHERE id = ?`
	statement, err := db.Prepare(sqlcmd)
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
}
func ToggleItem(id int, available bool) {
	// Open database
	db, err := sql.Open("sqlite3", configs.DBpath)
	if err != nil {
		log.Fatal(err)
	}
	sqlcmd := `UPDATE items SET available = ? WHERE id = ?`
	statement, err := db.Prepare(sqlcmd)
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec(available, id)
	if err != nil {
		log.Fatal(err)
	}
}
func StatusItem(id int, status string) {
	// Open database
	db, err := sql.Open("sqlite3", configs.DBpath)
	if err != nil {
		log.Fatal(err)
	}
	sqlcmd := `UPDATE items SET status = ? WHERE id = ?`
	statement, err := db.Prepare(sqlcmd)
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec(status, id)
	if err != nil {
		log.Fatal(err)
	}
}
func GetItems() string {
	// Open database
	db, err := sql.Open("sqlite3", configs.DBpath)
	if err != nil {
		log.Fatal(err)
	}
	// Get all data
	rows, err := db.Query("SELECT * FROM items")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	//Slice to hold items
	var items []data.Item
	// Loop through rows and extract data
	for rows.Next() {
		var item data.Item
		if err := rows.Scan(&item.Id, &item.Name, &item.Available, &item.Status); err != nil {
			log.Fatal(err)
		}
		items = append(items, item)
	}
	result, err := json.Marshal(items)
	return string(result)
}
