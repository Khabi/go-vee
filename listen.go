package govee

import (
	"context"
	"encoding/json"
	"net"
	"time"
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

	go func() {
		buffer := make([]byte, 8192)
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			return
		}

		var data scanRespose
		err = json.Unmarshal(buffer[:n], &data)
		if err != nil {
			return
		}

		ip := net.ParseIP(data.Msg.Data.IP)
		hw, err := net.ParseMAC(data.Msg.Data.Device)
		if err != nil {
			return
		}
		device := Device{
			IPAddr:   ip,
			MAC:      hw,
			LastSeen: time.Now(),
			Spec:     LookupBySKU(data.Msg.Data.Sku),
		}

		l.Result <- device
	}()

	<-ctx.Done()
	return nil

}
