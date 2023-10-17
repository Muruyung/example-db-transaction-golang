package coba_satu_usecase

import (
	"coba/services/coba_svc/domain/entity"
	"context"
	"fmt"

	"coba/pkg/logger"
)

// GetCobaSatuByID get coba satu by id
func (uc *cobaSatuInteractor) GetCobaSatuByID(ctx context.Context, id string) (*entity.CobaSatu, error) {
	const commandName = "UC-GET-COBA-SATU-BY-ID"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get coba satu process...",
		nil,
	)

	res, err := uc.CobaSatuSvc.GetCobaSatuByID(ctx, id)
	if err != nil {
		logger.DetailLoggerError(
			ctx,
			commandName,
			fmt.Sprintf("Error get by id=%v", id),
			err,
		)
		return nil, err
	}

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get coba satu success",
		nil,
	)
	return res, nil
}
