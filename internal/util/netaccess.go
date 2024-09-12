package util

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	"golang.org/x/xerrors"
)

// Get host ip address
func GetIpAddr() (string, error) {
	// 获取所有网络接口的信息
	interfaces, err := net.Interfaces()
	if err == nil {
		// 遍历所有网络接口
		for _, iface := range interfaces {
			// 排除 loopback 接口和无效接口
			if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
				continue
			}

			// 获取当前接口的地址信息
			addrs, err := iface.Addrs()
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			// 遍历当前接口的地址
			for _, addr := range addrs {
				// 确保地址为 IP 地址
				ipNet, ok := addr.(*net.IPNet)
				if ok && !ipNet.IP.IsLoopback() {
					if ipNet.IP.To4() != nil {
						return ipNet.IP.String(), nil
					}
				}
			}
		}
	}
	return "", xerrors.New("Get ip address failed")
}

// Request login url
func UrlRequest(url string) (string, error) {
	respond, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer respond.Body.Close()

	result, err := ioutil.ReadAll(respond.Body)

	if err != nil {
		return "", err
	}
	return string(result), nil

}
