package server

import (
	"context"
	"fmt"
	"localhost/lib/reciever"
	"localhost/lib/sender"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

// Server creates a new server and listens to REST requests
type Server struct{ *http.Server }

// NewServer creates a new server
func NewServer(port string) *Server {
	return &Server{&http.Server{
		Addr:         fmt.Sprintf("localhost:%s", port),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}}
}

// ConfigureAndStart will ensure the configurations have been set before starting the server
func (s *Server) ConfigureAndStart(reciever *reciever.ReceiveData, sender *sender.SendData, wg *sync.WaitGroup) {
	router := mux.NewRouter()
	// Create a simple http hander
	router.HandleFunc("/user_data/{userId}/", reciever.RecieverUserDetails).Methods(http.MethodGet)
	router.HandleFunc("/user_data/{userId}/", sender.SendUserDetails).Methods(http.MethodPost)
	http.Handle("/", router)
	wg.Add(1)
	go s.start(wg)
}

// Start will start the server
func (s *Server) start(wg *sync.WaitGroup) {
	// Start the server
	defer wg.Done()
	if err := s.ListenAndServe(); err != nil {
		log.Fatal("Error Server has been stopped: %s", err)
	}
}

// Stop will stop the server
func (s *Server) Stop() {
	if err := s.Shutdown(context.Background()); err != nil {
		log.Fatal("Error stopping server: %s", err)
	}
}
