package controllers

import (
	"fmt"
	_ "io/ioutil"
	_ "os"
	"path"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

type Pic2txtController struct {
	beego.Controller
}

func (c *Pic2txtController) Get() {
	c.TplName = "pic2txt.tpl"
}

func (c *Pic2txtController) Post() {
	f, h, _ := c.GetFile("uploadfile")
	ext := path.Ext(h.Filename) //获取文件后缀
	fileNewName := string(time.Now().Format("20060102150405")) + strconv.Itoa(time.Now().Nanosecond()) + ext
	f.Close()
	file_path := beego.AppConfig.String("tempfilepath") + fileNewName
	c.SaveToFile("uploadfile", fileNewName+ext)
	fmt.Println("asdf")
	cmd("tesseract", file_path, file_path, "-l", "chi_sim+eng")
	temp_file, er := os.Open(file_path + ".txt")
	if er != nil {
		fmt.Println(er)
		//fmt.Fprintln(writer, "识别失败")
		c.Controller.Data["data"] = "识别失败"
	} else {
		data, _ := ioutil.ReadAll(temp_file)
		//fmt.Fprintln(writer, string(data))
		c.Controller.Data["data"] = string(data)
	}
	os.Remove(file_path + ".txt")
	os.Remove(file_path)
}
