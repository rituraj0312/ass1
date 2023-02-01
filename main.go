package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error

type persons struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

func createNewPerson(w http.ResponseWriter, r *http.Request) { //router handler
	reqBody, _ := ioutil.ReadAll(r.Body)
	var person persons
	json.Unmarshal(reqBody, &person)
	db.Create(&person)
	fmt.Println("Create New Person")
	json.NewEncoder(w).Encode(person)
}

func main() {
	db, err = gorm.Open("mysql", "root:RsR@0310@tcp(localhost:3306)/go_db")
	if err != nil {
		fmt.Println("Connection failed to open! ")
	} else {
		fmt.Println("Connection Established!")
	}
	defer db.Close()
	db.AutoMigrate(&persons{})
	router := mux.NewRouter()
	router.HandleFunc("/addperson", createNewPerson).Methods("POST")
	http.ListenAndServe(":8000", router)

}
