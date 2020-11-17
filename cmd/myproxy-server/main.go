package main

import (
	"fmt"

	"github.com/myproxy"
)

var version = "master"

func main() {
	// log.SetFlags(log.Lshortfile)

	// // 优先从环境变量中获取监听端口
	// port, err := strconv.Atoi(os.Getenv("MY_PROXY_SERVER_PORT"))
	// // 服务端监听端口随机生成
	// if err != nil {
	// 	port, err = freeport.GetFreePort()
	// }

	// if err != nil {
	// 	port = 7448
	// }
	// config := &cmd.Config{}

	pwd := myproxy.RandomPassword()

	fmt.Printf("随机密码: %s\n", pwd)

}
