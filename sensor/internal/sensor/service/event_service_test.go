package service

import (
	"context"
	"fmt"
	"sensor/internal/sensor/repository/model"
	"testing"

	"sensor/internal/sensor/mocks"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGet_eventService(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		uid, _ := uuid.NewRandom()
		uid_sensor, _ := uuid.NewRandom()

		mockEventResp := &model.Event{
			ID:       uid,
			IDSensor: uid_sensor,
			Valor:    "20A",
		}

		mockEventRepository := new(mocks.MockEventRepository)
		s := NewEventService(&EventConfig{
			RepositoryEvent: mockEventRepository,
		})
		mockEventRepository.On("Get", mock.Anything, uid).Return(mockEventResp, nil)

		ctx := context.TODO()
		u, err := s.Get(ctx, uid)

		assert.NoError(t, err)
		assert.Equal(t, u, mockEventResp)
		mockEventRepository.AssertExpectations(t)
	})
	t.Run("Error", func(t *testing.T) {
		uid, _ := uuid.NewRandom()

		mockEventRepository := new(mocks.MockEventRepository)
		s := NewEventService(&EventConfig{
			RepositoryEvent: mockEventRepository,
		})
		mockEventRepository.On("Get", mock.Anything, uid).Return(nil, fmt.Errorf("Some error down the call chain"))

		ctx := context.TODO()
		u, err := s.Get(ctx, uid)

		assert.Equal(t, u, &model.Event{})
		assert.Error(t, err)
		mockEventRepository.AssertExpectations(t)
	})
}

func TestCreate_eventService(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		uid_sensor, _ := uuid.NewRandom()

		mockEvent := &model.Event{
			IDSensor: uid_sensor,
			Valor:    "20A",
		}

		mockEventRepository := new(mocks.MockEventRepository)
		s := NewEventService(&EventConfig{
			RepositoryEvent: mockEventRepository,
		})

		// We can use Run method to modify the user when the Create method is called.
		//  We can then chain on a Return method to return no error
		mockEventRepository.
			On("Create", mock.AnythingOfType("*context.emptyCtx"), mockEvent).
			Run(func(args mock.Arguments) {
				eventArg := args.Get(1).(*model.Event)
				eventArg.IDSensor = uid_sensor
			}).Return(nil)

		ctx := context.TODO()
		id_event, err := s.Create(ctx, mockEvent)

		assert.NoError(t, err)

		assert.NotNil(t, id_event)
		assert.Equal(t, uid_sensor, mockEvent.IDSensor)
		mockEventRepository.AssertExpectations(t)
	})

}
