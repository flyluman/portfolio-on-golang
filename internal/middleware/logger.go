package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/flyluman/portfolio-on-golang/internal/database"
	"github.com/flyluman/portfolio-on-golang/internal/local_time"
	"github.com/flyluman/portfolio-on-golang/internal/models"
)

func getClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	if xForwardedFor != "" {
		ips := strings.Split(xForwardedFor, ",")
		return strings.TrimSpace(ips[0])
	}

	ip := r.RemoteAddr
	host, _, err := net.SplitHostPort(ip)
	if err != nil {
		return ip
	}
	return host
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// start := time.Now()

		if r.URL.Path != "/query" && r.Method != "POST" {

			resp, err := http.Get(fmt.Sprintf("http://ipwhois.app/json/%s?objects=ip,isp,city,country", getClientIP(r)))
			if err != nil {
				log.Println("Error getting IP info:", err)
				next.ServeHTTP(w, r)
				return
			}
			defer resp.Body.Close()

			var ip_api_resp models.WhoAmI
			json.NewDecoder(resp.Body).Decode(&ip_api_resp)

			r.Header.Set("IP", ip_api_resp.IP)
			r.Header.Set("ISP", ip_api_resp.ISP)
			r.Header.Set("City", ip_api_resp.City)
			r.Header.Set("Country", ip_api_resp.Country)

			db_table := "log"

			if ip_api_resp.Country != "Bangladesh" {
				db_table = "foreign_log"
			}

			db_q := fmt.Sprintf("INSERT INTO %s (ip, isp, city, country, date, path, useragent) VALUES (?, ?, ?, ?, ?, ?, ?)", db_table)

			_, err = database.DB.Exec(db_q,
				r.Header.Get("IP"), r.Header.Get("ISP"), r.Header.Get("City"), r.Header.Get("Country"),
				time.Now().In(local_time.Location).Format(time.UnixDate), r.URL.Path, r.UserAgent())

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		next.ServeHTTP(w, r)
		// log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
	})
}
