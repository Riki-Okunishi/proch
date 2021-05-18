package proch

import (
	"fmt"
)

type netshRunnerMock struct {
	connectSsid string
	disconnect bool
	profiles []string
	networks []string
	cssid string
}

var _ netshRunnable = &netshRunnerMock{}

func (nrm *netshRunnerMock) Connect(ssid string) error {
	if ssid != nrm.connectSsid {
		return fmt.Errorf("error [netshRunnerMock]: cannot connect")
	}
	return nil
}

func (nrm *netshRunnerMock) Disconnect() error {
	if !nrm.disconnect {
		return fmt.Errorf("error [netshRunnerMock]: cannot disconnect")
	}
	return nil
}

func (nrm *netshRunnerMock) ShowProfiles() []string {
	return nrm.profiles
}

func (nrm *netshRunnerMock) ShowNetworks() []string {
	return nrm.networks
}

func (nrm *netshRunnerMock) GetCurrentSsid() string {
	return nrm.cssid
}