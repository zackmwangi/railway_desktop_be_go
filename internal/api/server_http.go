package api

import (
	"net/http"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"

	api "github.com/zackmwangi/railway_desktop_be_go/internal/api/http_endpoints"
	"github.com/zackmwangi/railway_desktop_be_go/internal/config"
	"go.uber.org/zap"
)

func initHTTPRoutingEngine(c *config.AppConfig) *gin.Engine {
	globalHTTPRoutingEngine := gin.New()
	var myLogger *zap.Logger

	if c.AppEnv == "prod" {
		gin.SetMode(gin.ReleaseMode)
		myLogger, _ = zap.NewProduction()

	} else {
		gin.SetMode(gin.DebugMode)
		myLogger, _ = zap.NewDevelopment()
	}

	globalHTTPRoutingEngine.Use(ginzap.Ginzap(myLogger, time.RFC3339, true))
	//TODO: Add HTTP middleware for validation
	//TODO: Add HTTP middleware for rate limiting/throttling
	return globalHTTPRoutingEngine
}

func AddHttpEndpoints(httpRoutingEngine *gin.Engine, c *config.AppConfig, s *Servers) *gin.Engine {

	api.AddHealthEndpoints(httpRoutingEngine, c, s.HealthCheckerLive, s.HealthCheckerReady)
	api.AddDocsEndpoints(httpRoutingEngine, c)
	api.AddGrpcGatewayEndpoints(httpRoutingEngine, c, s.Grpc.Gmux)

	httpRoutingEngine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error_info": "Requested resource not found"})
	})

	return httpRoutingEngine
}
