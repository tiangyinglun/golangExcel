package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"project/models"
	"strconv"
	"time"
)

type ShowController struct {
	beego.Controller
}

func (c *ShowController) Index() {
	data, err := models.GetAllImageData()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
	c.Data["data"] = data
	c.TplName = "demo/index.html"
}

func (c *ShowController) Second() {
	data := models.GetAlldata()

	fmt.Println(data)
	c.Data["data"] = data
	c.TplName = "demo/second.html"
}

func (c *ShowController) Show() {
	c.TplName = "demo/show.html"
}

func (c *ShowController) Three() {
	data := models.GetAlldata()
	c.Data["json"] = &data
	c.ServeJSON()
}

func (c *ShowController) Upload() {
	// Title := c.GetString("title")
	// Instr := c.GetString("instr")
	// Rank := c.GetString("rank")
	// ImageType := c.GetString("image_type")
	JPG := strconv.FormatInt(time.Now().Unix(), 10)
	f, h, err := c.GetFile("path")
	defer f.Close()
	if err != nil {
		fmt.Println("getfile err ", err)
	} else {
		fmt.Println(h.Filename)
		Name := JPG + h.Filename
		c.SaveToFile("path", "E:/gows/src/project/"+Name)
	}
	c.TplName = "demo/show.html"
}
