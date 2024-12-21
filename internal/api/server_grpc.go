package api

import (
	"context"

	v1 "github.com/zackmwangi/railway_desktop_be_go/internal/api_proto/v1"
	"github.com/zackmwangi/railway_desktop_be_go/internal/config"
)

type (
	MybackendSvcServerImpl struct {
		AppConfig *config.AppConfig
		v1.UnimplementedMybackendSvcServer
	}
)

func NewMybackendSvcServerImpl(
	ac *config.AppConfig,
) v1.MybackendSvcServer {
	return &MybackendSvcServerImpl{
		AppConfig: ac,
	}
}

func (s *MybackendSvcServerImpl) GetUserInfoById(ctx context.Context, req *v1.GetUserInfoByIdRequest) (*v1.GetUserInfoByIdResponse, error) {

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

func (s *MybackendSvcServerImpl) GetUserInfoByUsername(ctx context.Context, req *v1.GetUserInfoByUsernameRequest) (*v1.GetUserInfoByUsernameResponse, error) {

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