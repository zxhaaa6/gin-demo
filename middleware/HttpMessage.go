package middleware

type HttpMessage struct {
	Success bool
	Status  int
	Message string
	Data    interface{}
}

func GenJson(isSuccess bool, status int, message string, data interface{}) HttpMessage {
	return HttpMessage{Success: isSuccess, Status: status, Message: message, Data: data}
}
