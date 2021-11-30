package logger

import "github.com/micro/go-micro/v2/util/log"

// LogError 当存在错误时记录日志
func LogError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}