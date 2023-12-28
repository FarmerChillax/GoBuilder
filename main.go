package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/FarmerChillax/GoBuilder/config"
	"github.com/sirupsen/logrus"
)

func GetFileName() (fileName string) {
	if len(os.Args) > 1 {
		return os.Args[len(os.Args)-1]
	}
	return ""
}

func BuildProject(ctx context.Context, fileName, platform, sourcePath string) {
	logrus.Infof("Build project %s for %s", fileName, platform)
	platformInfo := strings.Split(platform, "/")
	if len(platformInfo) < 2 {
		logrus.Errorf("Invalid platform: %s", platform)
		return
	}
	targetOS, targetArch := platformInfo[0], platformInfo[1]
	outputPaht := fmt.Sprintf("./bin/%s/%s/%s", targetOS, targetArch, fileName)
	command := exec.CommandContext(ctx, "go", "build", "-o", outputPaht, "-v", "-ldflags", "-s -w", "-trimpath", sourcePath)
	command.Env = append(os.Environ(), fmt.Sprintf("CGO_ENABLED=0"), fmt.Sprintf("GOOS=%s", targetOS), fmt.Sprintf("GOARCH=%s", targetArch))
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	logrus.Infof("exec command %s for %s", command.String(), platform)
	if err := command.Run(); err != nil {
		logrus.Errorf("Build project %s failed: %s", fileName, err)
		return
	}
	logrus.Infof("Build project %s for %s success.", fileName, platform)
}

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	//log.SetLevel(log.InfoLevel)

	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	// 读取配置文件
	cfg, err := config.LoadConfig()
	if err != nil {
		logrus.Fatal(err)
		return
	}

	fileName := GetFileName()
	if fileName == "" {
		logrus.Fatal("No file name specified")
		return
	}

	// 并行编译项目
	ctx := context.Background()
	wg := &sync.WaitGroup{}
	for _, platform := range cfg.Platform {
		wg.Add(1)
		go func(platform string) {
			defer wg.Done()
			BuildProject(ctx, fileName, platform, cfg.SourcePath)
		}(platform)
	}
	wg.Wait()
}
