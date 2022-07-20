package database

type Document struct {
	BaseModel
	Name string `json:"name"`
	Type string `json:"type"`
	Path string
}
