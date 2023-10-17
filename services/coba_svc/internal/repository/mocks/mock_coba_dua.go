// Code generated by MockGen. DO NOT EDIT.
// Source: /home/robi/ngoding/pribadi/example-db-transaction-golang/services/coba_svc/domain/repository/coba_dua.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entity "coba/services/coba_svc/domain/entity"
	context "context"
	reflect "reflect"

	goutils "github.com/Muruyung/go-utilities"
	gomock "github.com/golang/mock/gomock"
)

// MockCobaDuaRepository is a mock of CobaDuaRepository interface.
type MockCobaDuaRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCobaDuaRepositoryMockRecorder
}

// MockCobaDuaRepositoryMockRecorder is the mock recorder for MockCobaDuaRepository.
type MockCobaDuaRepositoryMockRecorder struct {
	mock *MockCobaDuaRepository
}

// NewMockCobaDuaRepository creates a new mock instance.
func NewMockCobaDuaRepository(ctrl *gomock.Controller) *MockCobaDuaRepository {
	mock := &MockCobaDuaRepository{ctrl: ctrl}
	mock.recorder = &MockCobaDuaRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCobaDuaRepository) EXPECT() *MockCobaDuaRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockCobaDuaRepository) Delete(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCobaDuaRepositoryMockRecorder) Delete(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCobaDuaRepository)(nil).Delete), ctx, id)
}

// Get mocks base method.
func (m *MockCobaDuaRepository) Get(ctx context.Context, query goutils.QueryBuilderInteractor) (*entity.CobaDua, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, query)
	ret0, _ := ret[0].(*entity.CobaDua)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockCobaDuaRepositoryMockRecorder) Get(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockCobaDuaRepository)(nil).Get), ctx, query)
}

// GetCount mocks base method.
func (m *MockCobaDuaRepository) GetCount(ctx context.Context, query goutils.QueryBuilderInteractor) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCount", ctx, query)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCount indicates an expected call of GetCount.
func (mr *MockCobaDuaRepositoryMockRecorder) GetCount(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCount", reflect.TypeOf((*MockCobaDuaRepository)(nil).GetCount), ctx, query)
}

// GetList mocks base method.
func (m *MockCobaDuaRepository) GetList(ctx context.Context, query goutils.QueryBuilderInteractor) ([]*entity.CobaDua, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetList", ctx, query)
	ret0, _ := ret[0].([]*entity.CobaDua)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetList indicates an expected call of GetList.
func (mr *MockCobaDuaRepositoryMockRecorder) GetList(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetList", reflect.TypeOf((*MockCobaDuaRepository)(nil).GetList), ctx, query)
}

// Save mocks base method.
func (m *MockCobaDuaRepository) Save(ctx context.Context, data *entity.CobaDua) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockCobaDuaRepositoryMockRecorder) Save(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockCobaDuaRepository)(nil).Save), ctx, data)
}

// Update mocks base method.
func (m *MockCobaDuaRepository) Update(ctx context.Context, id string, data *entity.CobaDua) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, id, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockCobaDuaRepositoryMockRecorder) Update(ctx, id, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCobaDuaRepository)(nil).Update), ctx, id, data)
}
