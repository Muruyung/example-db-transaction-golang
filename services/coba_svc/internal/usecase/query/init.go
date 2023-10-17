package query

import (
	"coba/services/coba_svc/domain/service"
	"coba/services/coba_svc/domain/usecase/query"
	coba_aja_usecase "coba/services/coba_svc/internal/usecase/query/coba_aja"
)

func Init(svc *service.Wrapper) *query.Wrapper {
	return &query.Wrapper{
		CobaAjaUC: coba_aja_usecase.NewCobaAjaUseCase(svc),
	}
}
