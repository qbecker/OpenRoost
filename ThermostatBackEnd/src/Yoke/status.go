package yoke

type Status struct {
	SetTemp     int `json:"setTemp"`
	CurrentTemp int `json:"currentTemp"`
	HomeTemp    int `json:"homeTemp"`
	AwayTemp    int `json:"awayTemp"`
}
