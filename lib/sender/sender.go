package sender

import (
	"net/http"
	"sync"
)

type SendData struct {
	waitGroup *sync.WaitGroup
	stopChan  chan bool
}



// sendUserDetails
func (*SendData) SendUserDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/octet-stream")
	// ... Do Stuff
}

func (s *SendData) Stop() {
	s.stopChan <- true
}
