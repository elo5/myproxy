package myproxy

// Cipher 编码、解码 所用的密码
type Cipher struct {
	encodePassword *Password
	decodePassword *Password
}

// Encode 加密原始数据
func (cipher *Cipher) Encode(bs []byte) {

	for i, v := range bs {
		bs[i] = cipher.encodePassword[v]
	}
}

// Decode 解密数据
func (cipher *Cipher) Decode(bs []byte) {

	for i, v := range bs {
		bs[i] = cipher.decodePassword[v]
	}
}

// NewCipher 新建一个编码解码器
func NewCipher(encodePassword *Password) *Cipher {

	decodePassword := &Password{}
	for i, v := range encodePassword {
		encodePassword[i] = v
		decodePassword[v] = byte(i)
	}

	return &Cipher{
		encodePassword: encodePassword,
		decodePassword: decodePassword,
	}
}
