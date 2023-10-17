package coba_satu_service

import (
	"context"
	"fmt"

	"coba/pkg/logger"
)

// DeleteCobaSatu update coba satu
func (svc *cobaSatuInteractor) DeleteCobaSatu(ctx context.Context, id string) error {
	const commandName = "SVC-DELETE-COBA-SATU"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Delete coba satu process...",
		nil,
	)

	err := svc.repo.CobaSatuRepo.Delete(ctx, id)
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
