package data

type Pokemon struct {
	Rank      int64  `json:"rank"`
	Name      string `json:"name"`
	Type_1    string `json:"type_1"`
	Type_2    string `json:"type_2,omitempty"`
	Legendary bool   `json:"legendary"`
}
