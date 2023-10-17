package mocks

import "github.com/golang/mock/gomock"

type Wrapper struct {
	*MockCobaAjaUseCase
	*MockCobaDuaService
	*MockCobaSatuService
	*MockSvcTx
}

func Init(ctrl *gomock.Controller) Wrapper {
	return Wrapper{
		MockCobaAjaUseCase:  NewMockCobaAjaUseCase(ctrl),
		MockCobaDuaService:  NewMockCobaDuaService(ctrl),
		MockCobaSatuService: NewMockCobaSatuService(ctrl),
		MockSvcTx:           NewMockSvcTx(ctrl),
	}
}
