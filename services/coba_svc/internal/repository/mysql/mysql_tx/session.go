package mysqltx

import configmysql "coba/pkg/database/mysql"

// Session get session tx
func (db *mysqlTxRepository) Session() *configmysql.DB {
	if db.tx.UseTx {
		return db.tx
	}

	return &configmysql.DB{
		UseTx: false,
	}
}
