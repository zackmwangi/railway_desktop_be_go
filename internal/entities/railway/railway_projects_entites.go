package entities

import (
	"github.com/hasura/go-graphql-client"
)

type (
	FetchProjectNode struct {
		Id   graphql.String
		Name graphql.String
	}

	FetchProjectsQuery struct {
		Me struct {
			Projects struct {
				Edges []struct {
					Node FetchProjectNode
				}
			}
		}
	}
)
