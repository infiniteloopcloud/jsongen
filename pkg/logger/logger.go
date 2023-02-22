package logger

import "log"

func Info(v ...any) {
	log.Println("[INFO] ", v)
}

func Error(v ...any) {
	log.Println("[ERROR] ", v)
}
