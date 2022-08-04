package config

type ConfigLog struct {
	Name string
}

func LoadConfigLog() *ConfigLog {

	/*if erro := godotenv.Load(); erro != nil {
		log.Fatal(erro)
		panic("error load configurations log")
		return nil, erro
	}

	var config ConfigLog
	config.Name = os.Getenv("LOG_PROJECT_NAME")
	*/

	var config ConfigLog
	config.Name = "Teste"

	return &config

}
