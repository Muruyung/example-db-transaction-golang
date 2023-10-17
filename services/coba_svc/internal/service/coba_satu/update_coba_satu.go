package coba_satu_service

import (
	"coba/services/coba_svc/domain/entity"
	"coba/services/coba_svc/domain/service"
	"context"

	"coba/pkg/logger"
)

// UpdateCobaSatu update coba satu
func (svc *cobaSatuInteractor) UpdateCobaSatu(ctx context.Context, id string, dto service.DTOCobaSatu) error {
	const commandName = "SVC-UPDATE-COBA-SATU"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Create coba satu process...",
		nil,
	)

	entityDTO, err := entity.NewCobaSatu(entity.DTOCobaSatu{
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

	err = svc.repo.CobaSatuRepo.Update(ctx, id, entityDTO)
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
		"Update coba satu success",
		nil,
	)
	return nil
}
