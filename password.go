package myproxy

import (
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const passwordLength = 256

// Password length  256 byte
type Password [passwordLength]byte

func init() {
	// 更新随机种子，防止生成一样的随机密码
	rand.Seed(time.Now().Unix())
}

// 采用base64编码把密码转换为字符串
func (password *Password) String() string {
	return base64.StdEncoding.EncodeToString(password[:])
}

// func BytesToString(b []byte) string {
// 	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
// 	sh := reflect.StringHeader{bh.Data, bh.Len}
// 	return *(*string)(unsafe.Pointer(&sh))
// }

// ParsePassword 解析采用base64编码的字符串获取密码
func ParsePassword(passwordString string) (*Password, error) {
	bs, err := base64.StdEncoding.DecodeString(strings.TrimSpace(passwordString))
	if err != nil || len(bs) != passwordLength {
		return nil, errors.New("密码不合法")
	}
	password := Password{}
	copy(password[:], bs)
	bs = nil
	return &password, err
}

// RandomPassword 产生 256个byte随机组合的 密码，最后会使用base64编码为字符串存储在配置文件中
// 不能出现任何一个重复的byte位，必须又 0-255 组成，并且都需要包含
func RandomPassword() string {
	// 随机生成一个由  0~255 组成的 byte 数组
	intArr := rand.Perm(passwordLength)
	password := &Password{}

	// fmt.Println(intArr)
	// fmt.Println("----------------------------------- ")

	for i, v := range intArr {
		password[i] = byte(v)
		fmt.Println("for  loop :", i, " -- ", v)
		if i == v {
			// 确保不会出现如何一个byte位出现重复
			return RandomPassword()
		}
	}
	return password.String()
}
