package repositories

import (
	"context"
	"time"

	"github.com/aryanagn/ticket-go-backend/models"
	)

type EventRepository struct {
	db any
}

func (r *EventRepository) GetMany(ctx context.Context) ([]*models.Event, error) {
	events := []*models.Event{}

	events = append(events, &models.Event{
		ID:          "0391084801993113",
		Name:        "Name Test 1",	
		Location:  	 "Location Test 1",
		Date:		 time.Now(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	})

	return events, nil
}
func (r *EventRepository) GetOne(ctx context.Context, eventId string) (*models.Event, error) {
	return nil, nil
}
func (r *EventRepository) CreatedOne(ctx context.Context, event models.Event) (*models.Event, error) {
	return nil, nil
}

func NewEventRepository(db any) models.EventRepository {
	return &EventRepository{
		db: db,
	}
}
