package coba_satu_usecase

import (
	"coba/services/coba_svc/domain/service"
	"coba/services/coba_svc/domain/usecase"
	"context"
	"fmt"

	"coba/pkg/logger"
)

// UpdateCobaSatu update coba satu
func (uc *cobaSatuInteractor) UpdateCobaSatu(ctx context.Context, id string, dto usecase.DTOCobaSatu) error {
	const commandName = "UC-UPDATE-COBA-SATU"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Update coba satu process...",
		nil,
	)

	err := uc.CobaSatuSvc.UpdateCobaSatu(ctx, id, service.DTOCobaSatu{
		Age:       dto.Age,
		IsActive:  dto.IsActive,
		StartDate: dto.StartDate,
		Name:      dto.Name,
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
		"Update coba satu success",
		nil,
	)
	return nil
}
