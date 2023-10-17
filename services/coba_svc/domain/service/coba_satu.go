package service

import (
	"coba/services/coba_svc/domain/entity"
	"context"
	"time"

	goutils "github.com/Muruyung/go-utilities"
)

// DTOCobaSatu dto for coba satu service
type DTOCobaSatu struct {
	Name      string
	Age       int
	IsActive  bool
	StartDate time.Time
}

// CobaSatuService coba satu service template
type CobaSatuService interface {
	GetCobaSatuByID(ctx context.Context, id string) (*entity.CobaSatu, error)
	GetListCobaSatu(ctx context.Context, request *goutils.RequestOption) ([]*entity.CobaSatu, *goutils.MetaResponse, error)
	CreateCobaSatu(ctx context.Context, dto DTOCobaSatu) error
	UpdateCobaSatu(ctx context.Context, id string, dto DTOCobaSatu) error
	DeleteCobaSatu(ctx context.Context, id string) error
}
