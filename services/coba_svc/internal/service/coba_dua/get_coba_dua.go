package coba_dua_service

import (
	"coba/services/coba_svc/domain/entity"
	"context"
	"fmt"

	"coba/pkg/logger"

	goutils "github.com/Muruyung/go-utilities"
)

// GetCobaDuaByID get coba dua by id
func (svc *cobaDuaInteractor) GetCobaDuaByID(ctx context.Context, id string) (*entity.CobaDua, error) {
	const commandName = "SVC-GET-COBA-DUA-BY-ID"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get coba dua process...",
		nil,
	)

	var query = goutils.NewQueryBuilder()
	query.AddWhere("id", "=", "string")
	res, err := svc.repo.CobaDuaRepo.Get(ctx, query)
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
		"Get coba dua success",
		nil,
	)
	return res, nil
}
