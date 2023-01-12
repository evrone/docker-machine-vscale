package vscale

import (
	"fmt"
)

const (
	scaletsBasePath = "/v1/scalets"
	tasksBasePath   = "/v1/tasks"
)

type ScaletService interface {
	List() (*[]Scalet, *Response, error)
	GetByID(int) (*Scalet, *Response, error)
	Create(*ScaletCreateRequest) (*Scalet, *Response, error)
	Restart(int) (*Scalet, *Response, error)
	Rebuild(*ScaletRebuildRequest) (*Scalet, *Response, error)
	Halt(int) (*Scalet, *Response, error)
	Start(int) (*Scalet, *Response, error)
	UpdatePlan(*ScaletUpdatePlanRequest) (*Scalet, *Response, error)
	Delete(int) (*Scalet, *Response, error)
	AddSSHKeyToScalet(*SSHKeyAppendRequest) (*Scalet, *Response, error)
	Tasks() (*[]ScaletTask, *Response, error)
}

type ScaletServiceOp struct {
	client *Client
}

var _ ScaletService = &ScaletServiceOp{}

type Scalet struct {
	Name           string         `json:"name,omitempty"`
	Hostname       string         `json:"hostname,omitempty"`
	Locked         bool           `json:"locked,omitempty"`
	Location       string         `json:"location,omitempty"`
	Rplan          string         `json:"rplan,omitempty"`
	Active         bool           `json:"active,omitempty"`
	Keys           []SSHKey       `json:"keys,omitempty"`
	PublicAddress  *ScaletAddress `json:"public_address,omitempty"`
	Status         string         `json:"status,omitempty"`
	MadeFrom       string         `json:"made_from,omitempty"`
	CTID           int            `json:"ctid,omitempty"`
	PrivateAddress *ScaletAddress `json:"private_address,omitempty"`
}

type ScaletAddress struct {
	Address string `json:"address,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	Gateway string `json:"gateway,omitempty"`
}

type ScaletCreateRequest struct {
	MakeFrom string `json:"make_from"`
	Rplan    string `json:"rplan"`
	DoStart  bool   `json:"do_start"`
	Name     string `json:"name"`
	Keys     []int  `json:"keys"`
	Password string `json:"password"`
	Location string `json:"location"`
}

type ScaletRebuildRequest struct {
	ID       int
	Password string `json:"password"`
}

type ScaletUpdatePlanRequest struct {
	ID    int
	Rplan string `json:"rplan"`
}

type SSHKeyAppendRequest struct {
	CTID int
	Keys []int `json:"keys"`
}

type ScaletTask struct {
	ID         string `json:"id,omitempty"`
	Location   string `json:"location,omitempty"`
	InsertDate string `json:"d_insert,omitempty"`
	StartDate  string `json:"d_start,omitempty"`
	EndDate    string `json:"d_end,omitempty"`
	Done       bool   `json:"done,omitempty"`
	ScaletId   int    `json:"scalet,omitempty"`
	Error      bool   `json:"error,omitempty"`
	Method     string `json:"method,omitempty"`
}

func (s Scalet) String() string {
	return Stringify(s)
}

func (s ScaletServiceOp) List() (*[]Scalet, *Response, error) {
	path := scaletsBasePath
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	scalets := &[]Scalet{}
	resp, err := s.client.Do(req, scalets)
	if err != nil {
		return nil, nil, err
	}
	return scalets, resp, err
}

func (s ScaletServiceOp) GetByID(ctid int) (*Scalet, *Response, error) {
	if ctid < 1 {
		return nil, nil, NewArgError("scaletID", "cannot be less than 1")
	}
	path := fmt.Sprintf("%s/%d", scaletsBasePath, ctid)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	scalet := &Scalet{}
	resp, err := s.client.Do(req, scalet)
	if err != nil {
		return nil, nil, err
	}
	return scalet, resp, err
}

func (s ScaletServiceOp) Create(createRequest *ScaletCreateRequest) (*Scalet, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest("POST", scaletsBasePath, createRequest)
	if err != nil {
		return nil, nil, err
	}

	scalet := new(Scalet)
	resp, err := s.client.Do(req, scalet)
	if err != nil {
		return nil, nil, err
	}

	return scalet, resp, err
}

func (s ScaletServiceOp) Restart(ctid int) (*Scalet, *Response, error) {
	if ctid < 1 {
		return nil, nil, NewArgError("scaletID", "cannot be less than 1")
	}
	path := fmt.Sprintf("%s/%d/restart", scaletsBasePath, ctid)
	req, err := s.client.NewRequest("PATCH", path, nil)
	if err != nil {
		return nil, nil, err
	}

	scalet := &Scalet{}
	resp, err := s.client.Do(req, scalet)
	if err != nil {
		return nil, nil, err
	}
	return scalet, resp, err
}

func (s *ScaletServiceOp) Rebuild(rebuildRequest *ScaletRebuildRequest) (*Scalet, *Response, error) {
	if rebuildRequest == nil {
		return nil, nil, NewArgError("rebuildRequest", "cannot be nil")
	}

	path := fmt.Sprintf("%s/%d/rebuild", scaletsBasePath, rebuildRequest.ID)
	req, err := s.client.NewRequest("PATCH", path, rebuildRequest)
	if err != nil {
		return nil, nil, err
	}

	scalet := &Scalet{}
	resp, err := s.client.Do(req, scalet)
	if err != nil {
		return nil, nil, err
	}
	return scalet, resp, err
}

func (s *ScaletServiceOp) Halt(ctid int) (*Scalet, *Response, error) {
	if ctid < 1 {
		return nil, nil, NewArgError("scaletID", "cannot be less than 1")
	}
	path := fmt.Sprintf("%s/%d/stop", scaletsBasePath, ctid)
	req, err := s.client.NewRequest("PATCH", path, nil)
	if err != nil {
		return nil, nil, err
	}

	scalet := &Scalet{}
	resp, err := s.client.Do(req, scalet)
	if err != nil {
		return nil, nil, err
	}
	return scalet, resp, err
}

func (s *ScaletServiceOp) Start(ctid int) (*Scalet, *Response, error) {
	if ctid < 1 {
		return nil, nil, NewArgError("scaletID", "cannot be less than 1")
	}
	path := fmt.Sprintf("%s/%d/start", scaletsBasePath, ctid)
	req, err := s.client.NewRequest("PATCH", path, nil)
	if err != nil {
		return nil, nil, err
	}

	scalet := &Scalet{}
	resp, err := s.client.Do(req, scalet)
	if err != nil {
		return nil, nil, err
	}
	return scalet, resp, err
}

func (s *ScaletServiceOp) UpdatePlan(rplanUpdateRequest *ScaletUpdatePlanRequest) (*Scalet, *Response, error) {
	if rplanUpdateRequest == nil {
		return nil, nil, NewArgError("rplanUpdateRequest", "cannot be nil")
	}

	if rplanUpdateRequest.ID < 1 {
		return nil, nil, NewArgError("rplanUpdateRequest.ID", "cannot be less than 1")
	}

	if rplanUpdateRequest.Rplan == "" {
		return nil, nil, NewArgError("rplanUpdateRequest.rplan", "cannot be empty")
	}

	path := fmt.Sprintf("%s/%d/upgrade", scaletsBasePath, rplanUpdateRequest.ID)
	req, err := s.client.NewRequest("POST", path, rplanUpdateRequest)
	if err != nil {
		return nil, nil, err
	}

	scalet := new(Scalet)
	resp, err := s.client.Do(req, scalet)
	if err != nil {
		return nil, nil, err
	}
	return scalet, resp, err
}

func (s *ScaletServiceOp) Delete(ctid int) (*Scalet, *Response, error) {
	if ctid < 1 {
		return nil, nil, NewArgError("scaletID", "cannot be less than 1")
	}
	path := fmt.Sprintf("%s/%d", scaletsBasePath, ctid)
	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, nil, err
	}

	scalet := &Scalet{}
	resp, err := s.client.Do(req, scalet)
	if err != nil {
		return nil, nil, err
	}
	return scalet, resp, err
}

func (s *ScaletServiceOp) Tasks() (*[]ScaletTask, *Response, error) {
	req, err := s.client.NewRequest("GET", tasksBasePath, nil)
	if err != nil {
		return nil, nil, err
	}

	tasks := &[]ScaletTask{}
	resp, err := s.client.Do(req, tasks)
	if err != nil {
		return nil, nil, err
	}
	return tasks, resp, err
}

func (s *ScaletServiceOp) AddSSHKeyToScalet(appendSSHKeyRequest *SSHKeyAppendRequest) (*Scalet, *Response, error) {
	if appendSSHKeyRequest == nil {
		return nil, nil, NewArgError("appendSSHKeyRequest", "cannot be nil")
	}
	if appendSSHKeyRequest.CTID < 1 {
		return nil, nil, NewArgError("appendSSHKeyRequest.CTID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s/scalets/%d", sshBaseUrl, appendSSHKeyRequest.CTID)
	req, err := s.client.NewRequest("PATCH", path, appendSSHKeyRequest)
	if err != nil {
		return nil, nil, err
	}

	scalet := new(Scalet)
	resp, err := s.client.Do(req, scalet)
	if err != nil {
		return nil, nil, err
	}
	return scalet, resp, err
}
