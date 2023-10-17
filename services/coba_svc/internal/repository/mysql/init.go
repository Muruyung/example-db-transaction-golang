package repository_mysql

import (
	"coba/services/coba_svc/domain/repository"
	coba_dua_repo "coba/services/coba_svc/internal/repository/mysql/coba_dua"
	coba_satu_repo "coba/services/coba_svc/internal/repository/mysql/coba_satu"
)

func Init(db repository.SqlTx) *repository.Wrapper {
	return &repository.Wrapper{
		CobaDuaRepo:  coba_dua_repo.NewCobaDuaRepository(db),
		CobaSatuRepo: coba_satu_repo.NewCobaSatuRepository(db),
		SqlTx:        db,
	}
}
