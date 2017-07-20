package goexcel

const (
	Success     = 0
	ParamsError = iota
	NoExist
	SuffixErr
	NoKnowErr
	CreateFail
)

var Message = map[int]string{
	Success:     "返回成功",
	ParamsError: "参数错误",
	NoExist:     "文件不存在",
	SuffixErr:   "文件名类型错误",
	NoKnowErr:   "未知错误",
	CreateFail:  "创建文件失败",
}
