package vscale

type ConfigurationsService interface {
	Rplans() (*[]Rplan, *Response, error)
}

type Rplan struct {
	ID        string   `json:"id,omitempty"`
	Memory    int      `json:"memory,omitempty"`
	Disk      int      `json:"disk,omitempty"`
	Locations []string `json:"locations,omitempty"`
	Network   int      `json:"network,omitempty"`
	Addresses int      `json:"addresses,omitempty"`
	Cpus      int      `json:"cpus,omitempty"`
	Templates []string `json:"templates,omitempty"`
}

type ConfigurationsServiceOp struct {
	client *Client
}

var _ ConfigurationsService = &ConfigurationsServiceOp{}

func (s *ConfigurationsServiceOp) Rplans() (*[]Rplan, *Response, error) {
	path := "/v1/rplans"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	collection := &[]Rplan{}
	resp, err := s.client.Do(req, collection)
	if err != nil {
		return nil, nil, err
	}
	return collection, resp, err
}
