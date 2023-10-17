package coba_satu_usecase

import (
	"coba/services/coba_svc/domain/service"
	"coba/services/coba_svc/domain/usecase"
	"context"

	"coba/pkg/logger"
)

// CreateCobaSatu create coba satu
func (uc *cobaSatuInteractor) CreateCobaSatu(ctx context.Context, dto usecase.DTOCobaSatu) error {
	const commandName = "UC-CREATE-COBA-SATU"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Create coba satu process...",
		nil,
	)

	err := uc.CobaSatuSvc.CreateCobaSatu(ctx, service.DTOCobaSatu{
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
		"Create coba satu success",
		nil,
	)
	return nil
}
