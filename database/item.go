package database

type Item struct {
	BaseModel
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Collections []*Collection `json:"collections" gorm:"many2many:collection_items;"`
}

func (item *Item) Minimize() ItemMinimal {
	return ItemMinimal{
		BaseModel:   item.BaseModel,
		Name:        item.Name,
		Description: item.Description,
	}
}

type ItemMinimal struct {
	BaseModel
	Name        string `json:"name"`
	Description string `json:"description"`
}
