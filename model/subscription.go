package model

import (
	"polunzh/my-feed/ent"
	"time"
)

type Subscription struct {
	ID        int
	Name      string `json:"name" binding:"required"`
	URL       string `json:"url" binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func ToModel(data *ent.Subscription) *Subscription {
	return &Subscription{
		ID:        data.ID,
		Name:      data.Name,
		URL:       data.URL,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}
