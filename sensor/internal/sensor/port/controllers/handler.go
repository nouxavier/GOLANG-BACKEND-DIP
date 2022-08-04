package controllers

import (
	"net/http"
	"sensor/internal/sensor/repository/model"
	"time"

	"sensor/pkg/responses"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

type Handler struct {
	SensorEventService model.SensorEventService
	EventService       model.EventService
	SensorService      model.SensorService
	MaxBodyBytes       int64
	Logger             *zap.Logger
	Responses          *responses.Responses
}

type HandlerConfig struct {
	Router             *mux.Router
	SensorEventService model.SensorEventService
	EventService       model.EventService
	SensorService      model.SensorService
	MaxBodyBytes       int64
	TimeoutDuration    time.Duration
	Logger             *zap.Logger
	Responses          *responses.Responses
}

type Route struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	RequiredAuthentication bool
}

func NewHandler(config *HandlerConfig) {
	h := &Handler{
		SensorEventService: config.SensorEventService,
		EventService:       config.EventService,
		SensorService:      config.SensorService,
		MaxBodyBytes:       config.MaxBodyBytes,
		Logger:             config.Logger,
		Responses:          config.Responses,
	}

	var routeSensor = []Route{
		{
			URI:                    "/sensor",
			Method:                 http.MethodPost,
			Function:               h.CreateSensor,
			RequiredAuthentication: false,
		},
		{
			URI:                    "/sensors/{sensorId}",
			Method:                 http.MethodGet,
			Function:               h.GetSensor,
			RequiredAuthentication: false,
		},
	}

	var routeEvent = []Route{
		{
			URI:                    "/event",
			Method:                 http.MethodPost,
			Function:               h.CreateEvent,
			RequiredAuthentication: false,
		},
		{
			URI:                    "/event/{eventId}",
			Method:                 http.MethodGet,
			Function:               h.GetEvent,
			RequiredAuthentication: false,
		},
	}

	var routeEventSensor = []Route{
		{
			URI:                    "/sensorevent/{sensorId}",
			Method:                 http.MethodGet,
			Function:               h.GetSensorEvents,
			RequiredAuthentication: false,
		},
	}

	routes := routeEventSensor
	routes = append(routes, routeEvent...)
	routes = append(routes, routeSensor...)

	for _, route := range routes {
		config.Router.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

}
