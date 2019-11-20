package app

import (
	"github.com/coffee377/voc-admin/internal/app/config"
	"github.com/lexkong/log"
)

func InitLogger() (func(), error) {
	l := config.Global.Log
	log.Initialize(l.Writers, l.LoggerLevel, l.LoggerFile, l.RollingPolicy,
		l.LogFormatText, l.LogRotateDate, l.LogRotateSize, l.LogBackupCount)

	return func() {

	}, nil
}
