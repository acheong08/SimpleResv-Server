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
	if !db.AddUser("acheong@student.dalat.org", "Acheong08$$") {
		fmt.Println("Failed to add user")
	}
	authed := db.AuthUser("acheong@student.dalat.org", "Acheong08$$")
	if authed {
		fmt.Println("Authenticated")
	} else {
		fmt.Println("Failed")
	}
	db.DeleteUser("admin@duti.tech")
	// Test items
	db.AddItem("Macbook")
	db.AddItem("iPad")
	var items string = db.GetItems()
  fmt.Println(items)
  server.Run()
}
