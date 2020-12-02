package reciever

import "net/http"

type ReceiveData struct{}

// RecieverUserDetails
func (*ReceiveData) RecieverUserDetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/octet-stream")
	// ... Do Stuff
}
