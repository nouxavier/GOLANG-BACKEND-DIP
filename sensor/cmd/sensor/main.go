package main

import (
	"fmt"
	"net/http"
	i "sensor/cmd/sensor/injection"
	"sensor/pkg/config"
	"sensor/pkg/db"
	"sensor/pkg/log"
)

func main() {
	/*
	 * PKGS
	 */
	log := log.NewLog(config.LoadConfigLog())

	/*
	 * Initialize data sources
	 */
	ds, err := db.InitializeDataSources(log, config.LoadConfigDatabase(), config.LoadConfigMigrate())
	if err != nil {
		log.Sugar().DPanic("Unable to initialize data sources: %v\n", err)
	}

	router, err := i.Injection(ds, log)
	if err != nil {
		log.Sugar().DPanic("Unable to initialize routes: %v\n", err)
	}

	fmt.Println("Running API")
	log.Sugar().DPanic(http.ListenAndServe(":5001", router))

}
