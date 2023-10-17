package mocks

import "github.com/golang/mock/gomock"

type Wrapper struct {
	*MockCobaDuaUseCase
	*MockCobaSatuUseCase
}

func Init(ctrl *gomock.Controller) Wrapper {
	return Wrapper{
		MockCobaDuaUseCase: NewMockCobaDuaUseCase(ctrl),
		MockCobaSatuUseCase: NewMockCobaSatuUseCase(ctrl),
	}
}
