package coba_satu_usecase

import (
	"context"
	"fmt"

	"coba/pkg/logger"
)

// DeleteCobaSatu update coba satu
func (uc *cobaSatuInteractor) DeleteCobaSatu(ctx context.Context, id string) error {
	const commandName = "UC-DELETE-COBA-SATU"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Delete coba satu process...",
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
		"Delete coba satu success",
		nil,
	)
	return nil
}
