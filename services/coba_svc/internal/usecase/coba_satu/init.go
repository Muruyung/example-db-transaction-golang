package coba_satu_usecase

import (
	"coba/services/coba_svc/domain/service"
	"coba/services/coba_svc/domain/usecase"
)

type cobaSatuInteractor struct {
	*service.Wrapper
}

// NewCobaSatuUseCase initialize new coba satu use case
func NewCobaSatuUseCase(svc *service.Wrapper) usecase.CobaSatuUseCase {
	return &cobaSatuInteractor{
		Wrapper: svc,
	}
}
