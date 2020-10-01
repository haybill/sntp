//package netapp
//author: btfak.com
//create: 2013-9-24
//update: 2016-08-22

package netapp

import (
	"fmt"
	"github.com/btfak/sntp/netevent"
	"github.com/btfak/sntp/sntp"
	"net"
	"strconv"
	"strings"
	"time"
)

var handler *Handler

type Handler struct {
	netevent.UdpHandler
}

func GetHandler() *Handler {
	if handler == nil {
		handler = new(Handler)
	}
	return handler
}

// DatagramReceived
// every udp request trigger it.
func (p *Handler) DatagramReceived(data []byte, addr net.Addr) {
	res, err := sntp.Serve(data)
	if err == nil {
		ip, port := spliteAddr(addr.String())
		p.UdpWrite(string(res), ip, port)
		fmt.Printf("%s: NTP server responded to incoming request from: %s\n", time.Now().Format("2006-01-02 15:04:05.000000000"), addr.String())
	} else {
		fmt.Printf("%s: NTP server generated an error (%v) while responding to: %s\n", time.Now().Format("2006-01-02 15:04:05.000000000"), err, addr.String())
	}
}

func spliteAddr(addr string) (string, int) {
	ip := strings.Split(addr, ":")[0]
	port := strings.Split(addr, ":")[1]
	p, _ := strconv.Atoi(port)
	return ip, p
}
