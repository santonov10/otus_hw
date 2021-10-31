package logger

import (
	"fmt"
	"testing"

	"github.com/santonov10/otus_hw/hw12_13_14_15_calendar/internal/config"
)

func TestLogger(t *testing.T) {
	config.SetFilePath("../../configs/default.json")
	logger := Get()
	logger.Error().Err(fmt.Errorf("error")).Msg("this is the way to log errors")
}
