package command

import (
	"coba/services/coba_svc/domain/service"
	"coba/services/coba_svc/domain/usecase/command"
	coba_aja_usecase "coba/services/coba_svc/internal/usecase/command/coba_aja"
)

func Init(svc *service.Wrapper) *command.Wrapper {
	return &command.Wrapper{
		CobaAjaUC: coba_aja_usecase.NewCobaAjaUseCase(svc),
	}
}
