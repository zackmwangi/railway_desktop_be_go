package railwaySvc

import (

	//"github.com/Khan/genqlient/graphql"

	"context"

	graphql "github.com/hasura/go-graphql-client"
	v1 "github.com/zackmwangi/railway_desktop_be_go/internal/api_proto/v1"
	"github.com/zackmwangi/railway_desktop_be_go/internal/config"

	"testing"
)

func TestNewRailwayServicesSvc(t *testing.T) {

	appConfig := config.InitAppConfig()
	gc := graphql.NewClient("http://localhost:", nil)
	l := appConfig.AppLogger

	got := NewRailwayServicesSvc(gc, l)

	want := (RailwayServicesSvc)(got)

	if got != want {
		t.Errorf("TestNewRailwayServicesSvc fail: got %s, wanted %s", got, want)
	}

}

func TestCreateService(t *testing.T) {

	appConfig := config.InitAppConfig()
	gc := graphql.NewClient("http://localhost:", nil)
	l := appConfig.AppLogger

	svc := NewRailwayServicesSvc(gc, l)
	ctx := context.Background()
	req := &v1.CreateServiceRequest{}

	got, err := svc.CreateService(ctx, req)

	want, err2 := (*v1.CreateServiceResponse)(got), (error)(err)

	if (got != want) && (err != err2) {
		t.Errorf("TestCreateService fail: got %s, wanted %s", got, want)
	}

}

/*
func TestCreateServiceFromEmpty(t *testing.T) {

}

func TestCreateServiceFromImage(t *testing.T) {

}

func TestCreateServiceFromRepo(t *testing.T) {

}

func TestSendCreateServiceRequest(t *testing.T) {

}

func TestDeleteService(t *testing.T) {

}
*/
