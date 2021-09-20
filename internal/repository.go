package internal

import (
	"sync"
)

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

func (s *Server) CreateUser(id string, name string, lastName string, age string, RecordingDate int64) map[string][]User {
	s.Store.Lock()
	defer s.Store.Unlock()

	//user := User{
	//	Id:            id,
	//	Name:          name,
	//	LastName:      lastName,
	//	Age:           Age,
	//	RecordingDate: RecordingDate,
	//}

	//s.Store.Users[user.Id] = append(s.Store.Users[user.Id], user)


	//id := pkg.GenerateUUID()
	//RecordingDate := time.Now().Unix()

	var сreateUser User
	сreateUser.Id = id
	сreateUser.Name = name
	сreateUser.LastName = lastName
	сreateUser.Age = age
	сreateUser.RecordingDate = RecordingDate

	s.Store.Users[сreateUser.Id] = append(s.Store.Users[сreateUser.Id], сreateUser)

	return s.Store.Users
}

func (s *Server) GetUser(id string) []User {
	s.Store.Lock()
	defer s.Store.Unlock()

	t, ok := s.Store.Users[id]
	if ok {
		return t
	}
	return t
}



