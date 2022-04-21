package main

import (
	"fmt"
	db "github.com/acheong08/SimpleResv-Server/utilities/database"
	server "github.com/acheong08/SimpleResv-Server/utilities/server"
)

func main() {
	// Test database
	db.ResetDB()
	// Test authentication
	if !db.AddUser("admin@example.com", "password") {
		fmt.Println("Failed to add user")
	}
	authed := db.AuthUser("admin@example.com", "wrong password")
	if authed {
		fmt.Println("Authenticated")
	} else {
		fmt.Println("Failed")
	}
	db.DeleteUser("admin@example.com")
	// Test items
	db.AddItem("Macbook")
	db.AddItem("iPad")
	var items string = db.GetItems()
  fmt.Println(items)
  server.Run()
}
