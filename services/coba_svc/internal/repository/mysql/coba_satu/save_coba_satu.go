package coba_satu_repo

import (
	"context"
	"time"

	"coba/pkg/logger"

	"github.com/rocketlaunchr/dbq/v2"

	"coba/services/coba_svc/domain/entity"
	"coba/services/coba_svc/internal/repository/mapper"
	"coba/services/coba_svc/internal/repository/models"
)

// Save create data coba satu
func (db *mysqlCobaSatuRepository) Save(ctx context.Context, data *entity.CobaSatu) error {
	const commandName = "REPO-SAVE-COBA-SATU"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Save coba satu process...",
		nil,
	)

	data, _ = data.SetCreatedAt(time.Now())
	data, _ = data.SetUpdatedAt(time.Now())

	var (
		err            error
		tableName      = models.CobaSatuModels{}.GetTableName()
		cobaSatuMapper = mapper.NewCobaSatuMapper(data, nil).MapDomainToModels()
		arrColumn      = cobaSatuMapper.GetColumns()
		arrValue       = cobaSatuMapper.GetValStruct(arrColumn)
		sqlDB          dbq.ExecContexter
	)
	if db.Session().UseTx {
		sqlDB = db.Session().Tx
	} else {
		sqlDB = db.DB()
	}

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := dbq.INSERTStmt(tableName, arrColumn, len(arrValue), dbq.MySQL)
	_, err = dbq.E(ctx, sqlDB, stmt, nil, arrValue)
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
		"Save coba satu success",
		nil,
	)

	return err
}
