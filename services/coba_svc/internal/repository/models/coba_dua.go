package models

import (
	"sort"
	"time"

	"github.com/Muruyung/go-utilities/converter"
)

// CobaDuaModels coba dua models struct
type CobaDuaModels struct {
	ID        string     `dbq:"id" json:"id"`
	Name      string     `dbq:"name" json:"name"`
	Age       int        `dbq:"age" json:"age"`
	IsActive  bool       `dbq:"is_active" json:"is_active"`
	StartDate time.Time  `dbq:"start_date" json:"start_date"`
	CreatedAt time.Time  `dbq:"created_at,omitempty" json:"created_at,omitempty"`
	UpdatedAt time.Time  `dbq:"updated_at,omitempty" json:"updated_at,omitempty"`
	DeletedAt *time.Time `dbq:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}

// GetTableName get table name of coba dua models
func (models CobaDuaModels) GetTableName() string {
	return "coba_dua"
}

// GetModels get models of coba dua models
func (models CobaDuaModels) GetModels() interface{} {
	return models
}

// GetModelsMap get models map of coba dua models
func (models CobaDuaModels) GetModelsMap() map[string]interface{} {
	dataMap, _ := converter.ConvertInterfaceToMap(models)
	return dataMap
}

// GetColumns get columns of coba dua models
func (models CobaDuaModels) GetColumns() []string {
	var (
		modelsMap = models.GetModelsMap()
		arrColumn = make([]string, 0)
	)

	for column := range modelsMap {
		arrColumn = append(arrColumn, column)
	}
	sort.Strings(arrColumn)

	return arrColumn
}

// GetValStruct get value struct of coba dua models
func (models CobaDuaModels) GetValStruct(arrColumn []string) []interface{} {
	var (
		modelsMap = models.GetModelsMap()
		arrValue  = make([]interface{}, 0)
	)

	for _, column := range arrColumn {
		arrValue = append(arrValue, modelsMap[column])
	}

	return []interface{}{arrValue}
}
