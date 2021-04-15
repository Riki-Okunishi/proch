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
	//Password string `json:"password"`
	ProxyEnable bool `json:"proxyEnable"`
	ProxyServer string `json:"proxyServer"`
	ProxyOverride string `json:"proxyOverride"`

}

func ImportJson(filepath string) (*prochSetting, error) {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var ps prochSetting
	if err := json.Unmarshal(bytes, &ps); err != nil {
		return nil, err
	}
	
	return &ps, nil
}