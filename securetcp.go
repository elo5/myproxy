package myproxy

import (
	"io"
	"log"
	"net"
)

const (
	bufSize = 1024
)

// SecureTCPConn 加密传输的 TCP Socket
type SecureTCPConn struct {
	io.ReadWriteCloser
	Cipher *Cipher
}

// DecodeRead 从输入流里读取加密过的数据，解密后把原数据放到bs里面
func (ss *SecureTCPConn) DecodeRead(bs []byte) (n int, err error) {
	n, err = ss.Read(bs)
	if err != nil {
		return
	}
	ss.Cipher.Decode(bs[:n])
	return
}

// EncodeWrite 把放在bs里面的数据加密后，立即全部写入输入流
func (ss *SecureTCPConn) EncodeWrite(bs []byte) (n int, err error) {
	ss.Cipher.Encode(bs)
	return ss.Write(bs)
}

// EncodeCopy 从src中源源不断的 读取原数据加密后 写入dst，直到src中没有数据可以再读取
func (ss *SecureTCPConn) EncodeCopy(dst io.ReadWriteCloser) error {
	buf := make([]byte, bufSize)

	for {
		readCount, errRead := ss.Read(buf)
		if errRead != nil {
			if errRead != io.EOF {
				return errRead
			}
			return nil
		}

		if readCount > 0 {
			nSSTcpConn := &SecureTCPConn{
				ReadWriteCloser: dst,
				Cipher:          ss.Cipher,
			}
			// writeCount, errWrite := (&SecureTCPConn{
			// 	ReadWriteCloser: dst,
			// 	Cipher:          ss.Cipher,
			// }).EncodeWrite(buf[0:readCount])
			writeCount, errWrite := nSSTcpConn.EncodeWrite(buf[0:readCount])

			if errWrite != nil {
				return errWrite
			}

			if readCount != writeCount {
				return io.ErrShortWrite
			}
		}
	}
}

// DecodeCopy 从src中源源不断的 读取加密后 的数据解密后写入到dst，直到src中没有数据可以再读取
func (ss *SecureTCPConn) DecodeCopy(dst io.Writer) error {
	buf := make([]byte, bufSize)
	for {
		readCount, errRead := ss.DecodeRead(buf)
		if errRead != nil {
			if errRead != io.EOF {
				return errRead
			}
			return nil
		}
		if readCount > 0 {
			writeCount, errWrite := dst.Write(buf[0:readCount])
			if errWrite != nil {
				return errWrite
			}
			if readCount != writeCount {
				return io.ErrShortWrite
			}
		}
	}
}

// DialEncryptedTCP 连接
func DialEncryptedTCP(raddr *net.TCPAddr, cipher *Cipher) (*SecureTCPConn, error) {

	remoteConn, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {
		return nil, err
	}
	// Conn被关闭时直接清除所有数据 不管没有发送的数据
	remoteConn.SetLinger(0)

	return &SecureTCPConn{
		ReadWriteCloser: remoteConn,
		Cipher:          cipher,
	}, nil
}

// ListenEncryptedTCP listen
func ListenEncryptedTCP(laddr *net.TCPAddr, cipher *Cipher, handleConn func(localConn *SecureTCPConn), didListen func(listenAddr *net.TCPAddr)) error {

	listener, err := net.ListenTCP("tcp", laddr)
	if err != nil {
		return err
	}
	defer listener.Close()
	if didListen != nil {
		// didListen 可能有阻塞操作
		go didListen(listener.Addr().(*net.TCPAddr))
	}

	for {
		localConn, err := listener.AcceptTCP()
		if err != nil {
			log.Println(err)
			continue
		}
		// localConn被关闭时直接清除所有数据 不管没有发送的数据
		localConn.SetLinger(0)
		go handleConn(&SecureTCPConn{
			ReadWriteCloser: localConn,
			Cipher:          cipher,
		})

	}
}
