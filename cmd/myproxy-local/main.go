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
