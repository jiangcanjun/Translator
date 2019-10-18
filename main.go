package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strconv"
	"time"
)

func cmdpython(command string, arg ...string) {
	cmd := exec.Command(command, arg...)
	stdout, err := cmd.StdoutPipe()
	if err != nil { //获取输出对象，可以从该对象中读取输出结果
		fmt.Println(err)
	}
	defer stdout.Close()               // 保证关闭输出流
	if err = cmd.Start(); err != nil { // 运行命令
		fmt.Println(err)
	}
	if opBytes, err := ioutil.ReadAll(stdout); err != nil { // 读取输出结果
		fmt.Println(err)
	} else {
		fmt.Println(string(opBytes))
	}
}
func cmd(command string, arg ...string) {
	cmd := exec.Command(command, arg...)
	stdout, err := cmd.StdoutPipe()
	if err != nil { //获取输出对象，可以从该对象中读取输出结果
		fmt.Println(err)
	}
	defer stdout.Close()               // 保证关闭输出流
	if err = cmd.Start(); err != nil { // 运行命令
		fmt.Println(err)
	}
	if opBytes, err := ioutil.ReadAll(stdout); err != nil { // 读取输出结果
		fmt.Println(err)
	} else {
		fmt.Println(string(opBytes))
	}
}
func main() {
	http.HandleFunc("/img2txt", img2txt)
	http.HandleFunc("/tsimg", transimg)
	err := http.ListenAndServe(":7373", nil)
	if err != nil {
		fmt.Println("服务器启动失败", err.Error())
		return
	}

}
func img2txt(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		writer.Write([]byte(tpl))
	} else if request.Method == "POST" {

		request.ParseMultipartForm(32 << 20)
		//接收客户端传来的文件 uploadfile 与客户端保持一致
		file, handler, err := request.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		//上传的文件保存在ppp路径下
		ext := path.Ext(handler.Filename) //获取文件后缀
		fileNewName := string(time.Now().Format("20060102150405")) + strconv.Itoa(time.Now().Nanosecond()) + ext

		f, err := os.OpenFile("./"+fileNewName, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)

		cmd("tesseract", fileNewName, fileNewName, "-l", "chi_sim+eng")
		temp_file, er := os.Open(fileNewName + ".txt")
		if er != nil {
			fmt.Println(er)
			fmt.Fprintln(writer, "识别失败")
		} else {
			data, _ := ioutil.ReadAll(temp_file)
			fmt.Fprintln(writer, string(data))
		}
		os.Remove(fileNewName + ".txt")
		os.Remove(fileNewName)
	}
}
func transimg(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		writer.Write([]byte(tpl))
	} else if request.Method == "POST" {
		request.ParseMultipartForm(32 << 20)
		//接收客户端传来的文件 uploadfile 与客户端保持一致
		file, handler, err := request.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		//上传的文件保存在ppp路径下
		ext := path.Ext(handler.Filename) //获取文件后缀
		fileNewName := string(time.Now().Format("20060102150405")) + strconv.Itoa(time.Now().Nanosecond()) + ext

		f, err := os.OpenFile("./"+fileNewName, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		cmd("tesseract", fileNewName, fileNewName, "-l", "chi_sim+eng")
		cmd("python", "ts.py", fileNewName+".txt")
		temp_file, er := os.Open(fileNewName + ".txt" + ".translate")
		if er != nil {
			fmt.Println(er)
			fmt.Fprintln(writer, "识别失败")
		} else {
			data, _ := ioutil.ReadAll(temp_file)
			fmt.Fprintln(writer, string(data))
		}
		os.Remove(fileNewName + ".txt" + ".translate")
		os.Remove(fileNewName + ".txt")
		os.Remove(fileNewName)
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte(tpl))
}

const tpl = `<html>
<head>
<title>上传文件</title>
</head>
<body>
<form enctype="multipart/form-data" action="/tsimg" method="post">
<input type="file" name="uploadfile">
<input type="hidden" name="token" value="{...{.}...}">
<input type="submit" value="upload">
</form>
</body>
</html>`
