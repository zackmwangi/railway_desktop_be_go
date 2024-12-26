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

// Services+Deployments
// ServiceCreateFromImage()
func (s *MybackendGrpcSvcServerImpl) ServiceCreateFromImage(ctx context.Context, req *v1.ServiceCreateFromImageRequest) (*v1.ServiceCreateFromImageResponse, error) {

	//read the service image url

	return s.railwayGraphqlServicesCollection.RailwayServicesSvc.ServiceCreateFromImage(ctx, req)

}

func (s *MybackendGrpcSvcServerImpl) DeleteService(ctx context.Context, req *v1.DeleteServiceRequest) (*v1.DeleteServiceResponse, error) {
	return &v1.DeleteServiceResponse{}, nil
}

//Variables

//Volumes

//
