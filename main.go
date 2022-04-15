package main

import (
  db "github.com/acheong08/SimpleResv-Server/utilities/database"
  _ "github.com/acheong08/SimpleResv-Server/utilities/server"
)

func main()  {
  db.ResetDB()
  db.AddUser("acheong@student.dalat.org", "apassword")
}
