package proch

import (
	"io/ioutil"
	"encoding/json"
)

type prochSetting struct {
	Profiles []wlanProfile `json:"profiles"`
}

type wlanProfile struct {
	Ssid string `json:"ssid"`
	ProxyEnable bool `json:"proxyEnable"`
	ProxyServer string `json:"proxyServer"`
	ProxyOverride string `json:"proxyOverride"`
}

type jsonLoadable interface {
	Load(filepath string) ([]wlanProfile, error)
}

type jsonLoader struct {
	jsonLoadable
}

var _ jsonLoadable = &jsonLoader{}

func NewJsonLoader() jsonLoadable {
	return &jsonLoader{}
}

func (jl *jsonLoader) Load(filepath string) ([]wlanProfile, error) {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var ps prochSetting
	if err := json.Unmarshal(bytes, &ps); err != nil {
		return nil, err
	}
	
	return ps.Profiles, nil
}