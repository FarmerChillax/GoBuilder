package base

import (
	"context"
	"os/exec"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

func GetPlatForm(ctx context.Context, linux, darwin, windows bool) []string {
	if !linux && !darwin {
		return allPlatForm(ctx)
	}

	if runtime.GOOS == "windows" {
		return windowsGet(ctx, linux, darwin, windows)
	}

	return get(ctx, linux, darwin, windows)
}

func allPlatForm(ctx context.Context) []string {
	output, err := exec.CommandContext(ctx, "go", "tool", "dist", "list").Output()
	if err != nil {
		logrus.Errorln(err)
	}

	allPlatForm := strings.Split(string(output), "\n")
	return allPlatForm[:len(allPlatForm)-1]
}

func get(ctx context.Context, linux, darwin, windows bool) []string {
	cmdStr := "go tool dist list | grep"
	if linux {
		cmdStr += " -e linux"
	}

	if darwin {
		cmdStr += " -e darwin"
	}

	output, err := exec.CommandContext(ctx, "bash", "-c", cmdStr).Output()
	if err != nil {
		logrus.Error(err)
		return nil
	}

	platForms := strings.Split(string(output), "\n")
	return platForms[:len(platForms)-1]
}

func windowsGet(ctx context.Context, linux, darwin, windows bool) []string {
	cmdStr := `go tool dist list | findstr /i "`
	if linux {
		cmdStr += "linux"
	}

	if darwin {
		cmdStr += " darwin"
	}

	cmdStr += `"`

	cmd := exec.CommandContext(ctx, "powershell.exe", "/C", cmdStr)
	println(cmd.String())
	output, err := cmd.Output()
	if err != nil {
		logrus.Errorln(err)
	}

	platForms := strings.Split(string(output), "\r\n")
	return platForms[:len(platForms)-1]
}
