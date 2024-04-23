package domain

type Weather struct {
	Definition  string  `json:"definition"`
	Description string  `json:"description"`
	Temperature float64 `json:"temperature"`
	Humidity    int     `json:"humidity"`
}
