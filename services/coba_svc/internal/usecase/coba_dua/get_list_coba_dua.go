package coba_dua_usecase

import (
	"coba/services/coba_svc/domain/entity"
	"context"

	"coba/pkg/logger"

	goutils "github.com/Muruyung/go-utilities"
)

// GetListCobaDua get list coba dua
func (uc *cobaDuaInteractor) GetListCobaDua(ctx context.Context, request *goutils.RequestOption) ([]*entity.CobaDua, *goutils.MetaResponse, error) {
	const commandName = "UC-GET-LIST-COBA-DUA"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get list coba dua process...",
		nil,
	)

	res, metaRes, err := uc.CobaDuaSvc.GetListCobaDua(ctx, request)
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
		"Get list coba dua success",
		nil,
	)
	return res, metaRes, nil
}
