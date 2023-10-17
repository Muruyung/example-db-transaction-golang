package handler

import "coba/services/coba_svc/domain/usecase/command"

type Handler struct {
	cobaAjaCommand command.CobaAjaUseCase
}

func NewHandler(command *command.Wrapper) *Handler {
	return &Handler{
		cobaAjaCommand: command.CobaAjaUC,
	}
}
