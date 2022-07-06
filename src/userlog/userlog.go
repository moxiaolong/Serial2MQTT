package userlog

import (
	"log"
	"os"
)

func Warn(any ...any) {
	log.Printf("[Warn]%+v", any)
}
func Info(any ...any) {
	log.Printf("[Info]%+v", any)
}

func Debug(any ...any) {
	log.Printf("[Debug]%+v", any)
}

func Error(any ...any) {
	log.Printf("[Error]%+v", any)
}
func Fatal(any ...any) {
	log.Printf("[Fatal]%+v", any)
	os.Exit(1)
}
