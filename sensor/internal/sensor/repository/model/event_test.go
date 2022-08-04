package model_test

import (
	"sensor/internal/sensor/repository/model"
	"testing"
	//"sensor/pkg/domain/sensor/v1/modelo"
)

func TestEventoSensor_NewEvento(t *testing.T) {

	type testCase struct {
		test          string
		inputidSensor string
		inputValor    string
		expectedErr   error
	}

	testCases := []testCase{
		{
			test:          "Validação de valor do sensor",
			inputidSensor: "ad4c3174-fd37-42af-ada4-ceb2f457d457",
			inputValor:    "",
			expectedErr:   model.ErrInvalidValorEvento,
		},
		{
			test:          "Validação id sensor",
			inputidSensor: "ad4c3174-fd37-42af-ada4-ceb2f457d457",
			inputValor:    "Valor",
			expectedErr:   model.ErrInvalidIDSensor,
		},
		{
			test:          "Valida objeto",
			inputidSensor: "ad4c3174-fd37-42af-ada4-ceb2f457d457",
			inputValor:    "Valor",
			expectedErr:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := model.NewEvento(tc.inputValor, tc.inputidSensor)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}

		})
	}

}
