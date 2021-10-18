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

	sqlStat := `INSERT INTO person (name, last_name, phone) VALUES ($1, $2, $3) returning id`
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

	rows, err := db.Query("select * from person")
	if err != nil {
		panic(err)
	}
	defer rows.Close()


	type ManyUsers []Person

	var manyUsers ManyUsers
	for rows.Next() {
		var user Person
		err = rows.Scan(&user.Id, &user.Name, &user.LastName, &user.Phone)
		if err != nil {
			log.Fatal("Error in Scan rows!")
		}
		manyUsers = append(manyUsers, user)
	}


	//users := []Person
	//
	//for rows.Next(){
	//	u := users[]
	//	err := rows.Scan(&u.id, &u.Name, &u.LastName, &u.Phone)
	//	if err != nil{
	//		fmt.Println(err)
	//		continue
	//	}
	//	users = append(users, u)
	//}
	//
	peopleBytes, _ := json.MarshalIndent(manyUsers, "", "\t")

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

	rows, err := db.Query("select * from person where id = $1", id)
	if err != nil {
		log.Fatal(err)
	}

	type ManyUsers []Person

	var manyUsers ManyUsers
	for rows.Next() {
		var user Person
		err = rows.Scan(&user.Id, &user.Name, &user.LastName, &user.Phone)
		if err != nil {
			log.Fatal("Error in Scan rows!")
		}
		manyUsers = append(manyUsers, user)
	}

	//var people []Person
	//
	//var person Person
	//rows.Scan(&person.Name, &person.LastName, &person.Phone)
	//people = append(people, person)


	peopleBytes, _ := json.MarshalIndent(manyUsers, "", "\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(peopleBytes)

	defer rows.Close()
	defer db.Close()
}
