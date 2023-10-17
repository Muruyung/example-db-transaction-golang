package service

import (
	"coba/services/coba_svc/domain/repository"
	"coba/services/coba_svc/domain/service"
	coba_dua_service "coba/services/coba_svc/internal/service/coba_dua"
	coba_satu_service "coba/services/coba_svc/internal/service/coba_satu"
)

func Init(repo *repository.Wrapper) *service.Wrapper {
	return &service.Wrapper{
		CobaDuaSvc:  coba_dua_service.NewCobaDuaService(repo),
		CobaSatuSvc: coba_satu_service.NewCobaSatuService(repo),
	}
}
