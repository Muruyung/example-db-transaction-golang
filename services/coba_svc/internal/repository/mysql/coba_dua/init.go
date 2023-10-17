package coba_dua_repo

import "coba/services/coba_svc/domain/repository"

type mysqlCobaDuaRepository struct {
	repository.SqlTx
}

// NewCobaDuaRepository initialize new coba dua repository
func NewCobaDuaRepository(db repository.SqlTx) repository.CobaDuaRepository {
	return &mysqlCobaDuaRepository{
		SqlTx: db,
	}
}
