package internal

type Server struct {
	Store *UserStore
}

func NewUserServer() *Server {
	Store := New()
	return &Server{Store: Store}
}