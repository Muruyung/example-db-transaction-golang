package coba_satu_service

import (
	"coba/services/coba_svc/domain/entity"
	"context"
	"fmt"

	"coba/pkg/logger"

	goutils "github.com/Muruyung/go-utilities"
)

// GetCobaSatuByID get coba satu by id
func (svc *cobaSatuInteractor) GetCobaSatuByID(ctx context.Context, id string) (*entity.CobaSatu, error) {
	const commandName = "SVC-GET-COBA-SATU-BY-ID"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get coba satu process...",
		nil,
	)

	var query = goutils.NewQueryBuilder()
	query.AddWhere("id", "=", "string")
	res, err := svc.repo.CobaSatuRepo.Get(ctx, query)
	if err != nil {
		logger.DetailLoggerError(
			ctx,
			commandName,
			fmt.Sprintf("Error get by id=%v", id),
			err,
		)
		return nil, err
	}

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get coba satu success",
		nil,
	)
	return res, nil
}
