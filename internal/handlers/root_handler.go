package handlers

import "net/http"

func GetRoot(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
