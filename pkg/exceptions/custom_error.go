package exceptions

import (
	"net/http"
)

const (
	NOERROR       Status = 0
	ERRDOMAIN     Status = 1
	ERRBUSSINESS  Status = 2
	ERRSYSTEM     Status = 3
	ERRNOTFOUND   Status = 4
	ERRREPOSITORY Status = 5
	ERRUNKNOWN    Status = 6
	ERRAUTHORIZED Status = 7
	ERRFORBIDDEN  Status = 8
)

type Status int

type CustomError struct {
	Status Status
	Error  error
}

func GetCustomErrorCodeByHttpStatusCode(statusCode int) Status {
	var code Status
	switch statusCode {
	case http.StatusBadRequest:
		code = ERRBUSSINESS
	case http.StatusUnauthorized:
		code = ERRAUTHORIZED
	case http.StatusNotFound:
		code = ERRNOTFOUND
	case http.StatusInternalServerError:
		code = ERRSYSTEM
	case http.StatusGatewayTimeout:
		code = ERRREPOSITORY
	default:
		code = ERRSYSTEM
	}

	return code
}
