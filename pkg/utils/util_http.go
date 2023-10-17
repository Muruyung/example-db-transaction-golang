package utils

import (
	"encoding/json"
	"net/http"

	"coba/pkg/exceptions"
)

const (
	CODE_FAILED  = 1
	CODE_SUCCESS = 0
)

type BaseJSONResponse struct {
	Status DetailJSONHeadMessage `json:"status"`
}

type DetailJSONHeadMessage struct {
	Code               int      `json:"code"`
	Message            string   `json:"message"`
	DetailErrorMessage []string `json:"detailError,omitempty"`
}

type SuccessJSONResponse struct {
	BaseJSONResponse
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta,omitempty"`
}

type FailedJSONResponse struct {
	BaseJSONResponse
}

func NewSuccessJSONResponse(data interface{}, message string) SuccessJSONResponse {
	return SuccessJSONResponse{
		BaseJSONResponse: BaseJSONResponse{Status: DetailJSONHeadMessage{
			Code:    CODE_SUCCESS,
			Message: message,
		}},
		Data: data,
	}
}

func NewSuccessJSONResponseWithMeta(meta interface{}, data interface{}, message string) SuccessJSONResponse {
	return SuccessJSONResponse{
		BaseJSONResponse: BaseJSONResponse{Status: DetailJSONHeadMessage{
			Code:    CODE_SUCCESS,
			Message: message,
		}},
		Data: data,
		Meta: meta,
	}
}

func NewFailedJSONResponse(err error, message string) FailedJSONResponse {
	var messageErr string
	if err == nil {
		messageErr = message
	} else {
		messageErr = message + " : " + err.Error()
	}

	return FailedJSONResponse{
		BaseJSONResponse: BaseJSONResponse{Status: DetailJSONHeadMessage{
			Code:    CODE_FAILED,
			Message: messageErr,
		}},
	}
}

func RespondWithError(w http.ResponseWriter, statusCode int, err error) {
	jsonResponse := NewFailedJSONResponse(err, "Failed")
	response, _ := json.Marshal(jsonResponse)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, _ = w.Write(response)
}

func RespondWithCustomError(w http.ResponseWriter, err *exceptions.CustomError) {
	var statusCode int
	switch err.Status {
	case exceptions.ERRDOMAIN:
		statusCode = http.StatusBadRequest
	case exceptions.ERRBUSSINESS:
		statusCode = http.StatusBadRequest
	case exceptions.ERRSYSTEM:
		statusCode = http.StatusInternalServerError
	case exceptions.ERRNOTFOUND:
		statusCode = http.StatusNotFound
	case exceptions.ERRREPOSITORY:
		statusCode = http.StatusBadRequest
	case exceptions.ERRUNKNOWN:
		statusCode = http.StatusBadRequest
	case exceptions.ERRAUTHORIZED:
		statusCode = http.StatusUnauthorized
	case exceptions.ERRFORBIDDEN:
		statusCode = http.StatusForbidden
	}
	jsonResponse := NewFailedJSONResponse(err.Error, "Failed")
	response, _ := json.Marshal(jsonResponse)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, _ = w.Write(response)
}

func RespondWithSuccess(w http.ResponseWriter, statusCode int, data interface{}) {
	successResponse := NewSuccessJSONResponse(data, "Success")
	response, _ := json.Marshal(successResponse)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, _ = w.Write(response)
}

func RespondWithSuccessMeta(w http.ResponseWriter, statusCode int, data interface{}, meta interface{}) {
	successResponse := NewSuccessJSONResponseWithMeta(meta, data, "Success")
	response, _ := json.Marshal(successResponse)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, _ = w.Write(response)
}
