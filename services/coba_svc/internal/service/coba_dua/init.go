package coba_dua_service

import (
	"coba/services/coba_svc/domain/repository"
	"coba/services/coba_svc/domain/service"
)

type cobaDuaInteractor struct {
	repo *repository.Wrapper
}

// NewCobaDuaService initialize new coba dua service
func NewCobaDuaService(repo *repository.Wrapper) service.CobaDuaService {
	return &cobaDuaInteractor{
		repo: repo,
	}
}
