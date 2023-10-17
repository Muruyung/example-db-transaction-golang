package repository

import (
	"context"
	"database/sql"

	configmysql "coba/pkg/database/mysql"
)

// MapperCommon template for common mapper models
type MapperCommon interface {
	MapDomainToModels() ModelsCommon
	MapModelsToDomain(entityStruct interface{})
}

// ModelsCommon template for common models repository
type ModelsCommon interface {
	GetTableName() string
	GetModels() interface{}
	GetModelsMap() map[string]interface{}
	GetColumns() []string
	GetValStruct(arrColumn []string) []interface{}
}

// SqlTx template for common transaction sql
type SqlTx interface {
	BeginTx(ctx context.Context, operation func(ctx context.Context, repo *Wrapper) error) error
	Session() *configmysql.DB
	Wrapper() *Wrapper
	DB() *sql.DB
}
