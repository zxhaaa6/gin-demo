package util

type Error struct {
	Code    int
	Message string
}

func GenUniError(code int, message string) Error {
	return Error{Code: code, Message: message}
}
