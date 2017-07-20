# golangExcel

1.php调用golang解析xlsx格式文件

1)改服务是利用thrift rpc宽肩实现，可以使用php调用golang服务去解析大型xlsx格式文件
，phpexcel由于性能问题不能解析生成很大的xlsx文件，所以才这样实现


# 2.配置文件 config.ini

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

# 3.调用方法： php调用解析xlsx 格式文件 见文件夹 phpThriftdemo/demo.php 

 # 解析xlsx文件

$item=array();

$item["path"]= "E:/gows/src/Thrift/server/demo.xlsx"; //文件地址

$item["type"]= "json";//返回形式 type=json 就是以json格式返回（建议不要用如果xlsx文件很大有可能内存溢出），type=path 就是返货 xxx.txt文件  把xlsx解析成txt并返回路径 数据以 \n换行 \t分割



# 生成xlsx
//---------------php调用把文件或者json生成xlsx格式文件 开始--------------------------
// $item = array();

// $item["path"] = "E:/gows/src/Thrift/server/272bb94d-1335-480b-be3c-81643a31a0e1-1499911449301.txt";

// $item["type"] = "path"; //如果type=path就是传入txt文件去解析成excel txt格式必须 以制表符 "\t" 分割 \n 结尾  例如：83328838666    80200800400|0.3 -   -   打车租车

// $item["json"] ='[["电话","idfa","ema","银行内部","世界你好","自由自在"],["17782277993","","","1","2",""],["18055108397","","","1","3",""],,["18981961715","","","","",""]]';
 

# 调用方式

$ThrifClient = new ThrifClient();

$type=1; //如果type=2代表要生成 xlsx文件  type=1解析xlsx 文件

//输出json格式

$res = $ThrifClient->SocketT($type, $item);


# 返回值：

{"detail":"E:\\gows\\src\\ThriftExcel/31666e37d87b3490a2c471edc1d1d07c.xlsx","message":"返回成功","status":0}

Array
(
    [detail] => E:\gows\src\ThriftExcel/31666e37d87b3490a2c471edc1d1d07c.xlsx

    [message] => 返回成功

    [status] => 0
)

# 4.启动服务  可执行文件在目录里面

1)如果linix

./main_linux

2)如果windows

main.exe
