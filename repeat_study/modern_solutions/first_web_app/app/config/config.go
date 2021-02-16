package config

import "os"

type dbConfig struct {
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// Config collect all different opts
type Config struct {
	DBParams dbConfig
}

//NewConfig create mysql params with all opts
func NewConfig() *Config {
	return &Config{
		DBParams: dbConfig{
			User:     getEnv("USER", ""),
			Password: getEnv("PASSWORD", ""),
			DBName:   getEnv("DB_NAME", ""),
			SSLMode:  getEnv("SSL_MODE", ""),
		},
	}
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
