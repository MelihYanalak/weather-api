package domain

type Weather struct {
	Temperature int    `json:"temperature"`
	Condition   string `json:"condition"`
	Valid       bool   `json:"valid"`
}
