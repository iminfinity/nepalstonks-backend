package api

import (
	"fmt"
	"net/http"
)

// WakeUp func
func WakeUp(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Working")
	fmt.Println("Up and running")
}
