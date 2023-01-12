package vscale

import (
	"fmt"
	"net"
	"time"

	"github.com/docker/machine/libmachine/drivers"
	"github.com/docker/machine/libmachine/log"
	"github.com/docker/machine/libmachine/mcnflag"
	"github.com/docker/machine/libmachine/state"
	api "github.com/evrone/vscale_api"
)

type Driver struct {
	*drivers.BaseDriver
	AccessToken string
	ScaletID    int
	ScaletName  string
	Rplan       string
	MadeFrom    string
	Location    string
	SSHKeyID    int
}

const (
	defaultRplan    = "small"
	defaultLocation = "spb0"
	defaultMadeFrom = "ubuntu_14.04_64_002_master"
)

func (d *Driver) GetCreateFlags() []mcnflag.Flag {
	return []mcnflag.Flag{
		mcnflag.StringFlag{
			EnvVar: "VSCALE_ACCESS_TOKEN",
			Name:   "vscale-access-token",
			Usage:  "Vscale access token",
		},
		mcnflag.StringFlag{
			EnvVar: "VSCALE_LOCATION",
			Name:   "vscale-location",
			Usage:  "Vscale location",
			Value:  defaultLocation,
		},
		mcnflag.StringFlag{
			EnvVar: "VSCALE_RPLAN",
			Name:   "vscale-rplan",
			Usage:  "Vscale rplan",
			Value:  defaultRplan,
		},
		mcnflag.StringFlag{
			EnvVar: "VSCALE_MADE_FROM",
			Name:   "vscale-made-from",
			Usage:  "Vscale made from",
			Value:  defaultMadeFrom,
		},
	}
}

func NewDriver(hostName, storePath string) *Driver {
	return &Driver{
		Rplan:    defaultRplan,
		Location: defaultLocation,
		MadeFrom: defaultMadeFrom,
		BaseDriver: &drivers.BaseDriver{
			MachineName: hostName,
			StorePath:   storePath,
		},
	}
}

func (d *Driver) GetSSHHostname() (string, error) {
	return d.GetIP()
}

// DriverName returns the name of the driver
func (d *Driver) DriverName() string {
	return "vscale"
}

func (d *Driver) SetConfigFromFlags(flags drivers.DriverOptions) error {
	d.AccessToken = flags.String("vscale-access-token")
	d.Location = flags.String("vscale-location")
	d.Rplan = flags.String("vscale-rplan")
	d.MadeFrom = flags.String("vscale-made-from")

	d.SwarmMaster = flags.Bool("swarm-master")
	d.SwarmHost = flags.String("swarm-host")
	d.SwarmDiscovery = flags.String("swarm-discovery")
	d.SSHPort = 22

	if d.AccessToken == "" {
		return fmt.Errorf("vscale driver requres the --vscale-access-token option")
	}

	return nil
}

func (d *Driver) getClient() *api.Client {
	return api.New(d.AccessToken)
}

func (d *Driver) PreCreateCheck() error {
	client := d.getClient()
	if client == nil {
		return fmt.Errorf("Cannot create Vscale client. Check --vscale-access-token option")
	}

	return nil
}

func (d *Driver) Create() error {
	log.Infof("Creating SSH key...")
	key, err := d.createSSHKey()
	if err != nil {
		return err
	}
	d.SSHKeyID = key.ID

	log.Infof("Creating Vscale scalet...")

	client := d.getClient()
	createRequest := &api.ScaletCreateRequest{
		MakeFrom: d.MadeFrom,
		Rplan:    d.Rplan,
		DoStart:  true,
		Name:     d.MachineName,
		Keys:     []int{d.SSHKeyID},
		Location: d.Location,
	}

	newScalet, _, err := client.Scalet.Create(createRequest)
	if err != nil {
		return err
	}

	d.ScaletID = newScalet.CTID

	log.Info("Waiting for IP address to be assigned to the Scalet...")

	for {
		newScalet, _, err = client.Scalet.GetByID(d.ScaletID)
		if err != nil {
			return err
		}

		if newScalet.PublicAddress != nil {
			d.IPAddress = newScalet.PublicAddress.Address
		}

		if d.IPAddress != "" {
			break
		}

		time.Sleep(1 * time.Second)
	}

	log.Debugf("Created scalet with ID: %v, IPAddress: %v", d.ScaletID, d.IPAddress)
	return nil
}

func (d *Driver) GetURL() (string, error) {
	ip, err := d.GetIP()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("tcp://%s", net.JoinHostPort(ip, "2376")), nil
}

func (d *Driver) GetState() (state.State, error) {
	scalet, _, err := d.getClient().Scalet.GetByID(d.ScaletID)
	if err != nil {
		return state.Error, err
	}

	switch scalet.Status {
	case "started":
		return state.Running, nil
	case "stopped":
		return state.Stopped, nil
	case "defined":
		return state.Starting, nil
	}
	return state.None, nil
}

func (d *Driver) Start() error {
	_, _, err := d.getClient().Scalet.Start(d.ScaletID)
	return err
}

func (d *Driver) Stop() error {
	_, _, err := d.getClient().Scalet.Halt(d.ScaletID)
	return err
}

func (d *Driver) Remove() error {
	client := d.getClient()
	_, _, err := client.Scalet.Delete(d.ScaletID)
	if err != nil {
		return err
	}

	_, err = client.SSHKey.Delete(d.SSHKeyID)
	if err != nil {
		return err
	}

	return nil
}

func (d *Driver) Restart() error {
	_, _, err := d.getClient().Scalet.Restart(d.ScaletID)
	return err
}

func (d *Driver) Kill() error {
	_, _, err := d.getClient().Scalet.Halt(d.ScaletID)
	return err
}
