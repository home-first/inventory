package database

type User struct {
	BaseModel
	Username     string `json:"username"`
	Name         string `json:"name"`
	PasswordHash []byte
}
