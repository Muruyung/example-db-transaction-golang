package coba_satu_usecase

import (
	"coba/services/coba_svc/domain/entity"
	"context"

	"coba/pkg/logger"

	goutils "github.com/Muruyung/go-utilities"
)

// GetListCobaSatu get list coba satu
func (uc *cobaSatuInteractor) GetListCobaSatu(ctx context.Context, request *goutils.RequestOption) ([]*entity.CobaSatu, *goutils.MetaResponse, error) {
	const commandName = "UC-GET-LIST-COBA-SATU"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get list coba satu process...",
		nil,
	)

	res, metaRes, err := uc.CobaSatuSvc.GetListCobaSatu(ctx, request)
	if err != nil {
		logger.DetailLoggerError(
			ctx,
			commandName,
			"Error get list",
			err,
		)
		return nil, nil, err
	}

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get list coba satu success",
		nil,
	)
	return res, metaRes, nil
}
