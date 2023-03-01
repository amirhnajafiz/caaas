package main

import (
	"github.com/amirhnajafiz/authX/cmd"
	"github.com/amirhnajafiz/authX/internal/config"
	"github.com/amirhnajafiz/authX/pkg/logger"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func main() {
	root := cobra.Command{}

	// load configs
	cfg := config.LoadConfigs()

	// create logger
	l := logger.NewLogger(cfg.Logger)

	root.AddCommand(
		cmd.HTTP{
			Logger: l.Named("http"),
			Cfg:    cfg,
		}.Command(),
		cmd.Migrate{
			Logger: l.Named("migrate"),
			Cfg:    cfg,
		}.Command(),
	)

	if err := root.Execute(); err != nil {
		l.Error("failed to execute root", zap.Error(err))
	}
}
