package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/flyluman/portfolio-on-golang/internal/models"
)

func GetWhoami(w http.ResponseWriter, r *http.Request) {

	whoami := models.WhoAmI{
		IP:      r.Header.Get("IP"),
		ISP:     r.Header.Get("ISP"),
		City:    r.Header.Get("City"),
		Country: r.Header.Get("Country"),
	}

	err := json.NewEncoder(w).Encode(whoami)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
