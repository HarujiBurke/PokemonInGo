package entities

type Pokemon struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Weight      uint   `json:"weight"`
	Description string `json:"description"`
}
