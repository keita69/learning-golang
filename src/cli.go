package main

import (
	"github.com/Sirupsen/logrus"
	"os"
)

func main() {
	var logger = logrus.New()
	logger.Out = os.Stderr
	logger.Info("aaabbb")
}
