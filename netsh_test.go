package proch

import (
	"testing"
)

func TestGetWlanNetworks(t *testing.T){
	net := GetWlanNetworks()
	for i, ssid := range net {
		t.Logf("net[%d]= '%s'", i, ssid)
	}

}