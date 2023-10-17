package coba_satu_repo

import "coba/services/coba_svc/domain/repository"

type mysqlCobaSatuRepository struct {
	repository.SqlTx
}

// NewCobaSatuRepository initialize new coba satu repository
func NewCobaSatuRepository(db repository.SqlTx) repository.CobaSatuRepository {
	return &mysqlCobaSatuRepository{
		SqlTx: db,
	}
}
