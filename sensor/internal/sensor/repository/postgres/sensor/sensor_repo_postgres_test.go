package repository_test

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"sensor/internal/sensor/repository/model"
	repo "sensor/internal/sensor/repository/postgres/sensor"
	"sensor/pkg/config"
	dbpostgres "sensor/pkg/db/postgres"
	"sensor/pkg/log"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

func Test_sensorRepo_Create(t *testing.T) {

	type testCase struct {
		name       string
		ctx        context.Context
		input      *model.Sensor
		beforeTest func(sqlmock.Sqlmock)
		want       uuid.UUID
		wantErr    bool
	}

	log := log.NewLog(config.LoadConfigLog())
	configDatabase := config.LoadConfigDatabase()
	postgres := dbpostgres.NewPostgres(configDatabase, log)

	sensor_input := &model.Sensor{Nome: "Sensor Nouara 2 ", Nomeregiao: model.EnumCentroOeste, Nomepais: model.EnumBrasil}
	tests := []testCase{
		{
			name:  "Sucesso ao criar sensor",
			ctx:   context.TODO(),
			input: sensor_input,
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectQuery(regexp.QuoteMeta(`
						INSERT INTO sensores (nome, nome_regiao, nome_pais)
						VALUES ($1, $2, $3);`,
					)).
					WithArgs("Sensor Nouara 2", model.EnumCentroOeste, model.EnumBrasil).
					WillReturnError(errors.New("whoops, error"))
			},
			want:    sensor_input.ID,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			db, _ := postgres.PostgresConnection()
			u := repo.NewRepositoryPostgres(db, log)

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			got, err := u.Create(tt.ctx, tt.input)
			if err != nil {
				t.Errorf("sensorRepo.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == uuid.Nil {
				t.Errorf("sensorRepo.Create() = %v, want %v", got, tt.want)
			}

		})
	}

}

func Test_sensorRepo_Get(t *testing.T) {

	type testCase struct {
		name       string
		ctx        context.Context
		input      uuid.UUID
		beforeTest func(sqlmock.Sqlmock)
		want       model.Sensor
		wantErr    bool
	}
	id_test, _ := uuid.Parse("3a0944dd-2498-4fd3-93b2-30cc224956c2")
	log := log.NewLog(config.LoadConfigLog())
	configDatabase := config.LoadConfigDatabase()
	postgres := dbpostgres.NewPostgres(configDatabase, log)

	tests := []testCase{
		{
			name:  "Sucesso ao pegar sensor",
			ctx:   context.TODO(),
			input: id_test,
			beforeTest: func(mockSQL sqlmock.Sqlmock) {
				mockSQL.
					ExpectQuery(regexp.QuoteMeta(`
						select id, nome from sensores where id=$1`,
					)).
					WithArgs(id_test).
					WillReturnError(errors.New("whoops, error"))
			},
			want:    model.Sensor{ID: id_test, Nome: "Sensor Nouara"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			db, _ := postgres.PostgresConnection()
			u := repo.NewRepositoryPostgres(db, log)

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			got, err := u.Get(tt.ctx, tt.input)
			if err != nil {
				t.Errorf("sensorRepo.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Nome != tt.want.Nome {
				t.Errorf("sensorRepo.Create() = %v, want %v", got.Nome, tt.want.Nome)
			}
			if got.ID != tt.want.ID {
				t.Errorf("sensorRepo.Create() = %v, want %v", got.ID, tt.want.ID)
			}

		})
	}

}
