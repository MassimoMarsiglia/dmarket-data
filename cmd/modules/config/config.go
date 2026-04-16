package config

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func New[T any](cmd *cobra.Command, logger *zap.Logger) (*Config[T], error) {
	path, err := cmd.Flags().GetString("cfg")
	if err != nil {
		return nil, err
	}
	return &Config[T]{logger: logger, path: path}, nil
}

func InitFlags(cmd *cobra.Command) {
	cmd.Flags().String("cfg", "", "set config path")
}
