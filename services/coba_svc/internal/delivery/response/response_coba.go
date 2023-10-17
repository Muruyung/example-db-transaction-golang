package response

import (
	"github.com/Muruyung/go-utilities/converter"

	"coba/services/coba_svc/domain/entity"
)

type ResponseCoba struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	IsActive  bool   `json:"is_active"`
	StartDate string `json:"start_date"`
}

func MapCobaResponse(coba *entity.CobaSatu) ResponseCoba {
	return ResponseCoba{
		Name:      coba.GetName(),
		Age:       coba.GetAge(),
		IsActive:  coba.GetIsActive(),
		StartDate: coba.GetCreatedAt().Format(converter.DateLayoutFull),
	}
}