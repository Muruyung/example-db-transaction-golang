package coba_aja_usecase

import (
	"context"
	"fmt"

	"coba/pkg/logger"
)

// DeleteCobaAja update coba aja
func (uc *cobaAjaInteractor) DeleteCobaAja(ctx context.Context, id string) error {
	const commandName = "UC-DELETE-COBA-AJA"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Delete coba aja process...",
		nil,
	)

	err := uc.CobaSatuSvc.DeleteCobaSatu(ctx, id)
	if err != nil {
		logger.DetailLoggerError(
			ctx,
			commandName,
			fmt.Sprintf("Error delete by id=%v", id),
			err,
		)
		return err
	}

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Delete coba aja success",
		nil,
	)
	return nil
}
