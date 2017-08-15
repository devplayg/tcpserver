package collectors

import (
	"crypto/tls"
	"errors"
	"net"
	"strings"
	"time"

	"github.com/devlang2/collectserver/event"
)

const (
	tcpTimeout = time.Duration(1000 * time.Millisecond)
	msgBufSize = 512
)

type Collector interface {
	Start(chan<- *event.Event) error
	Addr() net.Addr
}

func NewCollector(proto, addr string, tlsConfig *tls.Config) (Collector, error) {
	if strings.ToLower(proto) == "tcp" {
		return &TCPCollector{
			addrStr:   addr,
			tlsConfig: tlsConfig,
		}, nil
	} else if strings.ToLower(proto) == "udp" {
		udpAddr, err := net.ResolveUDPAddr("udp", addr)
		if err != nil {
			return nil, err
		}
		return &UDPCollector{addr: udpAddr}, nil
	}
	return nil, errors.New("Unsupport collector protocol")
}
