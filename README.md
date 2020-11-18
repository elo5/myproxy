# myproxy
照着lightsocks工程自己写了一遍，编译及部署至自己的vps，并成功运行

1. 使用 `uname -a` 查看vps的系统内核版本等信息，如果是`x86_64`那么`GOARCH`应该是`amd64`

2. 编译服务端(cd至myproxy-server)

   `env GOOS=linux GOARCH=amd64 go build main.go`

3. 上传至服务端，使用等google cloud， SSH窗口，设置图标带有上传文件菜单，将编译后的main上传

![local-start](https://github.com/elo5/myproxy/blob/main/assets/server-started.png)

![local-start](https://github.com/elo5/myproxy/blob/main/assets/local-started.png)

注意：代理类型为socks5

# 编译手机库
## 如果你喜欢，也可以继续折腾。比如，编译aar包给安卓使用

1. 直接编译成.so (在我发现gomobile之前)，步骤也很简单, 需要使用//export 方法名声明方法，必须包含main方法（实现可以为空）必须import “C”

``` js
import  "C"
//export yourMethod
func yourMethod(x, y int) int {
	return x + y
}

func main() {}
```
编译
```  bash
export ANDROID_NDK_HOME=你的ndk目录

export GOARCH=arm
export GOOS=android
export CGO_ENABLED=1
export CC=$ANDROID_NDK_HOME/toolchains/llvm/prebuilt/darwin-x86_64/bin/armv7a-linux-androideabi21-clang
go build -buildmode=c-shared -o output/android/armeabi-v7a/libmyproxy.so main.go

echo "Build armeabi-v7a success"

export GOARCH=arm64
export GOOS=android
export CGO_ENABLED=1
export CC=$ANDROID_NDK_HOME/toolchains/llvm/prebuilt/darwin-x86_64/bin/aarch64-linux-android21-clang
go build -buildmode=c-shared -o output/android/arm64-v8a/libmyproxy.so main.go

echo "Build arm64-v8a success"
```

然后新建CMake文件等等等等，不赘述了

2. 使用gomobile编译成aar，比编译成so简单很多, 方法名需要大写
```  bash
export ANDROID_NDK_HOME=你的ndk目录
export ANDROID_HOME=你的sdk目录
gomobile bind --target=android .
```

3. 编译成aar后新建android项目运行，电脑连接手机ip及端口，代理成功！