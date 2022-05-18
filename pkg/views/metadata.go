package views

import (
	"time"

	"github.com/jonathan-innis/go-todo-app/pkg/models"
)

type Metadata struct {
	CreatedAt  time.Time `json:"createdAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
}

func NewMetadataModel(m *Metadata) *models.Metadata {
	return &models.Metadata{
		CreatedAt:  m.CreatedAt,
		ModifiedAt: m.ModifiedAt,
	}
}

func NewMetadataView(m *models.Metadata) *Metadata {
	return &Metadata{
		CreatedAt:  m.CreatedAt,
		ModifiedAt: m.ModifiedAt,
	}
}
