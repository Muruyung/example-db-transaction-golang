package coba_satu_repo

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

// GetList get list data coba satu
func (db *mysqlCobaSatuRepository) GetList(ctx context.Context, query goutils.QueryBuilderInteractor) ([]*entity.CobaSatu, error) {
	const commandName = "REPO-GET-LIST-COBA-SATU"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get list coba satu process...",
		nil,
	)

	var (
		err       error
		tableName = models.CobaSatuModels{}.GetTableName()
		data      = make([]interface{}, 0)
	)

	query.AddWhere("deleted_at", "!=", nil)
	stmt, val, _ := query.GetQuery(tableName, "")
	opts := &dbq.Options{
		SingleResult:   false,
		ConcreteStruct: models.CobaSatuModels{},
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

	cobaSatuModels := result.([]interface{})
	if len(cobaSatuModels) == 0 {
		err = fmt.Errorf("coba satu list data not found")
		logger.DetailLoggerError(
			ctx,
			commandName,
			"Data not found",
			err,
		)
		return nil, nil
	}

	cobaSatu := make([]*entity.CobaSatu, len(cobaSatuModels))
	for key, val := range cobaSatuModels {
		cobaSatuModel := val.(*models.CobaSatuModels)
		data = append(data, cobaSatuModel.GetModelsMap())
		cobaSatu[key] = &entity.CobaSatu{}
		cobaSatuMapper := mapper.NewCobaSatuMapper(nil, cobaSatuModel)
		cobaSatuMapper.MapModelsToDomain(cobaSatu[key])
	}

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get list coba satu success",
		data,
	)
	return cobaSatu, nil
}
