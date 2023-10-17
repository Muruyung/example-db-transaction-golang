package query

import (
	"context"

	goutils "github.com/Muruyung/go-utilities"

	"coba/services/coba_svc/domain/entity"
)

// CobaAjaUseCase coba aja query usecase template
type CobaAjaUseCase interface {
	GetCobaAjaByID(ctx context.Context, id string) (*entity.CobaSatu, error)
	GetListCobaAja(ctx context.Context, request *goutils.RequestOption) ([]*entity.CobaSatu, *goutils.MetaResponse, error)
}
