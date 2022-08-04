package controllers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sensor/internal/sensor/repository/model"

	"github.com/google/uuid"
)

func (h *Handler) CreateSensor(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.Responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var sensor model.Sensor
	if err = json.Unmarshal(bodyRequest, &sensor); err != nil {
		h.Responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = sensor.Validate(model.Create); err != nil {
		h.Responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	h.SensorService.Create(context.Background(), &sensor)

	h.Responses.JSON(w, http.StatusCreated, sensor)

}

func (h *Handler) GetSensor(w http.ResponseWriter, r *http.Request) {
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

	sensor, err := h.SensorService.Get(context.Background(), id)
	if err != nil {
		h.Responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	h.Responses.JSON(w, http.StatusOK, sensor)
}
