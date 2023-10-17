package coba_dua_usecase

import (
	"coba/services/coba_svc/domain/entity"
	"context"
	"fmt"

	"coba/pkg/logger"
)

// GetCobaDuaByID get coba dua by id
func (uc *cobaDuaInteractor) GetCobaDuaByID(ctx context.Context, id string) (*entity.CobaDua, error) {
	const commandName = "UC-GET-COBA-DUA-BY-ID"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get coba dua process...",
		nil,
	)

	res, err := uc.CobaDuaSvc.GetCobaDuaByID(ctx, id)
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
		"Get coba dua success",
		nil,
	)
	return res, nil
}
