package coba_aja_usecase

import (
	"context"
	"fmt"

	"coba/pkg/logger"

	"coba/services/coba_svc/domain/entity"
)

// GetCobaAjaByID get coba aja by id
func (uc *cobaAjaInteractor) GetCobaAjaByID(ctx context.Context, id string) (*entity.CobaSatu, error) {
	const commandName = "UC-GET-COBA-AJA-BY-ID"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get coba aja process...",
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
		"Get coba aja success",
		nil,
	)
	return res, nil
}
