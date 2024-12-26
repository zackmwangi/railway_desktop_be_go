package railwaySvc

import (
	"context"

	//"github.com/Khan/genqlient/graphql"
	graphql "github.com/hasura/go-graphql-client"
	v1 "github.com/zackmwangi/railway_desktop_be_go/internal/api_proto/v1"
	entities "github.com/zackmwangi/railway_desktop_be_go/internal/entities/railway"
	"go.uber.org/zap"
)

type (
	RailwayServicesSvc interface {
		ServiceCreateFromImage(context.Context, *v1.ServiceCreateFromImageRequest) (*v1.ServiceCreateFromImageResponse, error)
		DeleteService(context.Context, *v1.DeleteServiceRequest) (*v1.DeleteServiceResponse, error)
	}

	RailwayServicesSvcStruct struct {
		Applogger    *zap.Logger
		graphlClient *graphql.Client
	}
)

// Services+Deployments
func NewRailwayServicesSvc(gc *graphql.Client, l *zap.Logger) RailwayServicesSvc {

	return &RailwayServicesSvcStruct{
		Applogger:    l,
		graphlClient: gc,
	}

}

func (s *RailwayServicesSvcStruct) ServiceCreateFromImage(ctx context.Context, req *v1.ServiceCreateFromImageRequest) (*v1.ServiceCreateFromImageResponse, error) {
	//Execute against Railway Graphql API

	userInputs := entities.ServiceCreateFromImageRequestGqlInputs{
		ProjectId:     req.ProjectId,
		EnvironmentId: req.EnvironmentId,
		Source: entities.SourceInput{
			Image: req.ImageUrl,
		},
		Name: "My Fresh mage",
	}

	variableMix := map[string]interface{}{
		"input": userInputs,
	}

	//mutation := `
	//mutation serviceCreate($input: ServiceCreateInput!) {
	mutation := `
	mutation serviceCreate{
	serviceCreate(input: $input) {
			id
		}
	}
	`
	//var mutationResponse entities.ServiceCreateResponseGql

	err := s.graphlClient.Mutate(ctx, &mutation, variableMix) // &mutationResponse)

	if err != nil {
		s.Applogger.Error(err.Error())
		return &v1.ServiceCreateFromImageResponse{}, err
	}

	//
	return &v1.ServiceCreateFromImageResponse{
		ServiceId: "NewSVCID-RESP",
		Error:     nil,
	}, nil
}

func (s *RailwayServicesSvcStruct) DeleteService(ctx context.Context, req *v1.DeleteServiceRequest) (*v1.DeleteServiceResponse, error) {

	return &v1.DeleteServiceResponse{
		ServiceId:     "NewSVCID-RESP",
		EnvironmentId: "NewEnvID-RESP",
		Error:         nil,
	}, nil

}
