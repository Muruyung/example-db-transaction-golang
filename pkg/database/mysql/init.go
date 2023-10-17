package configmysql

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"time"

	"coba/pkg/logger"

	_ "github.com/go-sql-driver/mysql"

	"coba/pkg/dotenv"
)

type DB struct {
	*sql.Tx
	UseTx bool
}

// InitMysqlDB initialize mysql db
func InitMysqlDB(connection string) *sql.DB {
	logger.Logger.Info("initialize database mysql...")
	ctx := context.WithValue(context.Background(), logger.ActivityID, "INIT-MYSQL-DB")
	var (
		err    error
		dbConn *sql.DB
	)

	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", dotenv.APPTIMEZONE())
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	if dotenv.IsAppEnvProduction() {
		dbConn, err = sql.Open(`nrmysql`, dsn)
	} else {
		dbConn, err = sql.Open(`mysql`, dsn)
	}

	if err != nil {
		logger.DetailLoggerError(
			ctx,
			"INIT-MYSQL-DB",
			"error initialize database mysql",
			err,
		)
		return nil
	}

	dbConn.SetMaxOpenConns(300)
	dbConn.SetMaxIdleConns(25)
	dbConn.SetConnMaxLifetime(5 * time.Minute)

	return dbConn
}

// InitMysqlTx initialize sql transaction pattern
func InitMysqlTx(db *sql.DB) *DB {
	tx, _ := db.Begin()
	return &DB{
		Tx:    tx,
		UseTx: false,
	}
}
