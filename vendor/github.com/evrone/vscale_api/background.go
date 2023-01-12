package vscale

type BackgroundService interface {
	Locations() (*[]Location, *Response, error)
	Images() (*[]Image, *Response, error)
}

type BackgroundServiceOp struct {
	client *Client
}

var _ BackgroundService = &BackgroundServiceOp{}

type Location struct {
	ID                string   `json:"id,omitempty"`
	Active            bool     `json:"active,omitempty"`
	Description       string   `json:"description,omitempty"`
	Rplans            []string `json:"rplans,omitempty"`
	Templates         []string `json:"templates,omitempty"`
	PrivateNetworking bool     `json:"private_networking,omitempty"`
}

type Image struct {
	ID          string   `json:"id,omitempty"`
	Active      bool     `json:"active,omitempty"`
	Description string   `json:"description,omitempty"`
	Rplans      []string `json:"rplans,omitempty"`
	Locations   []string `json:"locations,omitempty"`
	Size        int      `json:"size,omitempty"`
}

func (s *BackgroundServiceOp) Locations() (*[]Location, *Response, error) {
	path := "/v1/locations"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	locations := &[]Location{}

	resp, err := s.client.Do(req, locations)
	if err != nil {
		return nil, nil, err
	}

	return locations, resp, err
}

func (s *BackgroundServiceOp) Images() (*[]Image, *Response, error) {
	path := "/v1/images"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	images := &[]Image{}
	resp, err := s.client.Do(req, images)
	if err != nil {
		return nil, nil, err
	}

	return images, resp, err
}
