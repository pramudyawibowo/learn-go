package controller

import "net/http"

type BaseResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Error   any    `json:"error"`
}

func SuccessResponse(data any) (int, BaseResponse) {
	return http.StatusOK, BaseResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    data,
	}
}

func InternalServerErrorResponse(err error) (int, BaseResponse) {
	return http.StatusInternalServerError, BaseResponse{
		Status:  http.StatusInternalServerError,
		Message: "Internal Server Error",
		Error:   err.Error(),
	}
}

func NotFoundResponse() (int, BaseResponse) {
	return http.StatusNotFound, BaseResponse{
		Status:  http.StatusNotFound,
		Message: "Not Found",
	}
}

func BadRequestResponse(err error) (int, BaseResponse) {
	return http.StatusBadRequest, BaseResponse{
		Status:  http.StatusBadRequest,
		Message: "Bad Request",
		Error:   err.Error(),
	}
}
