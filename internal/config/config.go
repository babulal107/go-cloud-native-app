package config

import "database/sql"

type Config struct {
	Port     string         `json:"port"`
	LogLevel string         `json:"log_level"`
	DataBase DatabaseConfig `json:"DataBase"`
}

type DatabaseConfig struct {
	DriverName string `json:"driverName"`
	Host       string `json:"host"`
	Name       string `json:"name"`
	UserName   string `json:"userName"`
	Password   string `json:"password"`
}

type AppContainer struct {
	DB *sql.DB
}
