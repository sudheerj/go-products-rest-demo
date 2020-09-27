package configs

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Database string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Host:     "127.0.0.1",
			Port:     3306,
			Username: "root",
			Password: "root@2020",
			Name: "product_reviews",
			Charset:  "utf8",
		},
	}
}
