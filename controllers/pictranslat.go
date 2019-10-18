package controllers

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

type PictranslatController struct {
	beego.Controller
}

func (c *PictranslatController) Get() {
	c.TplName = "pic4translat.tpl"
}
func (c *PictranslatController) Post() {
	f, h, _ := c.GetFile("uploadfile")
	ext := path.Ext(h.Filename) //获取文件后缀
	fileNewName := string(time.Now().Format("20060102150405")) + strconv.Itoa(time.Now().Nanosecond()) + ext
	f.Close()
	file_path := beego.AppConfig.String("tempfilepath") + fileNewName
	c.SaveToFile("uploadfile", file_path)

	cmd("tesseract", file_path, file_path, "-l", "chi_sim+eng")
	cmd("python", "ts.py", file_path+".txt")
	temp_file, er := os.Open(file_path + ".txt" + ".translate")
	if er != nil {
		fmt.Println(er)
		c.Controller.Data["data"] = "识别失败"
	} else {
		data, _ := ioutil.ReadAll(temp_file)
		c.Controller.Data["data"] = string(data)
	}
	os.Remove(file_path + ".txt" + ".translate")
	os.Remove(file_path + ".txt")
	os.Remove(file_path)
}
