package mysqltx

import (
	"database/sql"

	configmysql "coba/pkg/database/mysql"
	"coba/services/coba_svc/domain/repository"
	repository_mysql "coba/services/coba_svc/internal/repository/mysql"
)

type mysqlTxRepository struct {
	db      *sql.DB
	tx      *configmysql.DB
	wrapper *repository.Wrapper
}

func NewMysqlTx(db *sql.DB) repository.SqlTx {
	tx := &mysqlTxRepository{
		db: db,
	}

	tx.wrapper = repository_mysql.Init(tx)

	return tx
}
