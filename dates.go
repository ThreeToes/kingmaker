package kingmaker

type Date struct {
	Day       int `json:"day"`
	DayOfWeek int `json:"day_of_week"`
	Month     int `json:"month"`
	Year      int `json:"year"`
}

type DateStructure struct {
	Days   []*Day   `json:"days"`
	Months []*Month `json:"months"`
}

type Day struct {
	Name string `json:"name"`
}

type Month struct {
	Name string `json:"name"`
	Days int    `json:"days"`
}
