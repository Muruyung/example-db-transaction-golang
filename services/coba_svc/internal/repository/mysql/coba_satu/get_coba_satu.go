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

// Get get single data coba satu
func (db *mysqlCobaSatuRepository) Get(ctx context.Context, query goutils.QueryBuilderInteractor) (*entity.CobaSatu, error) {
	const commandName = "REPO-GET-COBA-SATU"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get coba satu process...",
		nil,
	)

	var (
		err           error
		tableName     = models.CobaSatuModels{}.GetTableName()
		cobaSatu      = &entity.CobaSatu{}
		cobaSatuModel *models.CobaSatuModels
	)

	query.AddPagination(goutils.NewPagination(1, 1))
	query.AddWhere("deleted_at", "!=", nil)
	stmt, val, _ := query.GetQuery(tableName, "")
	opts := &dbq.Options{
		SingleResult:   true,
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

	if result != nil {
		cobaSatuModel = result.(*models.CobaSatuModels)
		cobaSatuMapper := mapper.NewCobaSatuMapper(nil, cobaSatuModel)
		cobaSatuMapper.MapModelsToDomain(cobaSatu)
	} else {
		err = fmt.Errorf("coba satu data not found")
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
		"Get coba satu success",
		cobaSatuModel.GetModelsMap(),
	)
	return cobaSatu, nil
}
