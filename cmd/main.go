package main

import (
	"context"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/alexliesenfeld/health"
	"github.com/zackmwangi/railway_desktop_be_go/internal/api"
	"github.com/zackmwangi/railway_desktop_be_go/internal/config"
	applogger "github.com/zackmwangi/railway_desktop_be_go/internal/pkg/applogger"
)

func main() {

	//CPU Max
	runtime.GOMAXPROCS(runtime.NumCPU())
	//#Init App config
	appConfig := config.InitAppConfig()

	//Init logger
	AppLogger := applogger.InitAppLogger()
	defer AppLogger.Sync()
	appConfig.AppLogger = AppLogger

	healthCheckerReady := health.NewChecker(health.WithCacheDuration(1*time.Second),
		health.WithTimeout(10*time.Second))

	appConfig.HealthCheckerReady = healthCheckerReady

	healthCheckerReady.Start()
	healthCheckerLive := health.NewChecker()

	appConfig.HealthCheckerLive = healthCheckerLive

	//################################################################################################

	s := &api.Servers{
		AppConfig: appConfig,
	}

	//################################################################################################
	//Handle Signals
	ec := make(chan error, 2)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

	go func() {
		ec <- s.Run(context.Background())
	}()

	//######
	var err error

	select {

	case err = <-ec:
	case <-ctx.Done():

		haltCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		s.Shutdown(haltCtx)
		stop()

		err = <-ec
	}

	if err != nil {
		AppLogger.Sugar().Infof("service %s shutdown : %s", appConfig.SvcName, err)
	}

}
