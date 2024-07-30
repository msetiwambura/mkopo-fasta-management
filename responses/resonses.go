package responses

import (
	"usrmanagement/models"
)

type ResponseBody struct {
	Message     string `json:"Message"`
	Description string `json:"Description"`
	ItemId      uint   `json:"ItemId"`
}

type LoginResBody struct {
	Message     string `json:"Message"`
	Description string `json:"Description"`
	UserDetails models.ResUser
	Token       string `json:"Token"`
}

type CreateLoginResponse struct {
	ResponseHeader ResponseHeader `json:"ResponseHeader"`
	ResponseBody   LoginResBody   `json:"ResponseBody"`
}

type CreateResponse struct {
	ResponseHeader ResponseHeader `json:"ResponseHeader"`
	ResponseBody   ResponseBody   `json:"ResponseBody"`
}

type RolesResponseBody struct {
	Id   uint   `json:"Id"`
	Name string `json:"Name"`
}

type CreateRolesResponse struct {
	ResponseHeader ResponseHeader    `json:"ResponseHeader"`
	ResponseBody   RolesResponseBody `json:"ResponseBody"`
}

type ResponseHeader struct {
	Message    string `json:"Message"`
	StatusCode string `json:"StatusCode"`
	StatusDesc string `json:"StatusDesc"`
}

type Response[T any] struct {
	ResponseHeader ResponseHeader `json:"ResponseHeader"`
	ResponseBody   []T            `json:"ResponseBody"`
}
