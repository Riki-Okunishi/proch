package proch

import (
	"fmt"
	
	"github.com/getlantern/systray"

)

type ssidMenuItem struct {
	*systray.MenuItem
	ssid string
	proxyEnable bool
	proxyServer string
	proxyOverride string
	ssidCh chan string
	closeCh chan struct{}
}

func newSsidMenuItem(mi *systray.MenuItem, ssidCh chan string, ssid string, proxyEnable bool, proxyServer string, proxyOverride string) *ssidMenuItem {
	// TODO: Add validation for mi, ssid, proxyServer and proxyOverride (when proxyEnable == true)
	smi := &ssidMenuItem{MenuItem: mi, ssid: ssid, proxyEnable: proxyEnable, proxyServer: proxyServer, proxyOverride: proxyOverride, ssidCh: ssidCh, closeCh: make(chan struct{})}
	return smi
}

func (smi *ssidMenuItem) waitClick() {
	for {
	select {
	case <-smi.ClickedCh:
		fmt.Printf("%s waitClick() called.\n", smi.ssid)
		smi.ssidCh <- smi.ssid
	case <-smi.closeCh:
		fmt.Printf("Close goroutine %s\n", smi.ssid)
		close(smi.ClickedCh)
		close(smi.closeCh)
		return
		}
	}
}
