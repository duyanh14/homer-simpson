package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"simpson/internal/helper/logger"
	"simpson/internal/registry"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	if err := registry.BuidlContainerV2(ctx); err != nil {
		panic(err)
	}

	logger.Logger.Info("Registry successfully")

	lisner, err := net.Listen("tcp", ":9090")

	fmt.Println(lisner, err, ctx)
	fmt.Println("*** start server ***")

	signals := make(chan os.Signal, 1)
	shutdown := make(chan bool, 1)
	signal.Notify(signals, os.Interrupt)
	go func() {
		<-signals
		_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		shutdown <- true
	}()
	fmt.Println("*** shutdown ***")
	time.Sleep(3 * time.Second)
	<-shutdown
}
