package main

import (
	"batu/demo"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/tools"
	"net"
	"os"
	//"strconv"
	"time"
)

const (
	HOST = "127.0.0.1"
	PORT = "9099"
)

func main() {
	startTime := currentTimeMillis()

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	transport, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}

	useTransport := transportFactory.GetTransport(transport)
	client := demo.NewBatuThriftClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to "+HOST+":"+PORT, " ", err)
		os.Exit(1)
	}
	defer transport.Close()

	// for i := 0; i < 10; i++ {
	paramMap := make(map[string]string)
	paramMap["path"] = "E:/gows/src/Thrift/server/123456122334.xlsx"
	paramMap["type"] = "json"

	r1, _ := client.CallBack(time.Now().Unix(), 1, paramMap)
	//fmt.Println("value", r1)
	if r1 == "" {
	}

	fmt.Println(r1)

	// model := demo.Article{1, "golang", "Thrift", "testing"}
	// client.Put(&model)
	endTime := currentTimeMillis()
	str := fmt.Sprintf("本次调用用时:%d-%d=%d毫秒\n", endTime, startTime, (endTime - startTime))
	fmt.Println(str)
	tools.LogInfo(str)
	tools.LogInfo("")

}

func currentTimeMillis() int64 {
	return time.Now().UnixNano() / 1000000
}
