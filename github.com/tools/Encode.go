package tools

const (
	Success     = 0
	SuffixError = iota
	NoExist
	ParamsError
)

var Message = map[int]string{
	Success:     "返回成功",
	SuffixError: "文件后缀名错误",
	NoExist:     "文件不存在",
	ParamsError: "参数错误",
}
