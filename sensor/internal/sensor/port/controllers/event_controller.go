package controllers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sensor/internal/sensor/repository/model"

	"github.com/google/uuid"
)

func (h *Handler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.Responses.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var event model.Event
	if err = json.Unmarshal(bodyRequest, &event); err != nil {
		h.Responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = event.Validate(model.Create); err != nil {
		h.Responses.Erro(w, http.StatusBadRequest, err)
		return
	}

	h.EventService.Create(context.Background(), &event)

	h.Responses.JSON(w, http.StatusCreated, event)

}

func (h *Handler) GetEvent(w http.ResponseWriter, r *http.Request) {
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

	evento, err := h.EventService.Get(context.Background(), id)
	if err != nil {
		h.Responses.Erro(w, http.StatusInternalServerError, err)
		return
	}
	h.Responses.JSON(w, http.StatusOK, evento)
}
