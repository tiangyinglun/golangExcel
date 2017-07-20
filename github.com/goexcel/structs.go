package goexcel

const (
	Sufxlsx  = ".xlsx"
	Suftxt  = ".txt"
)

//定义Excel返回类型
type DataContent struct {
	Data [][]string
}

//定义返回类型
type CallBack struct {
	RBack map[string]interface{}
}
