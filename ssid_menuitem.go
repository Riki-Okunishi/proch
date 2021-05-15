package proch

import (
	"fmt"
	
	"github.com/getlantern/systray"

)

type ssidMenuItem struct {
	systray.MenuItem
	ssid string
	proxyEanble bool
	proxyServer string
	proxyOverride string
	ssidCh chan string
	closeCh chan struct{}
}

func newSsidMenuItem(ssidCh chan string, ssid string, proxyEnable bool, proxyServer string, proxyOverride string) *ssidMenuItem {
	mi := systray.AddMenuItemCheckbox(ssid, fmt.Sprintf("Connect to %s", ssid), false)
	smi := &ssidMenuItem{MenuItem: *mi, ssid: ssid, proxyEanble: proxyEnable, proxyServer: proxyServer, proxyOverride: proxyOverride, ssidCh: ssidCh, closeCh: make(chan struct{})}
	return smi
}

func (smi *ssidMenuItem) WaitClick() {
	for {
	select {
	case <-smi.ClickedCh:
		fmt.Printf("%s WaitClick() called.\n", smi.ssid)
		smi.ssidCh <- smi.ssid
	case <-smi.closeCh:
		fmt.Printf("Close goroutine %s\n", smi.ssid)
		close(smi.ClickedCh)
		close(smi.closeCh)
		return
		}
	}
}
