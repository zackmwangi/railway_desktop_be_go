package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func AddGrpcGatewayEndpoints(httpRoutingEngine *gin.Engine, s *runtime.ServeMux) *gin.Engine {

	vx := "v1"
	vxRouterGroup := httpRoutingEngine.Group(vx).Use().(*gin.RouterGroup)

	vxRouterGroup.GET("/user/*{http_to_grpc_gateway}", gin.WrapH(s))
	vxRouterGroup.POST("/user/*{http_to_grpc_gateway}", gin.WrapH(s))

	return httpRoutingEngine
}
