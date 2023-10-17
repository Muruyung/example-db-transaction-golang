package coba_dua_usecase

import (
	"coba/services/coba_svc/domain/service"
	"coba/services/coba_svc/domain/usecase"
	"context"
	"fmt"

	"coba/pkg/logger"
)

// UpdateCobaDua update coba dua
func (uc *cobaDuaInteractor) UpdateCobaDua(ctx context.Context, id string, dto usecase.DTOCobaDua) error {
	const commandName = "UC-UPDATE-COBA-DUA"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Update coba dua process...",
		nil,
	)

	err := uc.CobaDuaSvc.UpdateCobaDua(ctx, id, service.DTOCobaDua{
		Name:      dto.Name,
		Age:       dto.Age,
		IsActive:  dto.IsActive,
		StartDate: dto.StartDate,
	})
	if err != nil {
		logger.DetailLoggerError(
			ctx,
			commandName,
			fmt.Sprintf("Error update by id=%v", id),
			err,
		)
		return err
	}

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Update coba dua success",
		nil,
	)
	return nil
}
