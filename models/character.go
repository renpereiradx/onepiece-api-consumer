package models

type Character struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Jname string `json:"jname"`
	Rname string `json:"rname"`
}

type CharacterDetail struct {
	ID          string `json:"id"`
	Age         string `json:"age"`
	Birth       string `json:"birth"`
	Affiliation string `json:"affiliation"`
	Ocupation   string `json:"ocupation"`
	Origin      string `json:"origin"`
	Height      string `json:"height"`
	CharacterID string `json:"character_id"`
	DevilFruit  string `json:"devil_fruit"`
}

type CharacterImage struct {
	ID          string `json:"id"`
	CharacterID string `json:"character_id"`
	ImageURL    string `json:"image_url"`
}

type DevilFruit struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
