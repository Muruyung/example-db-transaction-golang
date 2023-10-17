package coba_aja_usecase

import (
	"context"
	"fmt"

	"coba/pkg/logger"
	"coba/services/coba_svc/domain/service"
	"coba/services/coba_svc/domain/usecase/command"
)

// UpdateCobaAja update coba aja
func (uc *cobaAjaInteractor) UpdateCobaAja(ctx context.Context, id string, dto command.DTOCobaAja) error {
	const commandName = "UC-UPDATE-COBA-AJA"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Update coba aja process...",
		nil,
	)

	err := uc.CobaSatuSvc.UpdateCobaSatu(ctx, id, service.DTOCobaSatu{
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
		"Update coba aja success",
		nil,
	)
	return nil
}
