package middleware

import "github.com/zxhaaa6/gin-demo/util"

type SimpleMessage struct {
	Success bool
	Status  int
	Message string
	Data    interface{}
}

type PageMessage struct {
}

func GenSimpleJson(data interface{}) SimpleMessage {
	return SimpleMessage{Success: true, Status: 200, Message: "success", Data: data}
}

func GenErrorJson(err util.Error) SimpleMessage {
	return SimpleMessage{Success: false, Status: err.Code, Message: err.Message}
}
