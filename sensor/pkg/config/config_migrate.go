package config

type ConfigMigrate struct {
	Path         string
	MigrateStart bool
}

func LoadConfigMigrate() *ConfigMigrate {

	/*if erro := godotenv.Load(); erro != nil {
		log.Fatal(erro)
		panic("error load configurations log")
		return nil, erro
	}

	var config ConfigLog
	config.Name = os.Getenv("LOG_PROJECT_NAME")
	*/

	var config ConfigMigrate
	config.MigrateStart = true
	config.Path = "file://db/postgres/migrations/sensor"

	return &config

}
