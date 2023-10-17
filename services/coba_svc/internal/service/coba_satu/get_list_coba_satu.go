package coba_satu_service

import (
	"coba/services/coba_svc/domain/entity"
	"context"

	"coba/pkg/logger"

	goutils "github.com/Muruyung/go-utilities"
)

// GetListCobaSatu get list coba satu
func (svc *cobaSatuInteractor) GetListCobaSatu(ctx context.Context, request *goutils.RequestOption) ([]*entity.CobaSatu, *goutils.MetaResponse, error) {
	const commandName = "SVC-GET-LIST-COBA-SATU"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get list coba satu process...",
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

	res, err := svc.repo.CobaSatuRepo.GetList(ctx, query)
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
		totalCount, err := svc.repo.CobaSatuRepo.GetCount(ctx, queryPagination)
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
		"Get list coba satu success",
		nil,
	)
	return res, metaRes, nil
}
