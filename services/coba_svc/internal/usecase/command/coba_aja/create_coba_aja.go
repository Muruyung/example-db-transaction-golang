package coba_aja_usecase

import (
	"context"

	"coba/pkg/logger"
	"coba/services/coba_svc/domain/service"
	"coba/services/coba_svc/domain/usecase/command"
)

// CreateCobaAja create coba aja
func (uc *cobaAjaInteractor) CreateCobaAja(ctx context.Context, dto command.DTOCobaAja) error {
	const commandName = "UC-CREATE-COBA-AJA"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Create coba aja process...",
		nil,
	)

	err := uc.SvcTx.BeginTx(ctx, func(ctx context.Context, svc *service.Wrapper) error {
		err := svc.CobaSatuSvc.CreateCobaSatu(ctx, service.DTOCobaSatu{
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

		err = svc.CobaDuaSvc.CreateCobaDua(ctx, service.DTOCobaDua{
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

		return nil
	})

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Create coba aja success",
		nil,
	)
	return err
}
