package coba_aja_usecase

import (
	"coba/services/coba_svc/domain/service"
	"coba/services/coba_svc/domain/usecase/query"
)

type cobaAjaInteractor struct {
	*service.Wrapper
}

// NewCobaAjaUseCase initialize new coba aja use case
func NewCobaAjaUseCase(svc *service.Wrapper) query.CobaAjaUseCase {
	return &cobaAjaInteractor{
		Wrapper: svc,
	}
}
