package main

import (
	"fmt"
	"github.com/andlabs/ui"
)

func myMain() {
	var w *ui.Window
	w = ui.NewWindow("Test", 400, 100) //创建窗口
	ui.AppQuit = w.Closing             //程序关闭，同时关闭主窗口
	ed1 := ui.NewLineEdit("Text")      //创建单行文本编辑框
	btn1 := ui.NewButton("Click")      //创建按钮
	g := ui.NewGrid(1,                 //创建网格，用来存放控件。这里可以多个网格嵌套的，只需要设置一下被显示的网格就行
		ed1,
		btn1)
	g.SetStretchy(0, 0) //设置网格里0,0位置的控件为被缩放控件
	g.SetFilling(1, 0)  //将1,0位置的控件设为填充区域
	w.Open(g)           //设置主显示的网格
mainloop:
	for {
		select {
		case <-w.Closing:
			break mainloop
		case <-btn1.Clicked: //按钮点击触发的事件
			fmt.Println("按钮被点击", ed1.Text())
		}
	}
}
func main() {
	err := ui.Go(myMain)
	if err != nil {
		panic(fmt.Errorf("error initializing UI library: %v", err))
	}
}
