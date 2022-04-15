package main

import (
  db "github.com/acheong08/SimpleResv/utilities/database"
  _ "github.com/acheong08/SimpleResv/utilities/server"
)

func main()  {
  db.ResetDB()
  db.AddUser("acheong@student.dalat.org", "apassword")
}
