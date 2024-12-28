package railwaySvc

//Projects+Environments
import (

	//"github.com/Khan/genqlient/graphql"
	"context"

	graphql "github.com/hasura/go-graphql-client"
	v1 "github.com/zackmwangi/railway_desktop_be_go/internal/api_proto/v1"
	entities "github.com/zackmwangi/railway_desktop_be_go/internal/entities/railway"
	"go.uber.org/zap"
)

type (
	RailwayProjectsSvc interface {
		FetchProjectsQuery(context.Context, *v1.FetchProjectRequest) (*v1.FetchProjectResponse, error)
	}

	RailwayProjectsSvcStruct struct {
		Applogger    *zap.Logger
		graphlClient *graphql.Client
	}
)

func NewRailwayProjectssSvc(gc *graphql.Client, l *zap.Logger) RailwayProjectsSvc {

	return &RailwayProjectsSvcStruct{
		Applogger:    l,
		graphlClient: gc,
	}

}

func (s *RailwayProjectsSvcStruct) FetchProjectsQuery(ctx context.Context, req *v1.FetchProjectRequest) (*v1.FetchProjectResponse, error) {

	myProjects := make([]*v1.ProjectDetails, 0)

	queryResponse := &entities.FetchProjectsQuery{}
	err := s.graphlClient.Query(ctx, queryResponse, nil, graphql.OperationName("me"))

	if err != nil {
		return &v1.FetchProjectResponse{
			Projects: nil,
			Error:    v1.ErrorBadRequest,
		}, err
	}

	for _, projectEdges := range queryResponse.Me.Projects.Edges {

		projectDetails := projectEdges.Node
		myProject := &v1.ProjectDetails{ProjectId: string(projectDetails.Id), ProjectName: string(projectDetails.Name)}

		myProjects = append(myProjects, myProject)

	}

	return &v1.FetchProjectResponse{
		Projects: myProjects,
		Error:    nil,
	}, nil

}
