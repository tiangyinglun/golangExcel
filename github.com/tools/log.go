package tools

import (
	"fmt"
	log "github.com/cihub/seelog"
)

func LogInfo(str string) {
	//日志目录
	var logPath string
	logPath = ReadValue("log", "logpath")
	if logPath == "" {
		if GetOs() == "linux" {
			logPath = "/tmp/log/gbfs.log"
		}

		if GetOs() == "windows" {
			logPath = "./log/gbfs.log"
		}
	}

	testConfig := `
<seelog  >
	<outputs formatid="main">
		<rollingfile type="size" filename="` + logPath + `" maxsize="10000" maxrolls="100000" />
	</outputs>
	<formats>
		<format id="main" format="%Date %Time [%LEV] %Msg%n"/>
	</formats>
</seelog>`

	logger, err := log.LoggerFromConfigAsBytes([]byte(testConfig))

	if err != nil {
		fmt.Println(err)
	}

	loggerErr := log.ReplaceLogger(logger)

	if loggerErr != nil {
		fmt.Println(loggerErr)
	}
	//log.Trace("dafrasfdf")
	log.Info(str)
}
