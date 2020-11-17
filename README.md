# myproxy
照着lightsocket工程自己写了一遍，编译及部署至自己的vps，并成功运行

1. 使用 `uname -a` 查看vps的系统内核版本等信息，如果是`x86_64`那么`GOARCH`应该是`amd64`

2. 编译服务端(cd至myproxy-server)

   `env GOOS=linux GOARCH=amd64 go build main.go`

3. 上传至服务端，使用等google cloud， SSH窗口，设置图标带有上传文件菜单，将编译后的main上传

