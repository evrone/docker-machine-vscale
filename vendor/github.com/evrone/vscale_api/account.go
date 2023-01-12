package vscale

type AccountService interface {
	Get() (*Account, *Response, error)
}

type AccountServiceOp struct {
	client *Client
}

var _ AccountService = &AccountServiceOp{}

// https://developers.vscale.io/documentation/api/v1/#api-Account-GetAccount
type Account struct {
	ActivateDate string `json:"actdate,omitempty"`
	Country      string `json:"country,omitempty"`
	FaceID       int    `json:"face_id,string,omitempty"`
	ID           int    `json:"id,string,omitempty"`
	State        int    `json:"state,string,omitempty"`
	Email        string `json:"email,omitempty"`
	Name         string `json:"name,omitempty"`
	MiddleName   string `json:"middlename,omitempty"`
	SurName      string `json:"surname,omitempty"`
}

type accountInfo struct {
	Info *Account `json:"info,omitempty"`
}

func (r Account) String() string {
	return Stringify(r)
}

func (s *AccountServiceOp) Get() (*Account, *Response, error) {
	path := "/v1/account"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	info := new(accountInfo)

	resp, err := s.client.Do(req, info)
	if err != nil {
		return nil, nil, err
	}

	return info.Info, resp, err
}
