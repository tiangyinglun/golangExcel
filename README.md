# golangExcel

1.php调用golang解析xlsx格式文件

1)改服务是利用thrift rpc宽肩实现，可以使用php调用golang服务去解析大型xlsx格式文件
，phpexcel由于性能问题不能解析生成很大的xlsx文件，所以才这样实现

2.配置文件 config.ini

[addr]

#服务ip地址

ip=127.0.0.1

#服务端口号

port=9098

[filepath]

#生成文件存放路径 必须以 / 结尾 ，文件路径可以自定义 必须是绝对路径 如果linux /server/... 必须有权限

filepath =E:/gows/src/Thrift/server/

[rootPath]

#根目录 

path=E:/gows/src/Thrift/server/

3.调用方法：
1）php调用解析xlsx 格式文件文件夹 phpThriftdemo/demo.php

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

    const ThriftPort = 9098;            //UserServer端口

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



$startTime = getMillisecond();

//php调用把文件或者json生成xlsx格式文件 开始

// $item = array();

// $item["path"] = "E:/gows/src/Thrift/server/272bb94d-1335-480b-be3c-81643a31a0e1-1499911449301.txt";

// $item["type"] = "path"; //如果type=path就是传入txt文件去解析成excel txt格式必须 以制表符 "\t" 分割 \n 结尾  例如：83328838666    80200800400|0.3 -   -   打车租车

// $item["json"] ='[["电话","idfa","ema","银行内部","世界你好","自由自在"],["17782277993","","","1","2",""],["18055108397","","","1","3",""],["13392855539","","","1","4",""],["18927414041","","","1","5",""],["18166455656","","","1","6",""],["13359228840","","","1","7",""],["18122168581","","","1","8",""],["15364710990","","","1","9",""],["13302906338","","","1","10",""],["13302906338","","","1","11",""],["17729835585","","","1","12",""],["13385280050","","","1","13",""],["18092180167","","","1","14",""],["15397051735","","","1","15",""],["13343432089","","","1","16",""],["18149206100","","","1","17",""],["18051089967","","","1","18",""],["15325717241","","","1","19",""],["18007131543","","","1","20",""],["17761737679","","","1","21",""],["18938870368","","","1","22",""],["18072948085","","","1","23",""],["18902279322","","","1","24",""],["18011521286","","","1","25",""],["15388130519","","","1","26",""],["18092335801","","","1","27",""],["18938692080","","","1","28",""],["18100623968","","","1","29",""],["18162504215","","","1","30",""],["18982186270","","","1","31",""],["15389093213","","","1","32",""],["18925025088","","","1","33",""],["17752557636","","","1","34",""],["18024018524","","","1","35",""],["18980733252","","","1","36",""],["13324588576","","","1","37",""],["13380062514","","","1","38",""],["13380062514","","","1","39",""],["18080950817","","","1","40",""],["18903060721","","","1","41",""],["18903060721","","","1","42",""],["15353661190","","","1","43",""],["18080112966","","","1","44",""],["17791840945","","","1","45",""],["17791840945","","","1","46",""],["18086069180","","","1","47",""],["17725166759","","","1","48",""],["17708385118","","","1","49",""],["17720150490","","","1","50",""],["13310977223","","","1","51",""],["13388186315","","","1","52",""],["13373931817","","","1","53",""],["18907193669","","","1","54",""],["17712410842","","","1","55",""],["17723600266","","","1","56",""],["15397646317","","","1","57",""],["18995657668","","","1","58",""],["13389276130","","","1","59",""],["18086602768","","","1","60","自由自在"],["17783528569","","","1","61",""],["17784222842","","","1","62",""],["17784222842","","","1","63",""],["17721528435","","","1","64",""],["15527797792","","","1","65","自由自在"],["13316834758","","","1","66",""],["13316887913","","","1","67",""],["15378178676","","","1","68",""],["18926199823","","","1","69",""],["18106599198","","","1","70","自由自在"],["13368101072","","","1","71",""],["18958166977","","","1","72",""],["13316115200","","","1","73","自由自在"],["18981961715","","","","",""]]';

//php调用把文件或者json生成xlsx格式文件 结束
 
//解析xlsx参数 ---开始

$item=array();

$item["path"]= "E:/gows/src/Thrift/server/demo.xlsx";

$item["type"]= "json";//返回形式 type=json 就是以json格式返回（建议不要用如果xlsx文件很大有可能内存溢出），type=path 就是返货 xxx.txt文件  把xlsx解析成txt并返回路径 数据以 \n换行 \t分割

//解析xlsx参数 ---结束


$ThrifClient = new ThrifClient();

$type=1; //如果type=2代表要生成 xlsx文件  type=1解析xlsx 文件

//输出json格式

$res = $ThrifClient->SocketT($type, $item);

echo $res;

$resArr = json_decode($res, true);

print_r($resArr);

$endTime = getMillisecond();

echo "\n使用时间:" . $endTime . "-" . $startTime . " =" . ($endTime - $startTime) . " 毫秒\n";


?>

返回值：

{"detail":"E:\\gows\\src\\ThriftExcel/31666e37d87b3490a2c471edc1d1d07c.xlsx","message":"返回成功","status":0}

Array
(
    [detail] => E:\gows\src\ThriftExcel/31666e37d87b3490a2c471edc1d1d07c.xlsx

    [message] => 返回成功

    [status] => 0
)

4.启动服务  可执行文件在目录里面

1)如果linix

./main_linux

2)如果windows

main.exe
