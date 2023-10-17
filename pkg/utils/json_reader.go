package utils

import (
	"fmt"
	"io"
	"os"

	"coba/pkg/logger"
)

func GetDataJSONFromStub(NameFile string) []byte {
	jsonFile, err := os.Open(NameFile)
	// if we os.Open returns an error then handle it
	if err != nil {
		logger.Logger.Error(err)
		fmt.Println(err)
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)

	return byteValue
}
