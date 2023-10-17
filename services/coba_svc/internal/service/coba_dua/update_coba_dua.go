package coba_dua_service

import (
	"coba/services/coba_svc/domain/entity"
	"coba/services/coba_svc/domain/service"
	"context"

	"coba/pkg/logger"
)

// UpdateCobaDua update coba dua
func (svc *cobaDuaInteractor) UpdateCobaDua(ctx context.Context, id string, dto service.DTOCobaDua) error {
	const commandName = "SVC-UPDATE-COBA-DUA"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Create coba dua process...",
		nil,
	)

	entityDTO, err := entity.NewCobaDua(entity.DTOCobaDua{
		Name:      dto.Name,
		Age:       dto.Age,
		IsActive:  dto.IsActive,
		StartDate: dto.StartDate,
	})
	if err != nil {
		logger.DetailLoggerError(
			ctx,
			commandName,
			"Error generate entity",
			err,
		)
		return err
	}

	err = svc.repo.CobaDuaRepo.Update(ctx, id, entityDTO)
	if err != nil {
		logger.DetailLoggerError(
			ctx,
			commandName,
			"Error update",
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
