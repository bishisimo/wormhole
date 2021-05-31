//@Time : 2020/10/13 下午3:34
//@Author : bishisimo
package utils

import (
	"github.com/sirupsen/logrus"
	"net"
	"regexp"
	"strconv"
	"strings"
)

type ncInfo struct {
	Mac       string
	Mask      string
	Broadcast string
	Intranet  string
	Extranet  string
	AllIp     map[string]string
}

var NcInfo *ncInfo

func init() {
	initNetIp()
}

func initNetIp() {
	//validNet := map[string]bool{
	//	"WLAN":     true,
	//	"以太网":    true,
	//	"wlan0":    true,
	//	"eth0":     true,
	//	"eth1":     true,
	//	"es33":     true,
	//	"enp4s0f1": true,
	//	"wlp3s0":   true,
	//	"enp1s0":   true,
	//}
	var mac string
	var mask net.IPMask
	var intranet string
	var extranet string
	reg := regexp.MustCompile(`.*WLAN|以太网|Ethernet|eth.*|enp.*|es.*|wlp|en0.*|eno*`)
	allIp := make(map[string]string)
	netInterfaces, err := net.Interfaces()
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	for _, netInterface := range netInterfaces {
		if netInterface.Flags&net.FlagUp != 0 {
			keys := strings.Split(netInterface.Name, " ")
			key := keys[len(keys)-1]
			addrs, err := netInterface.Addrs()
			if err != nil {
				logrus.Error(err.Error())
				return
			}
			for _, address := range addrs {
				if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
					if ipNet.IP.To4() != nil && reg.Match([]byte(netInterface.Name)) && intranet == "" {
						mac = netInterface.HardwareAddr.String()
						mask = ipNet.Mask
						intranet = ipNet.IP.String()
					}
					allIp[key] = ipNet.IP.String()
				}
			}
		}
	}

	//resp, err := http.Get("http://localhost/api/ip/raw")
	//if err != nil {
	//	logrus.Error(err.Error())
	//}
	//if resp != nil {
	//	defer resp.Subs.Close()
	//	content, _ := ioutil.ReadAll(resp.Subs)
	//	extranet = string(content)
	//	allIp["ex"] = extranet
	//}
	if mask == nil {
		return
	}
	m := strings.Split(intranet, ".")
	broadcast := ""
	for i, s := range m {
		sd, err := strconv.Atoi(s)
		if err != nil {
			logrus.Error(err.Error())
			return
		}
		dd := sd&int(mask[i]) + 255 - int(mask[i])
		broadcast += strconv.Itoa(dd)
		if i < len(m)-1 {
			broadcast += "."
		}
	}
	NcInfo = &ncInfo{
		Mac:       mac,
		Mask:      mask.String(),
		Broadcast: broadcast,
		Intranet:  intranet,
		Extranet:  extranet,
		AllIp:     allIp,
	}
	if NcInfo.Intranet == "" {
		logrus.Fatalln("ip 无法获取！")
	}
}
