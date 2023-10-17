package mocks

import "github.com/golang/mock/gomock"

type Wrapper struct {
	*MockRedisTemplate
}

func Init(ctrl *gomock.Controller) Wrapper {
	return Wrapper{
		MockRedisTemplate: NewMockRedisTemplate(ctrl),
	}
}
