package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math"
	"net/http"
	"regexp"
	"strings"
)

// Define the produce structure
type Produce struct {
	Name string `json:"name"`
	Code string `json:"code"`
	Unit_Price float64 `json:"unit_price"`
}

// Init produce var as a Produce slice
type ProduceList []Produce

// global var where all produce is kept
var produce ProduceList


// Get all produce
func getProduce(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produce)
}

// Get a single produce item
func getSingleProduce(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through produce and find one with the name from the params
	for _, item := range produce {
		if strings.ToLower(item.Name) == strings.ToLower(params["name"]) {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// Add a new produce item
func addProduce(w http.ResponseWriter, r *http.Request) {
	// Accept json and decode it into a slice of structs
	var newProduceItems ProduceList
	err := json.NewDecoder(r.Body).Decode(&newProduceItems)
	if err != nil {
		log.Panic(err)
	}

	var tempItems ProduceList
	re := regexp.MustCompile("^[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}-[a-zA-Z0-9]{4}$")

	// Iterate over each element in the POSTed JSON and validate. If not validated, error out and stop
	for idx, produceItem := range newProduceItems {
		if !re.MatchString(produceItem.Code) || len(produceItem.Name) <= 0 {
			errMsg := fmt.Sprintf("Item %d: Incorrect produce code sequence or product name. Example code sequence: A12T-4GH7-QPL9-3N4M", idx)
			http.Error(w, errMsg, http.StatusBadRequest)
			return
		}

		produceItem.Unit_Price = math.Round(produceItem.Unit_Price*100) / 100 //rounds to the nearest cent
		tempItems = append(tempItems, produceItem)
	}

	// after validation, append new items to the global accumulator and respond back with added items
	produce = append(produce, tempItems...)
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(newProduceItems); err != nil {
		log.Panic(err)
	}
}

// Delete produce item
func deleteProduceItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range produce {
		if item.Name == params["name"] {
			produce = append(produce[:index], produce[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(produce)
}

func initializeProduceData(){
	// Hardcoded data - @todo: add database
	produce = append(produce, Produce{Name: "Lettuce", Code: "A12T-4GH7-QPL9-3N4M", Unit_Price: 3.46})
	produce = append(produce, Produce{Name: "Peach", Code: "E5T6-9UI3-TH15-QR88", Unit_Price: 2.99})
	produce = append(produce, Produce{Name: "Green Pepper", Code: "YRT6-72AS-K736-L4AR", Unit_Price: 0.79})
	produce = append(produce, Produce{Name: "Gala Apple", Code: "TQ4C-VV6T-75ZX-1RMR", Unit_Price: 3.59})
}


func main() {
	// init router
	router := mux.NewRouter()
	// init starting produce data
	initializeProduceData()
	// Route handles & endpoints
	router.HandleFunc("/produce", getProduce).Methods("GET")
	router.HandleFunc("/produce/{name}", getSingleProduce).Methods("GET")
	router.HandleFunc("/produce", addProduce).Methods("POST")
	router.HandleFunc("/produce/{name}", deleteProduceItem).Methods("DELETE")
	// Start server
	log.Fatal(http.ListenAndServe(":8000", router))
}
