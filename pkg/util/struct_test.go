package util

import (
	"github.com/lexkong/log"
	"testing"
)

func TestStruct2Map(t *testing.T) {
	c := log.PassLagerCfg{
		Writers:        "file,stdout",
		LoggerLevel:    "DEBUG",
		LoggerFile:     "logs/log.log",
		LogFormatText:  false,
		RollingPolicy:  "size",
		LogRotateDate:  1,
		LogRotateSize:  1,
		LogBackupCount: 7,
	}
	m := Struct2Map(c)
	println(m)
}
