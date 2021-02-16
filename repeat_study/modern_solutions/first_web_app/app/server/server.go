package server

import (
	"fmt"
	"net/http"
)

// RunWebPortal run web portal on address
func RunWebPortal(address string) error {
	http.HandleFunc("/", rootHandler)
	return http.ListenAndServe(address, nil)
}

func rootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Run web app %s", req.RemoteAddr)
}
