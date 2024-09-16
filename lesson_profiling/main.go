package main

import (
	"net/http"
	_ "net/http/pprof"
)

func main() {

	http.HandleFunc("/", handleRoot)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}

}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello"))
}
