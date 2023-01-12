package vscale

import (
	"fmt"
)

type SSHKey struct {
	ID   int    `json:"id,omitempty"`
	Key  string `json:"key,omitempty"`
	Name string `json:"name,omitempty"`
}

type SSHKeyCreateRequest struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

type SSHService interface {
	List() (*[]SSHKey, *Response, error)
	Create(*SSHKeyCreateRequest) (*SSHKey, *Response, error)
	Delete(int) (*Response, error)
}

type SSHServiceOp struct {
	client *Client
}

const (
	sshBaseUrl = "/v1/sshkeys"
)

var _ SSHService = &SSHServiceOp{}

func (s *SSHServiceOp) List() (*[]SSHKey, *Response, error) {
	path := sshBaseUrl
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	sshkeys := &[]SSHKey{}
	resp, err := s.client.Do(req, sshkeys)
	if err != nil {
		return nil, nil, err
	}
	return sshkeys, resp, err
}

func (s *SSHServiceOp) Create(createRequest *SSHKeyCreateRequest) (*SSHKey, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	req, err := s.client.NewRequest("POST", sshBaseUrl, createRequest)
	if err != nil {
		return nil, nil, err
	}

	sshkey := new(SSHKey)
	resp, err := s.client.Do(req, sshkey)
	if err != nil {
		return nil, nil, err
	}

	return sshkey, resp, err
}

func (s SSHServiceOp) Delete(sshkeyID int) (*Response, error) {
	if sshkeyID < 1 {
		return nil, NewArgError("sshkeyID", "cannot be less than 1")
	}

	path := fmt.Sprintf("%s/%d", sshBaseUrl, sshkeyID)

	req, err := s.client.NewRequest("DELETE", path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(req, nil)
}

// import (
// 	"encoding/json"
// 	"fmt"
// )
//
// type SSHKey struct {
// 	ID        int
// 	Name, Key string
// }
//
// func (c *Client) SSHKeys() (*[]SSHKey, error) {
// 	ret, err := c.get("sshkeys", nil)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	keys := make([]SSHKey, 0)
// 	err = json.Unmarshal([]byte(ret), &keys)
// 	return &keys, err
// }
//
// func (c *Client) NewSSHKey(params map[string]interface{}) (*SSHKey, error) {
// 	ret, err := c.post("sshkeys", params)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	keys := make([]SSHKey, 0)
// 	fmt.Println(ret)
// 	err = json.Unmarshal([]byte(ret), &keys)
//
// 	if len(keys) > 0 {
// 		return &keys[0], nil
// 	} else {
// 		return nil, err
// 	}
// }
//
// func (c *Client) DeleteSSHKey(id int) (string, error) {
// 	return c.delete(fmt.Sprintf("sshkeys/%d", id), nil)
// }
