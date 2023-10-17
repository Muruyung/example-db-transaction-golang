package coba_dua_repo

import (
	"context"
	"fmt"

	"coba/pkg/logger"

	goutils "github.com/Muruyung/go-utilities"
	"github.com/rocketlaunchr/dbq/v2"

	"coba/services/coba_svc/domain/entity"
	"coba/services/coba_svc/internal/repository/mapper"
	"coba/services/coba_svc/internal/repository/models"
)

// Get get single data coba dua
func (db *mysqlCobaDuaRepository) Get(ctx context.Context, query goutils.QueryBuilderInteractor) (*entity.CobaDua, error) {
	const commandName = "REPO-GET-COBA-DUA"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get coba dua process...",
		nil,
	)

	var (
		err          error
		tableName    = models.CobaDuaModels{}.GetTableName()
		cobaDua      = &entity.CobaDua{}
		cobaDuaModel *models.CobaDuaModels
	)

	query.AddPagination(goutils.NewPagination(1, 1))
	query.AddWhere("deleted_at", "!=", nil)
	stmt, val, _ := query.GetQuery(tableName, "")
	opts := &dbq.Options{
		SingleResult:   true,
		ConcreteStruct: models.CobaDuaModels{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}

	result, err := dbq.Q(ctx, db.DB(), stmt, opts, val...)
	if err != nil {
		logger.DetailLoggerError(
			ctx,
			commandName,
			"Error execute query",
			err,
		)
		return nil, err
	}

	if result != nil {
		cobaDuaModel = result.(*models.CobaDuaModels)
		cobaDuaMapper := mapper.NewCobaDuaMapper(nil, cobaDuaModel)
		cobaDuaMapper.MapModelsToDomain(cobaDua)
	} else {
		err = fmt.Errorf("coba dua data not found")
		logger.DetailLoggerError(
			ctx,
			commandName,
			"Data not found",
			err,
		)
		return nil, nil
	}

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get coba dua success",
		cobaDuaModel.GetModelsMap(),
	)
	return cobaDua, nil
}
