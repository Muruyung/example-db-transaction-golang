package mysqltx

import (
	"context"
	"database/sql"

	"coba/pkg/logger"
	configmysql "coba/pkg/database/mysql"
	"coba/services/coba_svc/domain/repository"
	repository_mysql "coba/services/coba_svc/internal/repository/mysql"
)

// BeginTx begin sql transaction pattern
func (db *mysqlTxRepository) BeginTx(ctx context.Context, operation func(context.Context, *repository.Wrapper) error) error {
	const commandName = "REPO-BEGIN-TRANSACTION"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Begin transaction process...",
		nil,
	)

	rollback := func(tx *sql.Tx, err error) error {
		logger.DetailLoggerWarn(
			ctx,
			commandName,
			"Transaction rollback process...",
			err,
		)

		err = tx.Rollback()
		if err == nil || err == sql.ErrTxDone || err == sql.ErrConnDone {
			logger.DetailLoggerWarn(
				ctx,
				commandName,
				"Transaction rollback success",
				nil,
			)
			return nil
		}

		logger.DetailLoggerError(
			ctx,
			commandName,
			"Transaction rollback failed",
			err,
		)
		return err
	}

	tx, err := db.db.Begin()
	if err != nil {
		logger.DetailLoggerError(
			ctx,
			commandName,
			"Begin transaction failed",
			err,
		)
	}

	dbTx := &mysqlTxRepository{
		db: db.db,
		tx: &configmysql.DB{
			Tx:    tx,
			UseTx: true,
		},
	}

	ctxTx := context.WithValue(ctx, logger.IsUseES, true)
	err = operation(ctxTx, repository_mysql.Init(dbTx))
	if err != nil {
		errRollback := rollback(tx, err)
		if errRollback != nil {
			return errRollback
		}
		return err
	}

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Transaction commit process...",
		nil,
	)
	err = tx.Commit()
	if err != nil {
		err = rollback(tx, err)
		return err
	}

	dbTx.tx.UseTx = false

	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Transaction commit success",
		nil,
	)
	return nil
}
