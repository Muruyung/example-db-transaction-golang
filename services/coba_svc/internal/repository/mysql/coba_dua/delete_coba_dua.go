package coba_dua_repo

import (
	"context"
	"fmt"
	"time"

	"coba/pkg/logger"

	"github.com/Muruyung/go-utilities/converter"
	"github.com/rocketlaunchr/dbq/v2"

	"coba/services/coba_svc/internal/repository/models"
)

// Delete delete data coba dua
func (db *mysqlCobaDuaRepository) Delete(ctx context.Context, id string) error {
	const commandName = "REPO-DELETE-COBA-DUA"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Delete coba dua process...",
		nil,
	)

	var (
		err       error
		tableName = models.CobaDuaModels{}.GetTableName()
		sqlDB     dbq.ExecContexter
	)
	if db.Session().UseTx {
		sqlDB = db.Session().Tx
	} else {
		sqlDB = db.DB()
	}

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`UPDATE %s SET deleted_at = ? WHERE id = ?`, tableName)
	_, err = dbq.E(ctx, sqlDB, stmt, nil, converter.ConvertDateToString(time.Now()), id)
	if err != nil {
		logger.DetailLoggerError(
			ctx,
			commandName,
			"Error execute query",
			err,
		)
		return err
	}

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Delete coba dua success",
		nil,
	)

	return err
}
