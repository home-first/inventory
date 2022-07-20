package config

type Config struct {
	Port           string   `json:"port"`
	DatabaseFile   string   `json:"databaseFile"`
	Secret         string   `json:"secret"`
	TrustedProxies []string `json:"trustedProxies"`
	DocumentPath   string   `json:"documentPath"`
	LabelRoot      string   `json:"labelRoot"`
}

var Cfg Config

func Init() {
	Cfg = Config{
		Port:           "8125",
		DatabaseFile:   "./bin/database.db",
		Secret:         "secret",
		TrustedProxies: []string{"127.0.0.1"},
		DocumentPath:   "./bin/documents",
		LabelRoot:      "l",
	}
}
