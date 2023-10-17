package repository

import (
	"coba/services/coba_svc/domain/entity"
	"context"

	goutils "github.com/Muruyung/go-utilities"
)

// CobaSatuRepository coba satu repository template
type CobaSatuRepository interface {
	Get(ctx context.Context, query goutils.QueryBuilderInteractor) (*entity.CobaSatu, error)
	GetList(ctx context.Context, query goutils.QueryBuilderInteractor) ([]*entity.CobaSatu, error)
	GetCount(ctx context.Context, query goutils.QueryBuilderInteractor) (int, error)
	Save(ctx context.Context, data *entity.CobaSatu) error
	Update(ctx context.Context, id string, data *entity.CobaSatu) error
	Delete(ctx context.Context, id string) error
}
