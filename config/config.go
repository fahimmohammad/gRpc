package config

type DbConfig struct {
	DbName    string
	TableName string
}

func NewConfig() *DbConfig {
	return &DbConfig{
		DbName:    "TEST",
		TableName: "article",
	}
}
