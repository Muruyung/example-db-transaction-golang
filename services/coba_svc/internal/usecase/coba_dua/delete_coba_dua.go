package coba_dua_usecase

import (
	"context"
	"fmt"

	"coba/pkg/logger"
)

// DeleteCobaDua update coba dua
func (uc *cobaDuaInteractor) DeleteCobaDua(ctx context.Context, id string) error {
	const commandName = "UC-DELETE-COBA-DUA"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Delete coba dua process...",
		nil,
	)

	err := uc.CobaDuaSvc.DeleteCobaDua(ctx, id)
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
