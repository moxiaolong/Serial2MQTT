package userlog

import "log"

func Warn(any ...any) {
	log.Printf("[Warn]%+v", any)
}
func Info(any ...any) {
	log.Printf("[Info]%+v", any)
}

func Error(any ...any) {
	log.Printf("[Error]%+v", any)
}
