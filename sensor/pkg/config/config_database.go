package config

//"os"

type ConfigDatabase struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func LoadConfigDatabase() *ConfigDatabase {

	/*if erro := godotenv.Load(); erro != nil {
		panic("error load config database")
	}

	var config ConfigDatabase
	config.Host = os.Getenv("DATABASE_HOST")
	config.Port = os.Getenv("DATABASE_PORT")
	config.User = os.Getenv("DATABASE_USER")
	config.Password = os.Getenv("DATABASE_PASSWORD")
	config.Database = os.Getenv("DATABASE_NAME")*/

	var config ConfigDatabase
	config.Host = "localhost"
	config.Port = "5432"
	config.User = "postgres"
	config.Password = "postgres"
	config.Database = "sensor"

	return &config

}
