package models

type WhoAmI struct {
	IP      string `json:"ip"`
	ISP     string `json:"isp"`
	City    string `json:"city"`
	Country string `json:"country"`
}
