package coba_satu_repo

import (
	"context"
	"fmt"
	"time"

	"coba/pkg/logger"

	"github.com/rocketlaunchr/dbq/v2"

	"coba/services/coba_svc/domain/entity"
	"coba/services/coba_svc/internal/repository/mapper"
	"coba/services/coba_svc/internal/repository/models"
)

// Update update data coba satu
func (db *mysqlCobaSatuRepository) Update(ctx context.Context, id string, data *entity.CobaSatu) error {
	const commandName = "REPO-UPDATE-COBA-SATU"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Update coba satu process...",
		nil,
	)

	var (
		err            error
		tableName      = models.CobaSatuModels{}.GetTableName()
		cobaSatuMapper = mapper.NewCobaSatuMapper(data, nil).MapDomainToModels()
		cobaSatuModels = cobaSatuMapper.GetModelsMap()
		arrColumn      = cobaSatuMapper.GetColumns()
		values         = make([]interface{}, 0)
		sqlDB          dbq.ExecContexter
	)
	if db.Session().UseTx {
		sqlDB = db.Session().Tx
	} else {
		sqlDB = db.DB()
	}

	cobaSatuModels["updated_at"] = time.Now()
	arrColumn = append(arrColumn, "updated_at")
	lastIndex := len(arrColumn) - 1

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`UPDATE %s SET`, tableName)
	for key, val := range arrColumn {
		if cobaSatuModels[val] != nil {
			stmt = fmt.Sprintf(`%s %s = ?`, stmt, val)
			values = append(values, cobaSatuModels[val])
		}

		if key < lastIndex && cobaSatuModels[val] != nil {
			stmt += `, `
		} else if key == lastIndex {
			stmt = fmt.Sprintf(`%s WHERE id = ?`, stmt)
		}
	}
	values = append(values, id)

	_, err = dbq.E(ctx, sqlDB, stmt, nil, values...)
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
		"Update coba satu success",
		nil,
	)

	return err
}
