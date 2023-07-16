package subscribe

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from sub node!")
}

func display(w http.ResponseWriter, r *http.Request) {

}

func StartPubNode() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", ping)

	log.Println("Subscribe is running on port 3001")
	err := http.ListenAndServe(":3001", r)
	if err != nil {
		log.Fatal(err)
	}
}
