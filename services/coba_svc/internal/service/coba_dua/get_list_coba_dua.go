package coba_dua_service

import (
	"coba/services/coba_svc/domain/entity"
	"context"

	"coba/pkg/logger"

	goutils "github.com/Muruyung/go-utilities"
)

// GetListCobaDua get list coba dua
func (svc *cobaDuaInteractor) GetListCobaDua(ctx context.Context, request *goutils.RequestOption) ([]*entity.CobaDua, *goutils.MetaResponse, error) {
	const commandName = "SVC-GET-LIST-COBA-DUA"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get list coba dua process...",
		nil,
	)

	var (
		query           = goutils.NewQueryBuilder()
		queryPagination = goutils.NewQueryBuilder()
		metaRes         *goutils.MetaResponse
		page            int
		limit           int
	)

	if request != nil {
		query, page, limit = request.SetPaginationWithSort(query)
	}

	res, err := svc.repo.CobaDuaRepo.GetList(ctx, query)
	if err != nil {
		logger.DetailLoggerError(
			ctx,
			commandName,
			"Error get list",
			err,
		)
		return nil, nil, err
	}

	if request != nil && request.GetPagination() != nil {
		totalCount, err := svc.repo.CobaDuaRepo.GetCount(ctx, queryPagination)
		if err != nil {
			logger.DetailLoggerError(
				ctx,
				commandName,
				"Error get total count list",

				err,
			)
			return nil, nil, err
		}

		var meta = goutils.MapMetaResponse(totalCount, len(res), page, limit)
		metaRes = &meta
	}

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get list coba dua success",
		nil,
	)
	return res, metaRes, nil
}
