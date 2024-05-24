//go:generate go run catalog_gen.go
//go:generate go fmt catalog.go

package govee

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func Run(ctx context.Context) {
	g := new(errgroup.Group)

	br := NewBroadcaster(10 * time.Second)
	g.Go(func() error { return br.Run(ctx) })

	lsner := NewListener()
	g.Go(func() error { return lsner.Run(ctx) })

	if err := g.Wait(); err != nil {
		fmt.Println(err.Error())
	}
}
