package utils

import "log"

func HandleError(str string, args ...any) {
	log.Fatalf(str, args...)
}
