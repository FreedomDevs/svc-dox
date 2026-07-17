package main

import "github.com/FreedomDevs/svcLibs/go/svcLibs"

var (
	ErrIpNotProvided = svcLibs.ErrorResponseCode{Code: "ERR_IP_NOT_PROVIDED", Message: "IP для ПРОБИВА не указан", Status: 400}
)

var (
	SuccessIpDox = svcLibs.SuccessResponseCode{Code: "SUCCESS_IP_DOX", Message: "IP пробит успешно", Status: 200}
)
