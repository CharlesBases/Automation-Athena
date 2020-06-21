package utils

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	DBG = 36

	INF = 38
	WRN = 34
	ERR = 31
)

var level = map[int]string{
	DBG: "DBG",
	INF: "INF",
	WRN: "WRN",
	ERR: "ERR",
}

func ThrowCheck(err error) {
	if err != nil {
		Error(err)
		os.Exit(1)
	}
}

func Debug(vs ...interface{}) {
	print(DBG, vs)
}

func Info(vs ...interface{}) {
	print(INF, vs)
}

func Warn(vs ...interface{}) {
	print(WRN, vs)
}

func Error(vs ...interface{}) {
	print(ERR, vs)
}

func print(code int, vs interface{}) {
	fmt.Print(fmt.Sprintf(
		"%c[%d;%d;%dm[%s][%s] ==> %s%c[0m\n",
		0x1B, 0 /*字体*/, 0 /*背景*/, code /*前景*/, time.Now().Format("2006-01-02 15:04:05.000"), level[code],
		strings.Trim(fmt.Sprint(vs), "[]"),
		0x1B),
	)
}
