package internal

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"mime"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, req *http.Request) {

	log.Printf("handling task create at %s\n", req.URL.Path)

	db := OpenConnection()

	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	var jsonData Person
	errSecond := json.NewDecoder(req.Body).Decode(&jsonData)
	if errSecond != nil{
		http.Error(w, errSecond.Error(), http.StatusBadRequest)
	}

	sqlStat := `INSERT INTO person (name, last_name, phone) VALUES ($1, $2, $3)`
	_, err = db.Exec(sqlStat, jsonData.Name, jsonData.LastName, jsonData.Phone)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	defer db.Close()
}

func GetAllUsersHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling task all users at %s\n", req.URL.Path)

	db := OpenConnection()

	rows, err := db.Query("SELECT * FROM person")
	if err != nil {
		log.Fatal(err)
	}

	var people []Person

	for rows.Next() {
		var person Person
		rows.Scan(&person.Name, &person.LastName, &person.Phone)
		people = append(people, person)
	}

	peopleBytes, _ := json.MarshalIndent(people, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(peopleBytes)

	defer rows.Close()
	defer db.Close()
}

func GetUserHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling get user at %s\n", req.URL.Path)

	db := OpenConnection()

	vars := mux.Vars(req)
	id := vars["id"]

	rows, err := db.Query("SELECT id FROM person")
	if err != nil {
		log.Fatal(err)
	}

	var people []Person

	var person Person
	rows.Scan(&person.Name, &person.LastName, &person.Phone)
	people = append(people, person)


	peopleBytes, _ := json.MarshalIndent(people, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(peopleBytes)

	defer rows.Close()
	defer db.Close()
}
