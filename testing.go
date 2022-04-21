package main

import (
  db "github.com/acheong08/SimpleResv-Server/utilities/database"
  _ "github.com/acheong08/SimpleResv-Server/utilities/server"
  "fmt"
)

func main()  {
  db.ResetDB()
  if !db.AddUser("acheong@student.dalat.org", "Acheong08$$") {
    fmt.Println("Failed to add user")
  }
  authed := db.AuthUser("acheong@student.dalat.org", "Acheong08$$")
  if authed {
    fmt.Println("Authenticated")
  } else{
    fmt.Println("Failed")
  }
}
