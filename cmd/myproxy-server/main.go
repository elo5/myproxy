package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/myproxy"
	"github.com/myproxy/cmd"
	"github.com/myproxy/server"
	"github.com/phayes/freeport"
)

var version = "master"

func main() {
	log.SetFlags(log.Lshortfile)

	// 优先从环境变量中获取监听端口
	port, err := strconv.Atoi(os.Getenv("MY_PROXY_SERVER_PORT"))
	// 服务端监听端口随机生成
	if err != nil {
		port, err = freeport.GetFreePort()
	}

	if err != nil {
		port = 7448 // 随机端口失败就采用 7448
	}
	// 默认配置
	config := &cmd.Config{
		ListenAddr: fmt.Sprintf(":%d", port),
		// 密码随机生成
		Password: myproxy.RandomPassword(),
	}

	config.ReadConfig()
	config.SaveConfig()

	myserver, err := server.NewLsServer(config.Password, config.ListenAddr)
	if err != nil {
		log.Fatalln(err)
	}

	log.Fatalln(myserver.Listen(func(listenAddr *net.TCPAddr) {
		log.Println(fmt.Sprintf(`
myproxy-server:%s  started successfully, configuration as follow:
Listen to address:
%s
password:
%s`, version, listenAddr, config.Password))
	}))

	// --------------  test code  --------------

	// pwd := myproxy.RandomPassword()

	// fmt.Printf("随机密码: %s\n", pwd)

	// a := rand.Int()
	// b := rand.Intn(100)

	// fmt.Println(a)
	// fmt.Println(b)

	// rand.Seed(time.Now().Unix())

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(rand.Intn(100))
	// }

}

// func Test1() {
// 	rand.Seed(time.Now().Unix())
// 	for i := 0; i < 10; i++ {
// 		//设置随机数种子
// 		bytes := make([]byte, 5)
// 		for i := 0; i < 5; i++ {
// 			b := rand.Intn(26) + 65
// 			fmt.Println(b)
// 			bytes[i] = byte(b)
// 		}
// 		fmt.Println(string(bytes))
// 	}
// }
