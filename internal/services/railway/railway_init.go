package railwaySvc

import (
	"net/http"

	//"github.com/Khan/genqlient/graphql"
	graphql "github.com/hasura/go-graphql-client"
	"github.com/zackmwangi/railway_desktop_be_go/internal/config"
)

type (
	authedTransport struct {
		key     string
		wrapped http.RoundTripper
	}

	RailwayGraphqlServicesCollection struct {
		//Projects+Env
		//Services+Deployments
		RailwayServicesSvc RailwayServicesSvc
		//Variables
		//Volumes
	}
)

func (t *authedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer: "+t.key)
	return t.wrapped.RoundTrip(req)
}

// func (s *RailwayGraphqlServicesCollection)
func InitRailwayGraphqlServicesCollection(
	ac *config.AppConfig,
) *RailwayGraphqlServicesCollection {

	httpClient := http.Client{
		Transport: &authedTransport{
			key:     ac.RailwayGraphqlToken,
			wrapped: http.DefaultTransport,
		},
	}

	graphqlClient := graphql.NewClient(ac.RailwayGraphqlURL, &httpClient)
	railwayServicesSvc := NewRailwayServicesSvc(graphqlClient, ac.AppLogger)

	ac.AppLogger.Sugar().Infof("Sending to url %s, with token %s", ac.RailwayGraphqlURL, ac.RailwayGraphqlToken)

	return &RailwayGraphqlServicesCollection{
		//Applogger:          ac.AppLogger,
		RailwayServicesSvc: railwayServicesSvc,
	}

}
