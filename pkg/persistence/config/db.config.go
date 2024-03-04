package config

import "os"

type DataBaseConfig struct {
	Host string
	Port string
	User string
	Pass string
	Name string
}

func GetDataBaseConfig() DataBaseConfig {
	return DataBaseConfig{
		Host: os.Getenv("DATABASE_HOST"),
		Port: os.Getenv("DATABASE_PORT"),
		User: os.Getenv("DATABASE_USER"),
		Pass: os.Getenv("DATABASE_PASS"),
		Name: os.Getenv("DATABASE_NAME"),
	}
}

func (dbc DataBaseConfig) GetConnectionString() string {
	return "postgres://" + dbc.User + ":" + dbc.Pass + "@" + dbc.Host + ":" + dbc.Port + "/" + dbc.Name
}
