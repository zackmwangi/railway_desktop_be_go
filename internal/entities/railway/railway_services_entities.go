package entities

type (
	ServiceCreateFromImageRequestGqlInputs struct {
		ProjectId     string      `json:"projectId"`
		EnvironmentId string      `json:"environmentId"`
		Source        SourceInput `json:"source"`
		Name          string      `json:"name,omitempty"`
	}

	SourceInput struct {
		Image string `json:"image"`
	}

	//ServiceCreateFromImageResponseGql struct {
	ServiceCreateResponseGql struct {
		//ServiceCreate struct {
		Id string `json:"id"`
		//}
	}

	/*
		ServiceCreateMutation struct {
			ServiceCreate struct {
				Id string `json:"id"`
			} `json:"serviceCreate(input: $input)"`
		}
	*/

	DeleteServiceRequestGql struct {
	}

	DeleteServiceResponseGql struct {
	}
)
