package coba_satu_service

import (
	"coba/services/coba_svc/domain/repository"
	"coba/services/coba_svc/domain/service"
)

type cobaSatuInteractor struct {
	repo *repository.Wrapper
}

// NewCobaSatuService initialize new coba satu service
func NewCobaSatuService(repo *repository.Wrapper) service.CobaSatuService {
	return &cobaSatuInteractor{
		repo: repo,
	}
}
