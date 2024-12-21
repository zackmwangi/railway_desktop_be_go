package api

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	"net/http"

	v1 "github.com/zackmwangi/railway_desktop_be_go/internal/api_proto/v1"
	"github.com/zackmwangi/railway_desktop_be_go/internal/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type (

	// type
	MyHttpServer struct {
		Http   *http.Server
		Logger *zap.Logger
		Hmux   *gin.Engine
		Config *config.AppConfig
	}

	//type
	MyGrpcServer struct {
		Grpc   *grpc.Server
		Logger *zap.Logger
		Gmux   *runtime.ServeMux
		Config *config.AppConfig
	}

	// type
	Servers struct {
		AppConfig *config.AppConfig
		//HealthCheckerLive  health.Checker
		//HealthCheckerReady health.Checker

		Grpc   *MyGrpcServer
		Http   *MyHttpServer
		stopFn sync.Once
	}
)

// ################################################################
// # HTTP
func (s *MyHttpServer) Run(ctx context.Context, httpAddr string) error {

	httpListener, err := net.Listen("tcp", httpAddr)

	if err != nil {
		s.Logger.Fatal("error on http address : " + httpAddr)
		os.Exit(1)
	}

	s.Http = &http.Server{
		//TODO: add these to main config
		Handler:        s.Hmux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.Logger.Sugar().Infof("HTTP service listening at %s", httpAddr)
	return s.Http.Serve(httpListener)
}

func (s *MyHttpServer) Shutdown(ctx context.Context) {
	s.Logger.Sugar().Infof("HTTP service gracefully shutting down ")

	if s.Http != nil {
		if err := s.Http.Shutdown(ctx); err != nil {
			s.Logger.Fatal("graceful shutdown of HTTP service failed ")
		}
	}
}

// ################################################################
// # gRPC
func (s *MyGrpcServer) Run(ctx context.Context, grpcAddress string) error {

	var lc net.ListenConfig

	lis, err := lc.Listen(ctx, "tcp", grpcAddress)

	if err != nil {

		s.Logger.Sugar().Fatalf("error on grpc address : %s", err)

	}

	s.Grpc = grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 5 * time.Minute,
		}),
	)

	reflection.Register(s.Grpc)

	//############################################################
	//Register gRPC base services

	MybackendSvcServerImpl := NewMybackendSvcServerImpl(s.Config)
	v1.RegisterMybackendSvcServer(s.Grpc, MybackendSvcServerImpl)

	//########################
	/*
		//Register grpc-gateway for HTTP-like API
		grpcDialOpts := []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		}

		errX := v1.RegisterMybackendSvcHandlerFromEndpoint(ctx, s.Gmux, grpcAddress, grpcDialOpts)

		if errX != nil {
			s.Logger.Sugar().Fatalf("failed to serve gRPC gateway : %s", err)
		}
	*/
	//########################
	s.Logger.Sugar().Infof("gRPC service listening at %v", lis.Addr())
	return s.Grpc.Serve(lis)

}

func (s *MyGrpcServer) Shutdown(ctx context.Context) {
	s.Logger.Sugar().Infof("gRPC service gracefully shutting down ")

	done := make(chan struct{}, 1)

	go func() {
		if s.Grpc != nil {
			s.Grpc.GracefulStop()
		}
		done <- struct{}{}
	}()

	select {
	case <-done:
	case <-ctx.Done():
		if s.Grpc != nil {
			s.Grpc.Stop()
		}
		s.Logger.Fatal("graceful shutdown of gRPC server failed. ")
	}
}

//################################################################
// both

func (s *Servers) Run(ctx context.Context) (err error) {

	var ec = make(chan error, 2)

	ctx, cancel := context.WithCancel(ctx)

	grpcGwMux := s.getGrpcGwMux(s.AppConfig)
	s.Grpc = &MyGrpcServer{
		Logger: s.AppConfig.AppLogger,
		Gmux:   grpcGwMux,
		Config: s.AppConfig,
	}

	httpMux := s.getHttpMux(s.AppConfig)
	s.Http = &MyHttpServer{
		Logger: s.AppConfig.AppLogger,
		Hmux:   httpMux,
		Config: s.AppConfig,
	}

	go func() {
		err := s.Grpc.Run(ctx, net.JoinHostPort(s.AppConfig.AppListenHostname, s.AppConfig.AppListenPortGrpc))
		if err != nil {
			err = fmt.Errorf("gRPC server stop : %w", err)
		}
		ec <- err
	}()

	go func() {
		err := s.Http.Run(ctx, net.JoinHostPort(s.AppConfig.AppListenHostname, s.AppConfig.AppListenPortHttp))
		if err != nil {
			err = fmt.Errorf("HTTP Server stop : %w", err)
		}
		ec <- err
	}()

	var es []string

	for i := 0; i < cap(ec); i++ {
		if err := <-ec; err != nil {
			es = append(es, err.Error())
			if ctx.Err() == nil {
				s.Shutdown(context.Background())
			}
		}
	}

	if len(es) > 0 {
		err = errors.New(strings.Join(es, ", "))
	}

	cancel()

	return err
}

func (s *Servers) getGrpcGwMux(c *config.AppConfig) *runtime.ServeMux {
	return runtime.NewServeMux()
}

func (s *Servers) getHttpMux(c *config.AppConfig) *gin.Engine {

	httpRoutingEngine := InitHTTPRoutingEngine(c)
	AddHttpEndpoints(httpRoutingEngine, c, s)
	return httpRoutingEngine
}

func (s *Servers) Shutdown(ctx context.Context) {
	s.stopFn.Do(func() {
		s.Http.Shutdown(ctx)
		s.Grpc.Shutdown(ctx)
	})
}
