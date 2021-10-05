package init

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/AntiD2ta/validators-alert-mechanism/internal"
)

func Init() (config internal.Config, err error) {
	rawJson := os.Getenv("VALIDATORS")
	if rawJson == "" {
		return
	}
	err = json.Unmarshal([]byte(rawJson), &config)
	if err != nil {
		return
	}

	rawInterval := os.Getenv("INTERVAL")
	interval, err := strconv.ParseInt(rawInterval, 10, 64)
	if err != nil {
		return
	}
	config.Interval = int(interval)
	return
}
