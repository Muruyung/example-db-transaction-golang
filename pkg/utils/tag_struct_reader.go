package utils

import (
	"fmt"
	"reflect"

	"coba/pkg/logger"
)

func GetJSONTagNameFromStructField(strct interface{}, field string) (string, error) {
	t := reflect.TypeOf(strct)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	rawTag, ok := t.FieldByName(field)
	if ok {
		return rawTag.Tag.Get("json"), nil
	}

	err := fmt.Errorf("invalid field name of struct %s", t.Name())
	logger.Logger.Error(err)
	return "", err
}
