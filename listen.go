package govee

import (
	"context"
	"fmt"
	"net"
)

type Listener struct {
	Result  chan Device
	address string
}

func NewListener() *Listener {
	return &Listener{
		Result:  make(chan Device),
		address: "239.255.255.250:4002",
	}
}

func (l Listener) Run(ctx context.Context) error {
	addr, err := net.ResolveUDPAddr("udp4", l.address)
	if err != nil {
		return err
	}

	conn, err := net.ListenMulticastUDP("udp4", nil, addr)
	if err != nil {
		return err
	}

	conn.SetReadBuffer(8192)

	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			buffer := make([]byte, 8192)
			_, _, err := conn.ReadFromUDP(buffer)
			if err != nil {
				return err
			}
			fmt.Println(string(buffer))
		}
	}
}
