package codes

import "svc-dox/responses"

var (
	SuccessLiveOK   = responses.SuccessResponseCode{Code: "LIVE_OK", Message: "svc-dox жив", Status: 200}
	SuccessReadyOK  = responses.SuccessResponseCode{Code: "READY_OK", Message: "Сервис готов к приёму запросов", Status: 200}
	SuccessHealthOK = responses.SuccessResponseCode{Code: "HEALTH_OK", Message: "Сервис работает", Status: 200}
	SuccessIpDox    = responses.SuccessResponseCode{Code: "SUCCESS_IP_DOX", Message: "IP пробит успешно", Status: 200}
)
