package coba_dua_usecase

import (
	"coba/services/coba_svc/domain/service"
	"coba/services/coba_svc/domain/usecase"
	"context"

	"coba/pkg/logger"
)

// CreateCobaDua create coba dua
func (uc *cobaDuaInteractor) CreateCobaDua(ctx context.Context, dto usecase.DTOCobaDua) error {
	const commandName = "UC-CREATE-COBA-DUA"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Create coba dua process...",
		nil,
	)

	err := uc.CobaDuaSvc.CreateCobaDua(ctx, service.DTOCobaDua{
		Name:      dto.Name,
		Age:       dto.Age,
		IsActive:  dto.IsActive,
		StartDate: dto.StartDate,
	})
	if err != nil {
		logger.DetailLoggerError(
			ctx,
			commandName,
			"Error create",
			err,
		)
		return err
	}

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Create coba dua success",
		nil,
	)
	return nil
}
