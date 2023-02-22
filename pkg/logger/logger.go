package logger

import (
	"log"
	"os"
)

func Info(v ...any) {
	log.Println("[INFO] ", v)
}

func Error(v ...any) {
	log.Println("[ERROR] ", v)
	os.Exit(1)
}
