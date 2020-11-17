package main

import (
	"fmt"
	"log"
	"net"

	"github.com/myproxy/cmd"
	"github.com/myproxy/local"
)

// co
const (
	DefaultListenAddr = ":7448"
)

var version = "master"

func main() {

	log.SetFlags(log.Lshortfile)

	config := &cmd.Config{
		ListenAddr: DefaultListenAddr,
	}

	config.ReadConfig()
	config.SaveConfig()

	// ListenAddr string `json:"listen"`
	// RemoteAddr string `json:"remote"`
	// Password   string `json:"password"`

	// config := &cmd.Config{
	// 	ListenAddr: ":44087",
	// 	RemoteAddr: "34.94.217.10:44087",
	// 	Password:   "3wukioA4RoQMTci/3gJJif0U0+h/5Htn1nnRCXIvw3Us6dsz1yYY9M6ptOxUYKBqpqzEQO1rbWJoFgSRW3RZlLn3jJ74eoHyVQNp3PM0D4iwzKsuvOqNs/AewGyGuE69Y/mHYahTpXd9DUHLtWWZnKMRoaIZMJ87hY6TQwqaNhXx/zE6Ss3Bb14Q9Q7Jxl/nwrcgMlgaykcI4iPSrU9uF1cSE4PVkNCvKY8bleDYJAAfTEXu5XM8NbqY5kinPh2XsV27Jftwx4LP1N1aJ+EG/O+uKJL2K9l8vlxQnVE5AbJ2Qkt4tlb+fgX6RC2qImbjKpYH2us3IT9SPRxxZMWLmw==",
	// }

	mylocal, err := local.NewLsLocal(config.Password, config.ListenAddr, config.RemoteAddr)
	if err != nil {
		log.Fatalln(err)
	}

	log.Fatalln(mylocal.Listen(func(listenAddr *net.TCPAddr) {
		log.Println(fmt.Sprintf(`
	myproxy-local:%s Started successfuly, configuration as follow:
	listen to local address:
	%s
	remote server address:
	%s
	password:
	%s`, version, listenAddr, config.RemoteAddr, config.Password))
	}))

}
