package vscale

type NotificationsService interface {
	Get() (*NotificationsSettings, *Response, error)
	SetNotifyBalance(int) (*NotificationsSettings, *Response, error)
}

type NotificationsServiceOp struct {
	client *Client
}

var _ NotificationsService = &NotificationsServiceOp{}

type NotificationsRequest struct {
	NotifyBalance int `json:"notify_balance"`
}

type NotificationsSettings struct {
	NotifyBalance int    `json:"notify_balance"`
	Status        string `json:"status,omitempty"`
}

func (s *NotificationsServiceOp) Get() (*NotificationsSettings, *Response, error) {
	path := "/v1/billing/notify"
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	notifSett := new(NotificationsSettings)
	resp, err := s.client.Do(req, notifSett)
	if err != nil {
		return nil, nil, err
	}
	return notifSett, resp, err
}

func (s *NotificationsServiceOp) SetNotifyBalance(num int) (*NotificationsSettings, *Response, error) {
	path := "/v1/billing/notify"
	notifyRequest := &NotificationsRequest{NotifyBalance: num}
	req, err := s.client.NewRequest("PUT", path, notifyRequest)
	if err != nil {
		return nil, nil, err
	}

	notifSett := new(NotificationsSettings)
	resp, err := s.client.Do(req, notifSett)
	if err != nil {
		return nil, nil, err
	}
	return notifSett, resp, err
}
