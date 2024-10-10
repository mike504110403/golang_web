package socketdemo

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"os"

	"github.com/mike504110403/goutils/log"
)

type NetWork string

const (
	TCP  NetWork = "tcp"
	UDP  NetWork = "udp"
	ICMP NetWork = "ip4:icmp"
	// ...
)

// DialDemo : 用戶端連接服務端
// network : 網絡協議 -> tcp/udp/icmp/...
// add : 服務端地址
func DialDemo(network NetWork, add string) (net.Conn, error) {
	conn, err := net.Dial(string(network), add)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// SocketReadDemo : socket 讀取
func SocketReadDemo() {
	// 錯誤處理
	validateErr := func(err error) {
		if err != nil {
			log.Error(err.Error())
			os.Exit(1)
		}
	}

	if len(os.Args) != 2 {
		log.Error(fmt.Sprintf("Usage: %s host:port", os.Args[0]))
		os.Exit(1)
	}
	service := os.Args[1]
	conn, err := net.Dial("tcp", service)
	validateErr(err)
	// 發送請求
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	validateErr(err)

	result, err := fullyRead(conn)
	validateErr(err)
	log.Info(string(result))
}

// fullyRead : 完全讀取
func fullyRead(conn net.Conn) ([]byte, error) {
	defer conn.Close()

	result := bytes.NewBuffer(nil)
	buf := make([]byte, 512)
	// 循環讀取
	for {
		n, err := conn.Read(buf[0:1])
		result.Write(buf[0:n])
		if err != nil {
			// EOF 表示讀取完畢
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}
