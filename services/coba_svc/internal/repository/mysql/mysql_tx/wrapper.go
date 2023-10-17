package mysqltx

import "coba/services/coba_svc/domain/repository"

// Wrapper get repository wrapper
func (db *mysqlTxRepository) Wrapper() *repository.Wrapper {
	return db.wrapper
}
