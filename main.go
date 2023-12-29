package main

import (
	"github.com/FarmerChillax/GoBuilder/internal/build"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "gob",
	Short:   "Gobuilder is a multi-architecture rapid build tool with Go",
	Version: release,
}

func init() {
	rootCmd.AddCommand(build.CmdBuild)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Errorln(err)
	}
}
