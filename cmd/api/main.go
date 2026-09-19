package main

import (
	"context"
	"flag"
	"os/signal"
	"syscall"

	"github.com/D1ssolve/go-zero-backend-starter/internal/config"
	"github.com/D1ssolve/go-zero-backend-starter/internal/handler"
	"github.com/D1ssolve/go-zero-backend-starter/internal/svc"
	tgtransport "github.com/D1ssolve/go-zero-backend-starter/internal/transport/telegram"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/starter-api.yaml", "configuration file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	c.ExpandEnvironment()

	serviceContext, err := svc.NewServiceContext(context.Background(), c)
	if err != nil {
		logx.Must(err)
	}
	defer serviceContext.Close()

	server := rest.MustNewServer(c.RestConf)
	handler.RegisterHandlers(server, serviceContext)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if c.Telegram.Token != "" {
		telegram, err := tgtransport.New(c.Telegram.Token, serviceContext.Notes)
		if err != nil {
			logx.Must(err)
		}
		go telegram.Run(ctx)
		logx.Info("telegram transport started")
	} else {
		logx.Info("telegram transport disabled: TELEGRAM_BOT_TOKEN is empty")
	}

	go func() {
		logx.Infof("starting HTTP server at %s:%d", c.Host, c.Port)
		server.Start()
	}()

	<-ctx.Done()
	logx.Info("shutdown signal received")
	server.Stop()
}
