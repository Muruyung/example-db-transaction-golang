package svctx

import (
	"coba/services/coba_svc/domain/repository"
	"coba/services/coba_svc/domain/service"
)

type svcTxInteractor struct {
	repo *repository.Wrapper
}

func NewSvcTx(repo *repository.Wrapper) service.SvcTx {
	return &svcTxInteractor{
		repo: repo,
	}
}
