package models

type OnePieceApiResponse struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Image          string `json:"image"`
	Jname          string `json:"jName"`
	Rname          string `json:"rname"`
	Affiliation    string `json:"affiliation"`
	Ocupation      string `json:"ocupation"`
	Origin         string `json:"origin"`
	Age            string `json:"age"`
	Birth          string `json:"birth"`
	Height         string `json:"height"`
	DevilFruit     string `json:"dfname"`
	DevilFruitType string `json:"dftype"`
}
