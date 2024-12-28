package entities

import "github.com/hasura/go-graphql-client"

type (
	CreateServiceMutation struct {
		ServiceCreate struct {
			Id        string `json:"id"`
			ProjectId string `json:"projectId"`
		} `graphql:"serviceCreate(input: $input)"`
	}

	ServiceSourceInputImage struct {
		Image graphql.String `json:"image"`
	}

	ServiceSourceInputRepo struct {
		Repo graphql.String `json:"repo"`
	}

	//################################
	DeleteServiceMutation struct {
		ServiceDelete bool `graphql:"serviceDelete(id: $id, environmentId: $environmentId)"`
	}
)
