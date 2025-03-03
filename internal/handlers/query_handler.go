package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"

	"github.com/flyluman/portfolio-on-golang/internal/database"
	"github.com/flyluman/portfolio-on-golang/internal/models"
)

func PostQuery(w http.ResponseWriter, r *http.Request) {
	var q models.Query
	err := json.NewDecoder(r.Body).Decode(&q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	regexpPattern := regexp.MustCompile(`log|foreign-log|msg`)
	if !regexpPattern.MatchString(q.Name) || q.Pass != os.Getenv("QUERYPASS") {
		http.Error(w, "Unauthorized request to the Server", http.StatusUnauthorized)
		return
	}

	db_rlt, err_q := database.DB.Query(fmt.Sprintf("SELECT * FROM %s ORDER BY id DESC LIMIT 50", q.Name))

	if err_q != nil {
		http.Error(w, err_q.Error(), http.StatusInternalServerError)
		return
	}

	defer db_rlt.Close()

	ret_q := []models.Hit{}

	for db_rlt.Next() {
		tmp := models.Hit{}
		err_s := db_rlt.Scan(&tmp.ID, &tmp.IP, &tmp.ISP, &tmp.City, &tmp.Country, &tmp.Date, &tmp.Path, &tmp.Useragent)

		if err_s != nil {
			http.Error(w, err_s.Error(), http.StatusInternalServerError)
		}

		ret_q = append(ret_q, tmp)
	}

	json.NewEncoder(w).Encode(ret_q)
}
