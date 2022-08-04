package mocks

import (
	"context"
	"sensor/internal/sensor/repository/model"

	"github.com/stretchr/testify/mock"

	"github.com/google/uuid"
)

type MockEventRepository struct {
	mock.Mock
}

func (m *MockEventRepository) Get(ctx context.Context, id uuid.UUID) (*model.Event, error) {
	ret := m.Called(ctx, id)

	var r0 *model.Event
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*model.Event)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1

}

func (m *MockEventRepository) GetEventsToIDSensor(ctx context.Context, id_sensor uuid.UUID) ([]*model.Event, error) {

	ret := m.Called(ctx, id_sensor)

	var r0 []*model.Event
	if ret.Get(0) != nil {
		r0 = ret.Get(0).([]*model.Event)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1

}

func (m *MockEventRepository) Create(ctx context.Context, evento *model.Event) (uuid.UUID, error) {

	ret := m.Called(ctx, evento)

	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return uuid.New(), r0

}
