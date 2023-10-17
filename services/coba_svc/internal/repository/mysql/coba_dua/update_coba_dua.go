package coba_dua_repo

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

// Update update data coba dua
func (db *mysqlCobaDuaRepository) Update(ctx context.Context, id string, data *entity.CobaDua) error {
	const commandName = "REPO-UPDATE-COBA-DUA"
	logger.DetailLoggerInfo(
		ctx,
		commandName,
		"Update coba dua process...",
		nil,
	)

	var (
		err           error
		tableName     = models.CobaDuaModels{}.GetTableName()
		cobaDuaMapper = mapper.NewCobaDuaMapper(data, nil).MapDomainToModels()
		cobaDuaModels = cobaDuaMapper.GetModelsMap()
		arrColumn     = cobaDuaMapper.GetColumns()
		values        = make([]interface{}, 0)
		sqlDB         dbq.ExecContexter
		lastIndex     = len(arrColumn) - 1
	)
	if db.Session().UseTx {
		sqlDB = db.Session().Tx
	} else {
		sqlDB = db.DB()
	}

	cobaDuaModels["updated_at"] = time.Now()
	arrColumn = append(arrColumn, "updated_at")

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`UPDATE %s SET`, tableName)
	for key, val := range arrColumn {
		if cobaDuaModels[val] != nil {
			stmt = fmt.Sprintf(`%s %s = ?`, stmt, val)
			values = append(values, cobaDuaModels[val])
		}

		if key < lastIndex && cobaDuaModels[val] != nil {
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
		"Update coba dua success",
		nil,
	)

	return err
}
