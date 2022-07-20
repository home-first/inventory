package database

type Collection struct {
	BaseModel
	Name  string `json:"name"`
	Items []Item `json:"items" gorm:"foreignkey:ID"`
}
