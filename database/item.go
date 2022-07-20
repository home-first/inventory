package database

type Item struct {
	BaseModel
	Name        string `json:"name"`
	Description string `json:"description"`
}
