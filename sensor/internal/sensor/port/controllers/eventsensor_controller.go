package controllers

import (
	"context"
	"io/ioutil"
	"net/http"
	"sensor/internal/sensor/repository/model"


	"github.com/google/uuid"
)

type eventsSensorVO struct {
	eventos []*model.Event
	sensor  *model.Sensor
}

func (h *Handler) GetSensorEvents(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.Responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	id, err := uuid.Parse(string(bodyRequest))
	if err != nil {
		h.Responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	sensor, events, err := h.SensorEventService.GetSensorEvents(context.Background(), id)
	if err != nil {
		h.Responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	h.Responses.JSON(w, http.StatusOK, eventsSensorVO{eventos: events, sensor: sensor})
}
