package utils

import (
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
