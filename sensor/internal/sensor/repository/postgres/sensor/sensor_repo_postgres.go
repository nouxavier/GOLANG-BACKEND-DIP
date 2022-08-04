package repository

import (
	"context"
	"database/sql"
	"sensor/internal/sensor/repository/model"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type SensorRepositoryPostgres struct {
	db     *sql.DB
	logger *zap.Logger
}

func NewRepositoryPostgres(db *sql.DB, log *zap.Logger) *SensorRepositoryPostgres {
	return &SensorRepositoryPostgres{db, log}
}

func (s *SensorRepositoryPostgres) Get(ctx context.Context, id uuid.UUID) (*model.Sensor, error) {
	var sensor model.Sensor

	if err := s.db.QueryRowContext(ctx, qSensor_Get, id).
		Scan(&sensor.ID, &sensor.Nome); err != nil {
		return &model.Sensor{}, err
	}

	return &sensor, nil

}

func (s *SensorRepositoryPostgres) Create(ctx context.Context, sensor *model.Sensor) (uuid.UUID, error) {
	err := s.db.QueryRowContext(ctx, qSensor_Create, sensor.Nome, sensor.Nomepais, sensor.Nomeregiao).
		Scan(&sensor.ID)
	if err != nil {
		return uuid.Nil, err
	}

	return sensor.ID, nil

}
