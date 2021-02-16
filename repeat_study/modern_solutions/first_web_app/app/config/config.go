package config

import "os"

type dbConfig struct {
	MysqlDBName string
}

// Config collect all different opts
type Config struct {
	DBParams dbConfig
}

//NewConfig create mysql params with all opts
func NewConfig() *Config {
	return &Config{
		DBParams: dbConfig{
			MysqlDBName: getEnv("MYSQL_DB_NAME", ""),
		},
	}
}

func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
