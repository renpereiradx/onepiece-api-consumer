package models

type Character struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Jname string `json:"jname"`
	Rname string `json:"rname"`
}

type CharacterFull struct {
	Character      Character       `json:"character"`
	ID             string          `json:"id"`
	Age            string          `json:"age"`
	Birth          string          `json:"birth"`
	Affiliation    string          `json:"affiliation"`
	Ocupation      string          `json:"ocupation"`
	Origin         string          `json:"origin"`
	Height         string          `json:"height"`
	DevilFruit     *DevilFruit     `json:"devil_fruit"`
	CharacterImage *CharacterImage `json:"character_image"`
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
