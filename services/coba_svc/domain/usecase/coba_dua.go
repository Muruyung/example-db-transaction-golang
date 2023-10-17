package usecase

import (
	"coba/services/coba_svc/domain/entity"
	"context"
	"time"

	goutils "github.com/Muruyung/go-utilities"
)

// DTOCobaDua dto for coba dua usecase
type DTOCobaDua struct {
	Name      string
	Age       int
	IsActive  bool
	StartDate time.Time
}

// CobaDuaUseCase coba dua usecase template
type CobaDuaUseCase interface {
	GetCobaDuaByID(ctx context.Context, id string) (*entity.CobaDua, error)
	GetListCobaDua(ctx context.Context, request *goutils.RequestOption) ([]*entity.CobaDua, *goutils.MetaResponse, error)
	CreateCobaDua(ctx context.Context, dto DTOCobaDua) error
	UpdateCobaDua(ctx context.Context, id string, dto DTOCobaDua) error
	DeleteCobaDua(ctx context.Context, id string) error
}
