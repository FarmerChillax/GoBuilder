package utils

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// 获取程序执行时间
func GetRunTime(startTime time.Time) (runTime string) {
	endTime := time.Now()
	runTime = endTime.Sub(startTime).String()
	logrus.Infof("Run time: %s", runTime)
	return runTime
}

// GetFileName
func GetFileName() (fileName string) {
	if len(os.Args) > 1 {
		return os.Args[len(os.Args)-1]
	}
	return ""
}
