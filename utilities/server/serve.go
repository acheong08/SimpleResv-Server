package server

import (
  "log"
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
)

func homepage(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Experimenting")
  fmt.Println("Homepage accessed")
}

func handleRequests()  {
  r := mux.NewRouter()
  r.HandleFunc("/testing", homepage).Methods(http.MethodGet)
  fmt.Println("Server started")
  log.Fatal(http.ListenAndServe(":10000", r))
}

func main()  {
  handleRequests()
}
