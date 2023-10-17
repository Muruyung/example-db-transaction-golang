package coba_dua_service

import (
	"context"
	"fmt"

	"coba/pkg/logger"
)

// DeleteCobaDua update coba dua
func (svc *cobaDuaInteractor) DeleteCobaDua(ctx context.Context, id string) error {
	const commandName = "SVC-DELETE-COBA-DUA"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Delete coba dua process...",
		nil,
	)

	err := svc.repo.CobaDuaRepo.Delete(ctx, id)
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
		"Delete coba dua success",
		nil,
	)
	return nil
}
