package vscale

import (
	"io/ioutil"

	"github.com/docker/machine/libmachine/ssh"
	api "github.com/evrone/vscale_api"
)

func (d *Driver) publicSSHKeyPath() string {
	return d.GetSSHKeyPath() + ".pub"
}

func (d *Driver) createSSHKey() (*api.SSHKey, error) {
	if err := ssh.GenerateSSHKey(d.GetSSHKeyPath()); err != nil {
		return nil, err
	}

	publicKey, err := ioutil.ReadFile(d.publicSSHKeyPath())
	if err != nil {
		return nil, err
	}

	createRequest := &api.SSHKeyCreateRequest{
		Name: d.MachineName,
		Key:  string(publicKey),
	}

	key, _, err := d.getClient().SSHKey.Create(createRequest)
	return key, err
}
