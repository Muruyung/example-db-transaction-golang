package handler

import "coba/services/coba_svc/domain/usecase"

type Handler struct {
	cobaSatuUC usecase.CobaSatuUseCase
}

func NewHandler(uc *usecase.Wrapper) *Handler {
	return &Handler{
		cobaSatuUC: uc.CobaSatuUC,
	}
}
