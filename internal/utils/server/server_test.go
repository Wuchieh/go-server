package server_test

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"testing"
	"time"

	"github.com/Wuchieh/go-server/internal/utils/server"
)

func TestNew(t *testing.T) {
	s := server.New()
	s2 := server.New()

	var wg sync.WaitGroup
	//ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	wg.Go(func() {
		time.Sleep(10 * time.Second)
		s.Stop(ctx)
		cancel()
	})

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sc
		cancel()
	}()

	wg.Go(func() {
		err := s.Run(ctx, ":8080")
		if err != nil {
			t.Fatal(err)
		}
	})

	wg.Go(func() {
		err := s2.Run(ctx, ":8081")
		if err != nil {
			t.Fatal(err)
		}
	})

	wg.Wait()
}
