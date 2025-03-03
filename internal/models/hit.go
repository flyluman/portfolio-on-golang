package models

type Hit struct {
	ID        int    `json:"id"`
	IP        string `json:"ip"`
	ISP       string `json:"isp"`
	City      string `json:"city"`
	Country   string `json:"country"`
	Date      string `json:"date"`
	Path      string `json:"path"`
	Useragent string `json:"useragent"`
}
