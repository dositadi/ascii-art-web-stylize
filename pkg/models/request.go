package models

type Ascii struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
	Text   string `json:"text"`
	Banner string `json:"banner"`
}

type RequestResponse struct {
	Original    string
	Transformed string
	BannerFont  string
	CreatedAt   string
}
