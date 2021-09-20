package internal

import (
	"HTTP-API-TestTask/pkg"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"time"
)

//Обработчик запросов
func (s *Server) UserHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/api/user" {
		if req.Method == http.MethodPost {
			s.CreateUserHandler(w, req)
		} else if req.Method == http.MethodGet {
			id := req.URL.Query().Get("id")
			if len(id) > 0 {
				s.GetUserHandler(w, req, id)
			} else {
				s.GetAllUsersHandler(w, req)
			}
		} else {
			http.Error(w, fmt.Sprintf("expect method GET or POST at /api/user, got %v", req.Method), http.StatusMethodNotAllowed)
			return
		}
	}
}

func (s *Server) CreateUserHandler(w http.ResponseWriter, req *http.Request) {

	log.Printf("handling task create at %s\n", req.URL.Path)

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

	var jsonData User
	jsonDataFromHttp, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(jsonDataFromHttp), &jsonData)
	if err != nil {
		panic(err)
	}

	id := pkg.GenerateUUID()
	time := time.Now().Unix()
	data := s.CreateUser(id, jsonData.Name, jsonData.LastName, jsonData.Age, time)


	js, err := json.Marshal(ResponseUserRequest{Error: err, Data: data[id]})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (s *Server) GetAllUsersHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling task all users at %s\n", req.URL.Path)

	type responceAllUsers struct {
		Error error              `json:"error"`
		Data  *map[string][]User `json:"data"`
	}

	js, err := json.Marshal(responceAllUsers{Error: nil, Data: &s.Store.Users})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (s *Server) GetUserHandler(w http.ResponseWriter, req *http.Request, id string) {
	log.Printf("handling get user at %s\n", req.URL.Path)

	type responceUser struct {
		Error error  `json:"error"`
		Data  []User `json:"data"`
	}

	user := s.GetUser(id)

	js, err := json.Marshal(responceUser{Error: nil, Data: user})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
