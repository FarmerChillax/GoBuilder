package base

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
)

func BuildProject(ctx context.Context, fileName, platform, sourcePath string) {
	logrus.Infof("Build project %s for %s", fileName, platform)
	platformInfo := strings.Split(platform, "/")
	if len(platformInfo) < 2 {
		logrus.Errorf("Invalid platform: %s", platform)
		return
	}
	targetOS, targetArch := platformInfo[0], platformInfo[1]
	outputPath := fmt.Sprintf("./bin/%s/%s/%s", targetOS, targetArch, fileName)
	command := exec.CommandContext(ctx, "go", "build", "-o", outputPath, "-v", "-ldflags", "-s -w", "-trimpath", sourcePath)
	command.Env = append(os.Environ(), "CGO_ENABLED=0", fmt.Sprintf("GOOS=%s", targetOS), fmt.Sprintf("GOARCH=%s", targetArch))
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	logrus.Infof("exec command %s for %s", command.String(), platform)
	if err := command.Run(); err != nil {
		logrus.Errorf("Build project %s failed: %s", fileName, err)
		return
	}
	logrus.Infof("Build project %s for %s success.", fileName, platform)
}
