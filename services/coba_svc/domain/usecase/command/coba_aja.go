package command

import (
	"context"
	"time"
)

// DTOCobaAja dto for coba aja usecase
type DTOCobaAja struct {
	StartDate time.Time
	Name      string
	Age       int
	IsActive  bool
}

// CobaAjaUseCase coba aja command usecase template
type CobaAjaUseCase interface {
	CreateCobaAja(ctx context.Context, dto DTOCobaAja) error
	UpdateCobaAja(ctx context.Context, id string, dto DTOCobaAja) error
	DeleteCobaAja(ctx context.Context, id string) error
}
