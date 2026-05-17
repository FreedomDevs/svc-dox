package codes

import "svc-dox/response"

var (
	SuccessLiveOK   = response.SuccessResponseCode{Code: "LIVE_OK", Message: "svc-dox жив", Status: 200}
	SuccessReadyOK  = response.SuccessResponseCode{Code: "READY_OK", Message: "Сервис готов к приёму запросов", Status: 200}
	SuccessHealthOK = response.SuccessResponseCode{Code: "HEALTH_OK", Message: "Сервис работает", Status: 200}
	SuccessIpDox    = response.SuccessResponseCode{Code: "SUCCESS_IP_DOX", Message: "IP пробит успешно", Status: 200}
)
