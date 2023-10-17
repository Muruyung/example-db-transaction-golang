package coba_satu_service

import (
	"context"

	"github.com/docker/distribution/uuid"

	"coba/pkg/converter"
	"coba/pkg/logger"
	"coba/services/coba_svc/domain/entity"
	"coba/services/coba_svc/domain/repository"
	"coba/services/coba_svc/domain/service"
)

// CreateCobaSatu create coba satu
func (svc *cobaSatuInteractor) CreateCobaSatu(ctx context.Context, dto service.DTOCobaSatu) error {
	const commandName = "SVC-CREATE-COBA-SATU"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Create coba satu process...",
		nil,
	)

	entityDTO, err := entity.NewCobaSatu(entity.DTOCobaSatu{
		ID:        uuid.Generate().String(),
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

	err = svc.repo.BeginTx(ctx, func(ctx context.Context, repo *repository.Wrapper) error {
		data, _ := converter.ConvertInterfaceToMap(dto)
		logger.DetailLoggerInfo(
			ctx,
			commandName,
			"Data create coba satu",
			data,
		)

		err = repo.CobaSatuRepo.Save(ctx, entityDTO)
		if err != nil {
			logger.DetailLoggerError(
				ctx,
				commandName,
				"Error create 1",
				err,
			)
			return err
		}

		return nil
	})

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Create coba satu success",
		nil,
	)
	return nil
}
