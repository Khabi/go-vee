package govee

import (
	"context"
	"net"
	"time"
)

type Broadcaster struct {
	interval time.Duration
	address  string
	message  []byte
}

func NewBroadcaster(interval time.Duration) *Broadcaster {
	return &Broadcaster{
		interval: interval,
		address:  "239.255.255.250:4001",
		message:  []byte(`{"msg":{"cmd":"scan","data":{"account_topic":"reserve"}}}`),
	}
}

func (b Broadcaster) Run(ctx context.Context) error {
	addr, err := net.ResolveUDPAddr("udp4", b.address)
	if err != nil {
		return err
	}

	conn, err := net.DialUDP("udp4", nil, addr)
	if err != nil {
		return err
	}

	ticker := time.NewTicker(b.interval)

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			_, err = conn.Write(b.message)
			if err != nil {
				return err
			}
		}
	}
}
