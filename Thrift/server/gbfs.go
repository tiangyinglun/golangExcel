package main

import (
	"batu/demo" //注意导入Thrift生成的接口包
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/tools"
	"os"
)

const (
	NetworkAddr = "127.0.0.1:9099" //监听地址&端口
)

type batuThrift struct {
}

//请求返回
func (this *batuThrift) CallBack(callTime int64, types int32, paramMap map[string]string) (ret string, err error) {
	//请求数据dataRet
	ret, err = tools.HandleData(types, paramMap)
	if err != nil {
		fmt.Println(err)
	}
	tools.LogInfo("返回成功内容")
	return
}

func (this *batuThrift) Put(s *demo.Article) (err error) {
	fmt.Printf("Article--->id: %d\tTitle:%s\tContent:%t\tAuthor:%d\n", s.ID, s.Title, s.Content, s.Author)
	return nil
}

func main() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	var workAddr string
	ip := tools.ReadValue("addr", "ip")
	port := tools.ReadValue("addr", "port")
	tools.LogInfo("服务的ip" + ip)
	tools.LogInfo("服务的ip" + port)
	if ip != "" && port != "" {
		workAddr = ip + ":" + port
	} else {
		workAddr = NetworkAddr
	}

	serverTransport, err := thrift.NewTServerSocket(workAddr)
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}

	handler := &batuThrift{}
	processor := demo.NewBatuThriftProcessor(handler)

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("thrift server in", workAddr)
	server.Serve()
}
