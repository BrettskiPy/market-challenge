package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Define the produce structure
type Produce struct {
	Name string `json:"name"`
	Code string `json:"code"`
	Unit_Price string `json:"unit_price"`
}

// Init produce var as a Produce slice
var produce []Produce

// Get all produce
func getProduce(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produce)
}

func initializeProduceData(){
	// Hardcoded data - @todo: add database
	produce = append(produce, Produce{Name: "Lettuce", Code: "A12T-4GH7-QPL9-3N4M", Unit_Price: "3.46"})
	produce = append(produce, Produce{Name: "Peach", Code: "E5T6-9UI3-TH15-QR88", Unit_Price: "2.99"})
	produce = append(produce, Produce{Name: "Green Pepper", Code: "YRT6-72AS-K736-L4AR", Unit_Price: "0.79"})
	produce = append(produce, Produce{Name: "Gala Apple", Code: "TQ4C-VV6T-75ZX-1RMR", Unit_Price: "3.59"})
}


func main() {
	// init router
	router := mux.NewRouter()
	// init starting produce data
	initializeProduceData()
	// Route handles & endpoints
	router.HandleFunc("/produce", getProduce).Methods("GET")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", router))
}
