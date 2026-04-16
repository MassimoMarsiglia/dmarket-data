package config

import (
	"io"
	"os"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

func (c *Config[T]) Load() (T, error) {
	var cfg T

	c.logger.Debug("loading config file", zap.String("path", c.path))

	reader, err := os.Open(c.path)
	if err != nil {
		c.logger.Error("failed to open config", zap.String("path", c.path), zap.Error(err))
		return cfg, err
	}
	defer reader.Close()

	data, err := io.ReadAll(reader)
	if err != nil {
		c.logger.Error("failed to read config content", zap.Error(err))
		return cfg, err
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		c.logger.Error("yaml syntax error", zap.Error(err))
		return cfg, err
	}

	// zap.Any is the standard way to handle a generic T.
	// It will reflect the struct fields into the log entry.
	c.logger.Info("config variables initialized",
		zap.String("path", c.path),
		zap.Any("config_content", cfg),
	)

	return cfg, nil
}
