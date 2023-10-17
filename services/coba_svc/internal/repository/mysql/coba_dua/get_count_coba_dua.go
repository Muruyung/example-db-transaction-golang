package coba_dua_repo

import (
	"context"
	"fmt"

	"coba/pkg/logger"

	goutils "github.com/Muruyung/go-utilities"
	"github.com/rocketlaunchr/dbq/v2"

	"coba/services/coba_svc/internal/repository/models"
)

// GetCount get count data coba dua
func (db *mysqlCobaDuaRepository) GetCount(ctx context.Context, query goutils.QueryBuilderInteractor) (int, error) {
	const commandName = "REPO-GET-COUNT-COBA-DUA"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get count coba dua process...",
		0,
	)

	var (
		err       error
		tableName = models.CobaDuaModels{}.GetTableName()
		count     map[string]int
	)

	query.AddCount("id", "count")
	query.AddWhere("deleted_at", "!=", nil)
	stmt, val, _ := query.GetQuery(tableName, "")
	opts := &dbq.Options{
		SingleResult:   true,
		ConcreteStruct: map[string]int{},
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
		return 0, err
	}

	if result != nil {
		count = result.(map[string]int)
	} else {
		err = fmt.Errorf("coba dua data not found")
		logger.DetailLoggerError(
			ctx,
			commandName,
			"Data not found",
			err,
		)
		return 0, nil
	}

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Get count coba dua success",
		count,
	)
	return count["count"], nil
}
