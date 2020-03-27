package server

import (
	"net/http"
	"realtimechat/structs"
	"sync"

	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
)

// Subscriber -
type Subscriber func(msg string) error

// Server -
type Server struct {
	router      *chi.Mux
	upgrader    *websocket.Upgrader
	conf        *structs.Conf
	submutex    *sync.Mutex
	subscribers map[string]Subscriber
}

// New -
func New() *Server {
	router := chi.NewRouter()

	conf := config()

	upgrader := &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	serv := &Server{
		router:      router,
		upgrader:    upgrader,
		conf:        conf,
		submutex:    &sync.Mutex{},
		subscribers: map[string]Subscriber{},
	}

	serv.ApplyHandlers()

	return serv
}

// Start -
func (serv *Server) Start() error {
	return http.ListenAndServe(":"+serv.conf.Port, serv.router)
}
