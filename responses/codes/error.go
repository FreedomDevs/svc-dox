package codes

import "svc-dox/responses"

var (
	ErrIpNotProvided = responses.ErrorResponseCode{Code: "ERR_IP_NOT_PROVIDED", Message: "IP для ПРОБИВА не указан", Status: 400}
	ErrInternalError = func(err error) responses.ErrorResponseCode {
		return responses.ErrorResponseCode{Code: "INTERNAL_ERROR", Message: err.Error(), Status: 500}
	}
)
