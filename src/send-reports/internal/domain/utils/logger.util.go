package utils

import (
	"log"
	"runtime"
)

var (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Green = "\033[32m"
)

func ErrorLog(funcName string, err error) {
	_, file, line, _ := runtime.Caller(1)
	log.Printf(Red+"%s:%d type=\"error\" name=\"%s\" data=\"%v\""+Reset, file, line, funcName, err)
}

func Info(funcName string, message interface{}) {
	_, file, line, _ := runtime.Caller(1)
	log.Printf(Green+"%s:%d type=\"info\" name=\"%s\" data=\"%v\""+Reset, file, line, funcName, message)
}
