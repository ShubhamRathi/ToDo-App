package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Item struct {
	gorm.Model // What does this mean?
	Name  string
}

func AllItems(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	var items [] Item
	db.Find(&items)
	fmt.Println("{}", items)

	json.NewEncoder(w).Encode(items)
}

func AddItem(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	item := vars["text"]

	db.Create(&Item{Name: item})
	fmt.Fprintf(w, "\nNew Item Successfully Added")
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	item := vars["text"]

	var del Item
	db.Where("name = ?", item).Find(&del)
	db.Delete(&del)

	fmt.Fprintf(w, "Successfully Deleted Item")

}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Update Endpoint Hit")
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	vars := mux.Vars(r)
	item := vars["text"]
	update := vars["update"]

	var locate Item
	db.Where("name = ?", item).Find(&locate)

	locate.Name = update
	db.Save(&locate)

	fmt.Fprintf(w, "Updated Item")
}

func initialMigration() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Item{})
}

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/items", AllItems).Methods("GET")
	myRouter.HandleFunc("/items/{text}", DeleteItem).Methods("DELETE")
	myRouter.HandleFunc("/items/{text}/{update}", UpdateItem).Methods("PUT")
	myRouter.HandleFunc("/items/{text}", AddItem).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}


func main() {
	fmt.Println("Go To Do REST API")
	// Add the call to our new initialMigration function
	initialMigration()

	// Handle Subsequent requests
	HandleRequests()
}
