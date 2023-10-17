package coba_dua_service

import (
	"coba/services/coba_svc/domain/entity"
	"coba/services/coba_svc/domain/service"
	"context"

	"coba/pkg/logger"
)

// CreateCobaDua create coba dua
func (svc *cobaDuaInteractor) CreateCobaDua(ctx context.Context, dto service.DTOCobaDua) error {
	const commandName = "SVC-CREATE-COBA-DUA"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Create coba dua process...",
		nil,
	)

	entityDTO, err := entity.NewCobaDua(entity.DTOCobaDua{
		Age:       dto.Age,
		IsActive:  dto.IsActive,
		StartDate: dto.StartDate,
		Name:      dto.Name,
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

	err = svc.repo.CobaDuaRepo.Save(ctx, entityDTO)
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
