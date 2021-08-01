package main

import (
	"context"
	"github.com/xeonselina/go-practise/week-03/http"
	"golang.org/x/sync/errgroup"
"log"
"os"
"os/signal"
"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	eg, ctx := errgroup.WithContext(ctx)
	httpServ := http.NewHttp(ctx,8080)
	httpServDebug := http.NewHttp(ctx,8081)

	eg.Go(func() error {
		return httpServ.ListenAndServe()
	})
	eg.Go(func() error {
		return httpServDebug.ListenAndServe()
	})
	eg.Go(func() error {
		killSignal := make(chan os.Signal, 1)
		signal.Notify(killSignal, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
		for {
			select {
			case <-ctx.Done():
				log.Println("done received")
				return ctx.Err()
			case <-killSignal: //不知道为什么 ctrl+c 无法退出
				log.Println("kill signal shutdown")
				cancel()
				close(killSignal)
				return nil
			}
		}
	})
	if err := eg.Wait(); err != nil {
		panic(err)
	}

}