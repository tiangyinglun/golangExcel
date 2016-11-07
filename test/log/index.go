package main

import (
	"fmt"
	log "github.com/cihub/seelog"
)

func main() {
	defer log.Flush()
	defaultFormat()
	stdFormat()
	dateTimeFormat()
	dateTimeCustomFormat()
	logLevelTypesFormat()
	fileTypesFormat()
	funcFormat()
	xmlFormat()
}

func defaultFormat() {
	fmt.Println("Default format")

	testConfig := `
<seelog type="sync" />`

	logger, err := log.LoggerFromConfigAsBytes([]byte(testConfig))
	if err != nil {
		fmt.Println(err)
	}
	log.ReplaceLogger(logger)

	log.Trace("Test message!")
}

func stdFormat() {
	fmt.Println("Standard fast format")

	testConfig := `
<seelog >
	<outputs formatid="main">
	<rollingfile type="size" filename="./log/manyrolls.log" maxsize="1" maxrolls="4" />
	</outputs>
	<formats>
		<format id="main" format="%Ns [%Level] %Msg%n"/>
	</formats>
</seelog>`

	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
	log.ReplaceLogger(logger)

	log.Trace("Test message!")
}

func dateTimeFormat() {
	fmt.Println("Date time format")

	testConfig := `
<seelog  >
	<outputs formatid="main">
		<rollingfile type="size" filename="./log/manyrolls.log" maxsize="1" maxrolls="4" />
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

	log.Trace("Test message!")
}

func dateTimeCustomFormat() {
	fmt.Println("Date time custom format")

	testConfig := `
<seelog >
	<outputs formatid="main">
		<rollingfile type="size" filename="./log/manyrolls.log" maxsize="1" maxrolls="4" />
	</outputs>
	<formats>
		<format id="main" format="%Date(2006 Jan 02/3:04:05.000000000 PM MST) [%Level] %Msg%n"/>
	</formats>
</seelog>`

	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
	log.ReplaceLogger(logger)

	log.Trace("Test message!")
}

func logLevelTypesFormat() {
	fmt.Println("Log level types format")

	testConfig := `
<seelog type="sync">
	<outputs formatid="main">
		<rollingfile type="size" filename="./log/manyrolls.log" maxsize="1" maxrolls="4" />
	</outputs>
	<formats>
		<format id="main" format="%Level %Lev %LEVEL %LEV %l %Msg%n"/>
	</formats>
</seelog>`

	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
	log.ReplaceLogger(logger)

	log.Trace("Test message!")
}

func fileTypesFormat() {
	fmt.Println("File types format")

	testConfig := `
<seelog type="sync">
	<outputs formatid="main">
		<console/>
	</outputs>
	<formats>
		<format id="main" format="%File %FullPath %RelFile %Msg%n"/>
	</formats>
</seelog>`

	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
	log.ReplaceLogger(logger)

	log.Trace("Test message!")
}

func funcFormat() {
	fmt.Println("Func format")

	testConfig := `
<seelog type="sync">
	<outputs formatid="main">
		<console/>
	</outputs>
	<formats>
		<format id="main" format="%Func %Msg%n"/>
	</formats>
</seelog>`

	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))
	log.ReplaceLogger(logger)

	log.Trace("Test message!")
}

func xmlFormat() {
	fmt.Println("Xml format")

	testConfig := `
<seelog type="sync">
	<outputs formatid="main">
		<console/>
	</outputs>
	<formats>
		<format id="main" format="` +
		`&lt;log&gt;` +
		`&lt;time&gt;%Ns&lt;/time&gt;` +
		`&lt;lev&gt;%l&lt;/lev&gt;` +
		`&lt;msg&gt;%Msg&lt;/msg&gt;` +
		`&lt;/log&gt;"/>
	</formats>
</seelog>`

	logger, _ := log.LoggerFromConfigAsBytes([]byte(testConfig))

	log.ReplaceLogger(logger)

	log.Trace("Test message!")
}
