package controllers

import (
	"context"
	"io/ioutil"
	"net/http"
	"sensor/internal/sensor/repository/model"

	"sensor/pkg/config/sensor/responses"

	"github.com/google/uuid"
)

type eventsSensorVO struct {
	eventos []*model.Event
	sensor  *model.Sensor
}

func (h *Handler) GetSensorEvents(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	id, err := uuid.Parse(string(bodyRequest))
	if err != nil {
		responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	sensor, events, err := h.SensorEventService.GetSensorEvents(context.Background(), id)
	if err != nil {
		responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, eventsSensorVO{eventos: events, sensor: sensor})
}
