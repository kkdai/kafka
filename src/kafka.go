package kafka

import (
	"fmt"
	"net"
	"os"
)

const (
	MAGIC_DEFAULT = 0
	NETWORK       = "tcp"
)


type Broker struct {
	topic     string
	partition int
	hostname  string
}

func newBroker(hostname string, topic string, partition int) *Broker {
	return &Broker{topic: topic,
		partition: partition,
		hostname:  hostname}
}



func (b *Broker) connect() (conn *net.TCPConn, error os.Error) {
	raddr, err := net.ResolveTCPAddr(NETWORK, b.hostname)
	conn, err = net.DialTCP(NETWORK, nil, raddr)
	if err != nil {
		fmt.Println("Fatal Error: ", err)
		return nil, err
	}
	return conn, error
}

