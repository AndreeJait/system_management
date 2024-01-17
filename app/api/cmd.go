package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"system_management/internal/shared/constant"
)

const (
	FlagEnv = "ENV"
)

var rootCmd = &cobra.Command{
	Use: "application",
	Run: func(cmd *cobra.Command, args []string) {
		envModeString, _ := cmd.Flags().GetString(FlagEnv)
		envMode := constant.GetEnvMode(envModeString)
		startServer(envMode)
	},
}

func Run() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	rootCmd.Flags().String(FlagEnv, string(constant.Development), "to define env")

	err := rootCmd.Execute()
	if err != nil {
		logrus.Error(err)
	}
}
