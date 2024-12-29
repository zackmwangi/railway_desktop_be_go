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
		CreateService(context.Context, *v1.CreateServiceRequest) (*v1.CreateServiceResponse, error)
		//CreateServiceFromImage(context.Context, *v1.CreateServiceRequest) (*v1.CreateServiceResponse, error)
		//CreateServiceFromRepo(context.Context, *v1.CreateServiceRequest) (*v1.CreateServiceResponse, error)
		//SendCreateServiceRequest(context.Context, map[string]any, string) (*v1.CreateServiceResponse, error)
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

func (s *RailwayServicesSvcStruct) CreateService(ctx context.Context, req *v1.CreateServiceRequest) (*v1.CreateServiceResponse, error) {

	serviceBaseResourceType := req.ServiceBaseResourceType.String()

	serviceBaseResourceTypeImage := v1.ServiceBaseResourceType_name[0]
	serviceBaseResourceTypeRepo := v1.ServiceBaseResourceType_name[0]

	switch serviceBaseResourceType {
	case string(serviceBaseResourceTypeImage):

		return s.CreateServiceFromImage(ctx, req)

	case string(serviceBaseResourceTypeRepo):

		return s.CreateServiceFromRepo(ctx, req)

	default:
		//empty service, no image/repo etc given
		return s.CreateServiceFromEmpty(ctx, req)

	}

}

func (s *RailwayServicesSvcStruct) CreateServiceFromEmpty(ctx context.Context, req *v1.CreateServiceRequest) (*v1.CreateServiceResponse, error) {

	//Use source:image from some hub
	/*
		type ServiceCreateInput struct {
			ProjectId     graphql.String `json:"projectId"`
			EnvironmentId graphql.String `json:"environmentId"`

			Name graphql.String `json:"name"`
		}

		reqInputs := ServiceCreateInput{
			ProjectId:     graphql.String(req.ProjectId),
			EnvironmentId: graphql.String(req.EnvironmentId),

			Name: graphql.String(req.ServiceName),
		}

		variableInput := map[string]any{
			"input": reqInputs,
		}

		return s.SendCreateServiceRequest(ctx, variableInput, req.EnvironmentId)
	*/

	return &v1.CreateServiceResponse{
		Error: v1.ErrorBadRequestNoServiceBaseResourceType,
	}, v1.ErrorBadRequestNoServiceBaseResourceType

}

func (s *RailwayServicesSvcStruct) CreateServiceFromImage(ctx context.Context, req *v1.CreateServiceRequest) (*v1.CreateServiceResponse, error) {

	//Use source:image from some hub
	type ServiceCreateInput struct {
		ProjectId     graphql.String                   `json:"projectId"`
		EnvironmentId graphql.String                   `json:"environmentId"`
		Source        entities.ServiceSourceInputImage `json:"source"`
		Name          graphql.String                   `json:"name"`
	}

	reqInputs := ServiceCreateInput{
		ProjectId:     graphql.String(req.ProjectId),
		EnvironmentId: graphql.String(req.EnvironmentId),
		Source: entities.ServiceSourceInputImage{
			Image: graphql.String(req.ServiceBaseResourceUrl),
		},
		Name: graphql.String(req.ServiceName),
	}

	variableInput := map[string]any{
		"input": reqInputs,
	}

	return s.SendCreateServiceRequest(ctx, variableInput, req.EnvironmentId)

}

func (s *RailwayServicesSvcStruct) CreateServiceFromRepo(ctx context.Context, req *v1.CreateServiceRequest) (*v1.CreateServiceResponse, error) {
	//Use source:repository and buildfile/Nix
	type ServiceCreateInput struct {
		ProjectId     graphql.String                  `json:"projectId"`
		EnvironmentId graphql.String                  `json:"environmentId"`
		Source        entities.ServiceSourceInputRepo `json:"source"`
		Name          graphql.String                  `json:"name"`
	}

	reqInputs := ServiceCreateInput{
		ProjectId:     graphql.String(req.ProjectId),
		EnvironmentId: graphql.String(req.EnvironmentId),
		Source: entities.ServiceSourceInputRepo{
			Repo: graphql.String(req.ServiceBaseResourceUrl),
		},
		Name: graphql.String(req.ServiceName),
	}

	variableInput := map[string]any{
		"input": reqInputs,
	}

	return s.SendCreateServiceRequest(ctx, variableInput, req.EnvironmentId)

}

func (s *RailwayServicesSvcStruct) SendCreateServiceRequest(ctx context.Context, variableInput map[string]any, environmentId string) (*v1.CreateServiceResponse, error) {

	var mutation entities.CreateServiceMutation

	err := s.graphlClient.Mutate(ctx, &mutation, variableInput, graphql.OperationName("serviceCreate"))

	if err != nil {
		return &v1.CreateServiceResponse{
			NewServiceId: "",
			Error:        v1.ErrorBadRequest,
		}, err
	}

	return &v1.CreateServiceResponse{
		NewServiceId:  mutation.ServiceCreate.Id,
		ProjectId:     mutation.ServiceCreate.ProjectId,
		EnvironmentId: environmentId,
		Error:         nil,
	}, nil
}

func (s *RailwayServicesSvcStruct) DeleteService(ctx context.Context, req *v1.DeleteServiceRequest) (*v1.DeleteServiceResponse, error) {

	variableInput := map[string]any{
		"id":            req.ServiceId,
		"environmentId": req.EnvironmentId,
	}

	var mutation entities.DeleteServiceMutation
	//var mutation entities.CreateServiceMutation

	err := s.graphlClient.Mutate(ctx, &mutation, variableInput, graphql.OperationName("serviceDelete"))

	if err != nil {
		return &v1.DeleteServiceResponse{
			ServiceId:     "",
			EnvironmentId: "",
			Error:         v1.ErrorBadRequest,
		}, err
	}

	if !mutation.ServiceDelete {
		return &v1.DeleteServiceResponse{
			ServiceId:     "",
			EnvironmentId: "",
			Error:         v1.ErrorInternal,
		}, err

	}

	return &v1.DeleteServiceResponse{
		ServiceId:     req.ServiceId,
		EnvironmentId: req.EnvironmentId,
		Error:         nil,
	}, nil

}
