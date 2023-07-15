package config

import (
	"encoding/json"

	"go.uber.org/zap"
)

func InitializeLogger() error {
	configJson := []byte(`{
		"level" : "debug",
		"encoding": "json",   
		"outputPaths":["stderr"],
		"errorOutputPaths":["stderr"],
		"encoderConfig": {
			"messageKey":"message",
			"levelKey":"level",
			"levelEncoder":"lowercase"
		}
	}`)
	var cfg zap.Config
	if err := json.Unmarshal(configJson, &cfg); err != nil {
		return err
	}

	logger, err := cfg.Build()

	// no need to return the logger as we can access our logger with zap.L()
	zap.ReplaceGlobals(logger)

	defer logger.Sync() // flushing the buffer at the end of execution

	if err != nil {
		return err
	}

	return err
}
