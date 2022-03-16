package model

import (
	"log"

	"github.com/go-sohunjug/logger"
)

var DefaultLogger logger.Logger = logger.NewStdLogger(log.Writer())
