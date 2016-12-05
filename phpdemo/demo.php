<?php
/**
 * Thrift RPC - PHPClient
 * @author liuxinming
 * @time 2015.10.8
 **/
namespace batu\testDemo;
header("Content-type: text/html; charset=utf-8");
$ROOT_DIR = realpath(dirname(__FILE__) . '/');
require_once $ROOT_DIR . '/Thrift/ClassLoader/ThriftClassLoader.php';

use Thrift\ClassLoader\ThriftClassLoader;
use Thrift\Protocol\TBinaryProtocol;
use Thrift\Transport\TSocket;
use Thrift\Transport\TSocketPool;
use Thrift\Transport\TFramedTransport;
use Thrift\Transport\TBufferedTransport;

class ThrifClient {
    const ThriftHost = '127.0.0.1'; //UserServer接口服务器IP
    const ThriftPort = 9099;            //UserServer端口

    public function __construct() {

    }

    public function ThriftLoader() {
        $loader = new ThriftClassLoader();
        $GEN_DIR = realpath(dirname(__FILE__) . '/');
        $loader->registerNamespace('Thrift', realpath(dirname(__FILE__) . '/'));
        $loader->registerDefinition('ExcelServer', $GEN_DIR);
        $loader->register();
    }

    public function SocketT($types, $item) {
        $this->ThriftLoader();
        $socket = new TSocket(self::ThriftHost, self::ThriftPort);
        $socket->setSendTimeout(10000);#Sets the send timeout.
        $socket->setRecvTimeout(20000);#Sets the receive timeout.
        //$transport = new TBufferedTransport($socket); #传输方式：这个要和服务器使用的一致 [go提供后端服务,迭代10000次2.6 ~ 3s完成]
        $transport = new TFramedTransport($socket); #传输方式：这个要和服务器使用的一致[go提供后端服务,迭代10000次1.9 ~ 2.1s完成，比TBuffer快了点]
        $protocol = new TBinaryProtocol($transport);  #传输格式：二进制格式
        $client = new \ExcelServer\batuThriftClient($protocol);# 构造客户端
        $transport->open();
        $socket->setDebug(TRUE);
        $result = $client->CallBack(time(), $types, $item); # 对服务器发起rpc调用
        $transport->close();
        return $result;
    }
}

function getMillisecond() {
    list($t1, $t2) = explode(' ', microtime());
    return (float)sprintf('%.0f', (floatval($t1) + floatval($t2)) * 1000);
}

// for($i=1;$i<11;$i++){


//echo "value".implode('',$result)."\n";
// }

// $s = new \batu\demo\Article();
// $s->id = 1;
// $s->title = '插入一篇测试文章';
// $s->content = '我就是这篇文章内容';
// $s->author = 'liuxinming';
// $client->put($s);

// $s->id = 2;
// $s->title = '插入二篇测试文章';
// $s->content = '我就是这篇文章内容';
// $s->author = 'liuxinming';
// $client->put($s);


$startTime = getMillisecond();
$item = array();
$item["path"] = "E:/gows/src/Thrift/server/demo.xlsx";
$types = 1;
$item["type"] = "json";
$ThrifClient = new ThrifClient();
//输出json格式
$res = $ThrifClient->SocketT($types, $item);
echo $res;
$ss = json_decode($res, true);
$excel = $ss['detail'];
$endTime = getMillisecond();
echo "\n使用时间:" . $endTime . "-" . $startTime . " =" . ($endTime - $startTime) . " 毫秒\n";
 
 
