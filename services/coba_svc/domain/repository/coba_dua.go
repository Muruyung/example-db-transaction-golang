package repository

import (
	"coba/services/coba_svc/domain/entity"
	"context"

	goutils "github.com/Muruyung/go-utilities"
)

// CobaDuaRepository coba dua repository template
type CobaDuaRepository interface {
	Get(ctx context.Context, query goutils.QueryBuilderInteractor) (*entity.CobaDua, error)
	GetList(ctx context.Context, query goutils.QueryBuilderInteractor) ([]*entity.CobaDua, error)
	GetCount(ctx context.Context, query goutils.QueryBuilderInteractor) (int, error)
	Save(ctx context.Context, data *entity.CobaDua) error
	Update(ctx context.Context, id string, data *entity.CobaDua) error
	Delete(ctx context.Context, id string) error
}
