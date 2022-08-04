package responses

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type Responses struct {
	log *zap.Logger
}

func NewResponses(log *zap.Logger) *Responses {
	return &Responses{log}
}

// JSON returns a JSON response to the request
func (r *Responses) JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		if erro := json.NewEncoder(w).Encode(data); erro != nil {
			r.log.Sugar().Fatal(erro)
		}
	}

}

// Error returns an error in JSON format
func (r *Responses) Erro(w http.ResponseWriter, statusCode int, erro error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	}); erro != nil {
		r.log.Sugar().Fatal(erro)
	}

}
