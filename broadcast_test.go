package govee

import (
	"context"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBroadcast(t *testing.T) {
	bcast := &Broadcaster{
		interval: 1 * time.Second,
		address:  "127.0.0.1:4100",
		message:  []byte(`test message`),
	}
	ctx, cancel := context.WithCancel(context.TODO())

	go bcast.Run(ctx)

	addr, err := net.ResolveUDPAddr("udp4", "239.255.255.250:4100")
	if err != nil {
		assert.NoError(t, err)
	}

	conn, err := net.ListenMulticastUDP("udp4", nil, addr)
	if err != nil {
		assert.NoError(t, err)
	}

	buffer := make([]byte, 8192)
	n, _, err := conn.ReadFromUDP(buffer)
	if err != nil {
		assert.NoError(t, err)
	}

	cancel()

	assert.Equal(t, string("test message"), string(buffer[:n]))
}
