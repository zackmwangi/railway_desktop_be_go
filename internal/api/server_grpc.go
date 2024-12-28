package api

import (
	"context"

	v1 "github.com/zackmwangi/railway_desktop_be_go/internal/api_proto/v1"
	"github.com/zackmwangi/railway_desktop_be_go/internal/config"
	railwaySvc "github.com/zackmwangi/railway_desktop_be_go/internal/services/railway"
)

type (
	MybackendGrpcSvcServerImpl struct {
		railwayGraphqlServicesCollection *railwaySvc.RailwayGraphqlServicesCollection
		v1.UnimplementedMybackendGrpcSvcServer
	}
)

//RailwayGraphqlServicesCollection railwaySvc.RailwayGraphqlServicesCollection

func NewMybackendGrpcSvcServerImpl(
	ac *config.AppConfig,
) v1.MybackendGrpcSvcServer {

	railwayGraphqlServicesCollection := railwaySvc.InitRailwayGraphqlServicesCollection(ac)

	return &MybackendGrpcSvcServerImpl{
		railwayGraphqlServicesCollection: railwayGraphqlServicesCollection,
	}
}

func (s *MybackendGrpcSvcServerImpl) GetUserInfoById(ctx context.Context, req *v1.GetUserInfoByIdRequest) (*v1.GetUserInfoByIdResponse, error) {

	userIdById := req.GetUserId()

	userNameById := "JohnCena1ById"
	userFirsnameById := "JohnfnameById"

	res := &v1.GetUserInfoByIdResult{
		UserId:        userIdById,
		UserName:      userNameById,
		UserFirstname: userFirsnameById,
	}

	return &v1.GetUserInfoByIdResponse{
		Result: res,
		Error:  nil,
	}, nil
}

func (s *MybackendGrpcSvcServerImpl) GetUserInfoByUsername(ctx context.Context, req *v1.GetUserInfoByUsernameRequest) (*v1.GetUserInfoByUsernameResponse, error) {

	userNameByUsername := req.GetUserName()
	userIdByUsername := "UUID-JOHNCENA"
	userFirstnameByUsername := "JohnfnameByUsername"

	res := &v1.GetUserInfoByUsernameResult{
		UserName:      userNameByUsername,
		UserId:        userIdByUsername,
		UserFirstname: userFirstnameByUsername,
	}

	return &v1.GetUserInfoByUsernameResponse{
		Result: res,
		Error:  nil,
	}, nil
}

//################################
//Railway
//Projects+Environments

func (s *MybackendGrpcSvcServerImpl) FetchProjects(ctx context.Context, req *v1.FetchProjectRequest) (*v1.FetchProjectResponse, error) {
	return s.railwayGraphqlServicesCollection.RailwayProjectsSvc.FetchProjectsQuery(ctx, req)
}

//###############
// Services+Deployments

func (s *MybackendGrpcSvcServerImpl) CreateService(ctx context.Context, req *v1.CreateServiceRequest) (*v1.CreateServiceResponse, error) {
	/*
		switch req.BaseResourceType {
		case v1.BaseResourceType:

			break

		default:
			//empty service

			break

		}
	*/

	return s.railwayGraphqlServicesCollection.RailwayServicesSvc.CreateService(ctx, req)

}

/*
func (s *MybackendGrpcSvcServerImpl) CreateServiceFromEmpty(ctx context.Context, req *v1.CreateServiceRequest) (*v1.CreateServiceResponse, error) {

	return s.railwayGraphqlServicesCollection.RailwayServicesSvc.CreateServiceFromEmpty(ctx, req)

}

// CreateServiceFromImage()
func (s *MybackendGrpcSvcServerImpl) CreateServiceFromImage(ctx context.Context, req *v1.CreateServiceFromImageRequest) (*v1.CreateServiceResponse, error) {

	return s.railwayGraphqlServicesCollection.RailwayServicesSvc.CreateServiceFromImage(ctx, req)

}

func (s *MybackendGrpcSvcServerImpl) CreateServiceFromRepo(ctx context.Context, req *v1.CreateServiceFromRepoRequest) (*v1.CreateServiceResponse, error) {

	return s.railwayGraphqlServicesCollection.RailwayServicesSvc.CreateServiceFromRepo(ctx, req)

}
*/

func (s *MybackendGrpcSvcServerImpl) DeleteService(ctx context.Context, req *v1.DeleteServiceRequest) (*v1.DeleteServiceResponse, error) {
	//return &v1.DeleteServiceResponse{}, nil
	return s.railwayGraphqlServicesCollection.RailwayServicesSvc.DeleteService(ctx, req)
}

//###############
//Variables

//###############
//Volumes

//
