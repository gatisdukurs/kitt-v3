package kitt

import "log"

var logger = log.Default()

func Log(msg string) {
	logger.Println(msg)
}
