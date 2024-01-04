package build

import (
	"context"
	"sync"

	"github.com/FarmerChillax/GoBuilder/internal/base"
	"github.com/FarmerChillax/GoBuilder/pkg/utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var CmdBuild = &cobra.Command{
	Use:   "build",
	Short: "Build your multi arch application",
	Long:  "You can use `gob build app` to build your multi arch application. Example: gob build app",
	Run:   run,
}

var (
	sourcePath string
	linux      bool
	darwin     bool
	windows    bool
)

func init() {
	CmdBuild.Flags().StringVarP(&sourcePath, "source-path", "o", "./", "source path")
	CmdBuild.Flags().BoolVarP(&linux, "linux", "l", false, "build linux")
	CmdBuild.Flags().BoolVarP(&darwin, "darwin", "d", false, "build darwin")
	CmdBuild.Flags().BoolVarP(&windows, "windows", "w", false, "build windows")
}

func run(_ *cobra.Command, args []string) {
	logrus.SetLevel(logrus.DebugLevel)

	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	ctx := context.Background()

	platfroms := base.GetPlatForm(ctx, linux, darwin, windows)

	fileName := utils.GetFileName()
	if fileName == "" {
		logrus.Fatal("No file name specified")
		return
	}

	println(fileName)

	// 并行编译项目
	wg := &sync.WaitGroup{}
	for _, platform := range platfroms {
		wg.Add(1)
		go func(platform string) {
			defer wg.Done()
			base.BuildProject(ctx, fileName, platform, sourcePath)
		}(platform)
	}
	wg.Wait()
}
