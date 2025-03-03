package handlers

import (
	"net/http"
	"time"

	"github.com/flyluman/portfolio-on-golang/internal/database"
	"github.com/flyluman/portfolio-on-golang/internal/local_time"
)

func PostMessenger(w http.ResponseWriter, r *http.Request) {
	_, err := database.DB.Exec(
		"INSERT INTO msg (ip, isp, city, country, date, useragent, name, email, msg) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		r.Header.Get("IP"), r.Header.Get("ISP"), r.Header.Get("City"), r.Header.Get("Country"),
		time.Now().In(local_time.Location).Format(time.UnixDate), r.UserAgent(),
		r.FormValue("name"), r.FormValue("email"), r.FormValue("msg"),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "https://flyluman.github.io", http.StatusSeeOther)
}
