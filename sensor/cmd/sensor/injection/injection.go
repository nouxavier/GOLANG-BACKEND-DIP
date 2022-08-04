package injection

import (
	handler "sensor/internal/sensor/port/controllers"
	repoEvent "sensor/internal/sensor/repository/postgres/event"
	repoSensor "sensor/internal/sensor/repository/postgres/sensor"
	"sensor/internal/sensor/service"
	db "sensor/pkg/db"
	"sensor/pkg/responses"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func Injection(datasources *db.DataSources, log *zap.Logger) (*mux.Router, error) {

	/*
	 * Responses
	 */
	rps := responses.NewResponses(log)

	/*
	 * Repository Layer
	 */
	repoSensor := repoSensor.NewRepositoryPostgres(datasources.DB, log)
	repoEvent := repoEvent.NewRepositoryPostgres(datasources.DB, log)

	/*
	 * Service Layer
	 */
	sensorEventService := service.NewEventoSensorService(
		&service.EventSensorConfig{RepositorySensor: repoSensor, RepositoryEvent: repoEvent})

	eventService := service.NewEventService(
		&service.EventConfig{RepositoryEvent: repoEvent})

	sensorService := service.NewSensorService(
		&service.SensorConfig{RepositorySensor: repoSensor})

	/*
	 * Routers
	 */
	router := mux.NewRouter()
	handler.NewHandler(&handler.HandlerConfig{
		Router:             router,
		SensorEventService: sensorEventService,
		EventService:       eventService,
		SensorService:      sensorService,
		MaxBodyBytes:       500,
		TimeoutDuration:    time.Duration(time.Duration(5000) * time.Second),
		Logger:             log,
		Responses:          rps,
	})

	return router, nil

}
