package mocks

import "github.com/golang/mock/gomock"

type Wrapper struct {
	*MockCobaDuaRepository
	*MockCobaSatuRepository
	*MockSqlTx
}

func Init(ctrl *gomock.Controller) Wrapper {
	return Wrapper{
		MockCobaDuaRepository: NewMockCobaDuaRepository(ctrl),
		MockCobaSatuRepository: NewMockCobaSatuRepository(ctrl),
		MockSqlTx: NewMockSqlTx(ctrl),
	}
}
