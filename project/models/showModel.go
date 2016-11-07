package models

import (
	//"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type TImage struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Path      string `json:"path"`
	Instr     string `json:"instr"`
	Rank      int    `json:"rank"`
	ImageType int    `json:"image_type"`
	CreatedAt string `json:"created_at"`
}

type TKImage struct {
	Kid int `json:"kid"`
	TImage
}

func init() {
	orm.RegisterModel(new(TImage))
}

func GetAllImageData() ([]*TImage, error) {
	o := orm.NewOrm()
	imglist := make([]*TImage, 0)
	_, err := o.QueryTable("t_image").All(&imglist)
	return imglist, err
}

func GetAllImageDataById() (TImage, error) {
	o := orm.NewOrm()
	var imglist TImage
	err := o.QueryTable("t_image").Filter("id", 1).One(&imglist)
	return imglist, err
}

func GetAlldata() []TKImage {
	data, err := GetAllImageData()
	if err != nil {
		fmt.Println("err")
	}
	List := make([]TKImage, len(data))
	for k, v := range data {
		List[k].Kid = k + 1
		List[k].TImage.Id = v.Id
		List[k].TImage.Title = v.Title
		List[k].TImage.Path = v.Path
		List[k].TImage.Instr = v.Instr
		List[k].TImage.ImageType = v.ImageType
		List[k].TImage.CreatedAt = v.CreatedAt
		List[k].TImage.Rank = v.Rank
	}
	return List
}
