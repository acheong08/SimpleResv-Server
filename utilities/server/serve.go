package server

// Native
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
  "encoding/json"
)

// Third party
import (
	"github.com/gorilla/mux"
)

// Local
import (
	data "github.com/acheong08/SimpleResv-Server/Data"
	configs "github.com/acheong08/SimpleResv-Server/Data/configs"
	db "github.com/acheong08/SimpleResv-Server/utilities/database"
)

// Default
func getItems(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, db.GetItems())
}

// Admin
func admin(w http.ResponseWriter, r *http.Request) {
	// Read post body
	reqBody, _ := ioutil.ReadAll(r.Body)
	// Convert JSON to struct
	var request data.Request
	json.Unmarshal(reqBody, &request)
	// Authenticate
	if request.Email != "admin@example.com" {
		fmt.Fprintf(w, "Not admin")
	} else if !db.AuthUser(request.Email, request.Password) {
		fmt.Fprintf(w, "Authentication failed")
	} else {
		// Proceed if authentication succeeds
		// Check action (AddItem, DeleteItem, AddUser, DeleteUser, ResetDB)
		var status bool
		switch request.Action {
		case "AddItem":
			status = db.AddItem(request.Name, request.Details)
    case "DeleteItem":
      status = db.DeleteItem(request.Id)
    case "AddUser":
      status = db.AddUser(request.AddEmail, request.AddPassword)
    case "DeleteUser":
      status = db.DeleteUser(request.AddEmail)
    case "Reset":
      status = db.ResetDB()
		default:
			status = false
		}
		if status == false {
			fmt.Fprintf(w, "Server Error")
		} else {
			fmt.Fprintf(w, "Success")
		}
	}
}

// User
func user(w http.ResponseWriter, r *http.Request) {
  // Read post body
	reqBody, _ := ioutil.ReadAll(r.Body)
  // Convert JSON to struct
	var request data.Request
	json.Unmarshal(reqBody, &request)
  // Authenticate
  if !db.AuthUser(request.Email, request.Password) {
    fmt.Fprintf(w, "Authentication failed")
  } else{
    switch request.Action {
    case "ToggleItem":
      if db.ToggleItem(request.Id, request.Available) && db.StatusItem(request.Id, request.Status) {
				fmt.Fprintf(w, "Failed")
      } else{
				fmt.Fprintf(w, "Done")
			}
    }
  }
}
func checkAuth(w http.ResponseWriter, r *http.Request) {
  // Read post body
	reqBody, _ := ioutil.ReadAll(r.Body)
  // Convert JSON to struct
	var request data.Request
	json.Unmarshal(reqBody, &request)
  if db.AuthUser(request.Email, request.Password){
    fmt.Fprintf(w, "true")
  } else {
    fmt.Fprintf(w, "false")
  }
}

// Handler
func handleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/api/GetItems", getItems).Methods(http.MethodGet)
	r.HandleFunc("/api/Admin", admin).Methods(http.MethodPost)
  r.HandleFunc("/api/User", user).Methods(http.MethodPost)
  r.HandleFunc("/api/CheckAuth", checkAuth).Methods(http.MethodPost)
	fmt.Println("Server started")
	log.Fatal(http.ListenAndServe(configs.Port, r))
}

func Run() {
	fmt.Println("Listening on " + string(configs.Port))
	handleRequests()
}
