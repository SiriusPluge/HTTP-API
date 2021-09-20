package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"mime"
	"net/http"
	"os"
	"sync"
	"time"
)

type Server struct {
	store *UserStore
}

func NewTaskServer() *Server {
	store := New()
	return &Server{store: store}
}

type User struct {
	Id            string `json:"id"`
	Name          string `json:"name,omitempty"`
	LastName      string `json:"lastName" default:""`
	Age           string `json:"age,omitempty"`
	RecordingDate int64  `json:"recordingDate"`
}

type UserStore struct {
	sync.Mutex
	Users map[string][]User
}

func New() *UserStore {
	us := &UserStore{}
	us.Users = make(map[string][]User)

	return us
}

type ResponseUserRequest struct {
	Error error  `json:"error"`
	Data  []User `json:"data"`
}

//Обработчик запросов
func (s *Server) userHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/api/user" {
		if req.Method == http.MethodPost {
			s.createUserHandler(w, req)
		} else  if req.Method == http.MethodGet {
			id := req.URL.Query().Get("id")
			if len(id) > 0 {
				s.getUserHandler(w, req, id)
			} else {
				s.getAllUsersHandler(w, req)
			}
		} else {
			http.Error(w, fmt.Sprintf("expect method GET or POST at /api/user, got %v", req.Method), http.StatusMethodNotAllowed)
			return
		}
	}
}

//генерация id
func GenerateUUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	uuid := fmt.Sprintf("%x%x%x%x%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}


func (s *Server) createUserHandler(w http.ResponseWriter, req *http.Request) {

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

	id := GenerateUUID()
	RecordingDate := time.Now().Unix()

	var сreateUser User
	сreateUser.Id = id
	сreateUser.Name = jsonData.Name
	сreateUser.LastName = jsonData.LastName
	сreateUser.Age = jsonData.Age
	сreateUser.RecordingDate = RecordingDate

	s.store.Users[сreateUser.Id] = append(s.store.Users[сreateUser.Id], сreateUser)

	js, err := json.Marshal(ResponseUserRequest{Error: err, Data: s.store.Users[id]})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (s *Server) getAllUsersHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("handling task all users at %s\n", req.URL.Path)

	type responceAllUsers struct {
		Error error  `json:"error"`
		Data  *map[string][]User `json:"data"`
	}

	js, err := json.Marshal(responceAllUsers{Error: nil, Data: &s.store.Users})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	fmt.Println(&s.store.Users)
}

func (s *Server) getUserHandler(w http.ResponseWriter, req *http.Request, id string) {
	log.Printf("handling get user at %s\n", req.URL.Path)

	type responceUser struct {
		Error error  `json:"error"`
		Data  []User `json:"data"`
	}

	user := s.store.Users[id]

	js, err := json.Marshal(responceUser{Error: nil, Data: user})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	mux := http.NewServeMux()
	server := NewTaskServer()

	mux.HandleFunc("/api/user", server.userHandler)
	mux.HandleFunc("api/user?id=", server.userHandler)

	log.Fatal(http.ListenAndServe("localhost:"+os.Getenv("SERVERPORT"), mux))
}
