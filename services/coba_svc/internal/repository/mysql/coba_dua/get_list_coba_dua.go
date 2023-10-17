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

// GetList get list data coba dua
func (db *mysqlCobaDuaRepository) GetList(ctx context.Context, query goutils.QueryBuilderInteractor) ([]*entity.CobaDua, error) {
	const commandName = "REPO-GET-LIST-COBA-DUA"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get list coba dua process...",
		nil,
	)

	var (
		err       error
		tableName = models.CobaDuaModels{}.GetTableName()
		data      = make([]interface{}, 0)
	)

	query.AddWhere("deleted_at", "!=", nil)
	stmt, val, _ := query.GetQuery(tableName, "")
	opts := &dbq.Options{
		SingleResult:   false,
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

	cobaDuaModels := result.([]interface{})
	if len(cobaDuaModels) == 0 {
		err = fmt.Errorf("coba dua list data not found")
		logger.DetailLoggerError(
			ctx,
			commandName,
			"Data not found",
			err,
		)
		return nil, nil
	}

	cobaDua := make([]*entity.CobaDua, len(cobaDuaModels))
	for key, val := range cobaDuaModels {
		cobaDuaModel := val.(*models.CobaDuaModels)
		data = append(data, cobaDuaModel.GetModelsMap())
		cobaDua[key] = &entity.CobaDua{}
		cobaDuaMapper := mapper.NewCobaDuaMapper(nil, cobaDuaModel)
		cobaDuaMapper.MapModelsToDomain(cobaDua[key])
	}

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get list coba dua success",
		data,
	)
	return cobaDua, nil
}
