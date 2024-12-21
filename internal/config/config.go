package config

import (
	"log"

	"github.com/alexliesenfeld/health"
	"github.com/joho/godotenv"
	"github.com/zackmwangi/railway_desktop_be_go/internal/pkg/envo"
	"go.uber.org/zap"
)

type (
	AppConfig struct {
		AppLogger          *zap.Logger
		HealthCheckerLive  health.Checker
		HealthCheckerReady health.Checker

		//################################

		SvcName string
		AppEnv  string

		AppListenHostname string
		AppListenPortHttp string
		AppListenPortGrpc string

		HealthEndpointPrefix string
		HealthEndpointLive   string
		HealthEndpointReady  string
	}
)

func InitAppConfig() *AppConfig {

	envConfigFile := ".env"
	err := godotenv.Load(envConfigFile)

	if err != nil {
		log.Printf("could not load .env file %s", err)
	} else {
		log.Printf("loaded .env file at %s", envConfigFile)
	}

	//#HARD DEFAULTS
	//########################
	svcNameDefault := "mybackend-svc-api"

	appEnvDefault := "dev"

	//appEnvDefault := "prod"

	appListenHostnameDefault := "127.0.0.1"
	//appListenHostnameDefault := "172.20.0.28"
	appListenPortHttpDefault := "8081"
	appListenPortGrpcDefault := "8082"

	healthEndpointPrefixDefault := "/"
	healthEndpointLiveDefault := "health"
	healthEndpointReadyDefault := "ready"

	//################################################################################################

	//# Swap hard defaults with .env/Configmap supplied values
	svcName := envo.EnvString("SVC_NAME", svcNameDefault)
	appEnv := envo.EnvString("APP_ENV", appEnvDefault)
	//ginMode := envo.EnvString("GIN_MODE", ginModeDefault)

	appListenHostname := envo.EnvString("APP_LISTEN_HOSTNAME", appListenHostnameDefault)
	appListenPortHttp := envo.EnvString("APP_LISTEN_PORT_HTTP", appListenPortHttpDefault)
	appListenPortGrpc := envo.EnvString("APP_LISTEN_PORT_GRPC", appListenPortGrpcDefault)

	healthEndpointPrefix := envo.EnvString("HEALTH_ENDPOINT_PREFIX", healthEndpointPrefixDefault)
	healthEndpointLive := envo.EnvString("HEALTH_ENDPOINT_LIVE", healthEndpointLiveDefault)
	healthEndpointReady := envo.EnvString("HEALTH_ENDPOINT_READY", healthEndpointReadyDefault)

	//################################################################################################

	return &AppConfig{

		SvcName: svcName,
		AppEnv:  appEnv,

		AppListenHostname: appListenHostname,
		AppListenPortHttp: appListenPortHttp,
		AppListenPortGrpc: appListenPortGrpc,

		HealthEndpointPrefix: healthEndpointPrefix,
		HealthEndpointLive:   healthEndpointLive,
		HealthEndpointReady:  healthEndpointReady,
	}

}
