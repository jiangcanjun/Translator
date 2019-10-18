package controllers

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

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
