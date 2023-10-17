package coba_aja_usecase

import (
	"coba/services/coba_svc/domain/service"
	"coba/services/coba_svc/domain/usecase/command"
)

type cobaAjaInteractor struct {
	*service.Wrapper
}

// NewCobaAjaUseCase initialize new coba aja use case
func NewCobaAjaUseCase(svc *service.Wrapper) command.CobaAjaUseCase {
	return &cobaAjaInteractor{
		Wrapper: svc,
	}
}
