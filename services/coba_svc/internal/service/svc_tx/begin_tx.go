package svctx

import (
	"context"

	"coba/services/coba_svc/domain/repository"
	"coba/services/coba_svc/domain/service"
	svc "coba/services/coba_svc/internal/service"
)

func (svcTx svcTxInteractor) BeginTx(ctx context.Context, operation func(ctx context.Context, svc *service.Wrapper) error) error {
	return svcTx.repo.BeginTx(ctx, func(ctx context.Context, repo *repository.Wrapper) error {
		return operation(ctx, svc.Init(repo, svcTx))
	})
}
