package service

import (
	"coba/services/coba_svc/domain/entity"
	"context"
	"time"

	goutils "github.com/Muruyung/go-utilities"
)

// DTOCobaDua dto for coba dua service
type DTOCobaDua struct {
	Age       int
	IsActive  bool
	StartDate time.Time
	Name      string
}

// CobaDuaService coba dua service template
type CobaDuaService interface {
	GetCobaDuaByID(ctx context.Context, id string) (*entity.CobaDua, error)
	GetListCobaDua(ctx context.Context, request *goutils.RequestOption) ([]*entity.CobaDua, *goutils.MetaResponse, error)
	CreateCobaDua(ctx context.Context, dto DTOCobaDua) error
	UpdateCobaDua(ctx context.Context, id string, dto DTOCobaDua) error
	DeleteCobaDua(ctx context.Context, id string) error
}
