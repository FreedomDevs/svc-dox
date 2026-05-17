package codes

import "svc-dox/response"

var (
	ErrIpNotProvided = response.ErrorResponseCode{Code: "ERR_IP_NOT_PROVIDED", Message: "IP для ПРОБИВА не указан", Status: 400}
	ErrInternalError = func(err error) response.ErrorResponseCode {
		return response.ErrorResponseCode{Code: "INTERNAL_ERROR", Message: err.Error(), Status: 500}
	}
)
