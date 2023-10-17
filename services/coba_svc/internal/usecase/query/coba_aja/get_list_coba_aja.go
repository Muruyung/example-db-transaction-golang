package coba_aja_usecase

import (
	"context"

	"coba/pkg/logger"

	goutils "github.com/Muruyung/go-utilities"

	"coba/services/coba_svc/domain/entity"
)

// GetListCobaAja get list coba aja
func (uc *cobaAjaInteractor) GetListCobaAja(ctx context.Context, request *goutils.RequestOption) ([]*entity.CobaSatu, *goutils.MetaResponse, error) {
	const commandName = "UC-GET-LIST-COBA-AJA"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get list coba aja process...",
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
		"Get list coba aja success",
		nil,
	)
	return res, metaRes, nil
}
