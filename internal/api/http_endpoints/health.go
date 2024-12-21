package api

import (
	"github.com/alexliesenfeld/health"
	"github.com/gin-gonic/gin"
	"github.com/zackmwangi/railway_desktop_be_go/internal/config"
)

func AddHealthEndpoints(httpRoutingEngine *gin.Engine, c *config.AppConfig, hcl health.Checker, hcr health.Checker) *gin.Engine {

	httpRoutingEngine.GET(c.HealthEndpointPrefix+c.HealthEndpointLive, gin.WrapH(health.NewHandler(hcl)))
	httpRoutingEngine.GET(c.HealthEndpointPrefix+c.HealthEndpointReady, gin.WrapH(health.NewHandler(hcr)))

	return httpRoutingEngine
}
