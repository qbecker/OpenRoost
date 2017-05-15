package yoke

type Status struct {
	setTemp     int `json:"setTemp"`
	currentTemp int `json:"currentTemp"`
	homeTemp    int `json:"homeTemp"`
	awayTemp    int `json:"awayTemp"`
}
