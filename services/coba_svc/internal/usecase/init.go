package usecase

import (
	"coba/services/coba_svc/domain/service"
	"coba/services/coba_svc/domain/usecase"
	coba_dua_usecase "coba/services/coba_svc/internal/usecase/coba_dua"
	coba_satu_usecase "coba/services/coba_svc/internal/usecase/coba_satu"
)

func Init(svc *service.Wrapper) *usecase.Wrapper {
	return &usecase.Wrapper{
		CobaDuaUC:  coba_dua_usecase.NewCobaDuaUseCase(svc),
		CobaSatuUC: coba_satu_usecase.NewCobaSatuUseCase(svc),
	}
}
