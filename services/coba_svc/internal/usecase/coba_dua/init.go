package coba_dua_usecase

import (
	"coba/services/coba_svc/domain/service"
	"coba/services/coba_svc/domain/usecase"
)

type cobaDuaInteractor struct {
	*service.Wrapper
}

// NewCobaDuaUseCase initialize new coba dua use case
func NewCobaDuaUseCase(svc *service.Wrapper) usecase.CobaDuaUseCase {
	return &cobaDuaInteractor{
		Wrapper: svc,
	}
}
