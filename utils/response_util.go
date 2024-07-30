package utils

import (
	"strconv"
	"usrmanagement/models"
	"usrmanagement/responses"
)

func LoginResponse(message, description, token string, user models.ResUser) responses.CreateLoginResponse {
	return responses.CreateLoginResponse{
		ResponseHeader: responses.ResponseHeader{
			StatusCode: strconv.Itoa(1000),
			Message:    "Success",
			StatusDesc: "OK",
		},
		ResponseBody: responses.LoginResBody{
			Message:     message,
			Description: description,
			UserDetails: user,
			Token:       token,
		},
	}
}

func CreateSuccessResponse[T any](message, statusDesc string, data []T) responses.Response[T] {
	header := responses.ResponseHeader{
		Message:    message,
		StatusCode: "200",
		StatusDesc: statusDesc,
	}

	return responses.Response[T]{
		ResponseHeader: header,
		ResponseBody:   data,
	}
}

func CreateErrorResponse(message, statusDesc string) responses.Response[interface{}] {
	header := responses.ResponseHeader{
		Message:    message,
		StatusCode: "500",
		StatusDesc: statusDesc,
	}

	return responses.Response[interface{}]{
		ResponseHeader: header,
		ResponseBody:   nil,
	}
}
