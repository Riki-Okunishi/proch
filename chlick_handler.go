package proch

import (
	"fmt"
	"os/exec"
	"syscall"
	
	"golang.org/x/sys/windows/registry"

	"github.com/getlantern/systray"
)
type clickEvent struct {
	systray.MenuItem
	WlanProfile wlanProfile
	eventCh chan string
	CloseCh chan struct{}
}

func (ce *clickEvent) WaitClick() {
	for {
		select {
		case <-ce.ClickedCh:
		fmt.Printf("\t%s WaitClick() called!\n", ce.WlanProfile.Ssid)
		ce.eventCh <- ce.WlanProfile.Ssid
		case <-ce.CloseCh:
			fmt.Printf("Close goroutine %s\n", ce.WlanProfile.Ssid)
			close(ce.ClickedCh)
			close(ce.CloseCh)
			return
		}
	}
}

func (ce *clickEvent) Connect() error {
	/* execute netsh and reg*/

	// disconnect current wlan
	netsh_disconnect := exec.Command("C:\\Windows\\system32\\netsh.exe", "wlan", "disconnect")
	netsh_disconnect.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := netsh_disconnect.Output()
	if err != nil {
		fmt.Printf("Error: faild to disconnect\n\t%s\n", err)
		return err
	}
	fmt.Printf("Disconnect wlan:\n\t%s\n\n", string(out))

	// Connect wlan
	netsh_connect := exec.Command("C:\\Windows\\system32\\netsh.exe", "wlan", "connect", fmt.Sprintf("name=%s", ce.WlanProfile.Ssid))
	netsh_connect.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err = netsh_connect.Output()
	if err != nil {
		fmt.Printf("Error: failed to connect (wlan=%s)\n\t%s\n", ce.WlanProfile.Ssid, err)
		return err
	}
	fmt.Printf("Connect wlan %s\n\t%s\n\n", ce.WlanProfile.Ssid, string(out))

	// Setting Proxy by Editing Regidtry
	var dword uint32
	if ce.WlanProfile.ProxyEnable {
		dword = 1
	} else {
		dword = 0
	}

	k, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.WRITE)
	if err != nil {
		fmt.Printf("Error: registry.OpenKey(...)\n\t%s\n\n", err)
	}
	defer k.Close()

	// Edit value of ProxyEnable
	err = k.SetDWordValue("ProxyEnable", dword)
	if err != nil {
		fmt.Printf("Error: Key.SetDWordValue(...)\n\t%s\n\n", err)
		return err
	}
	
	if dword == 0 {
		return nil
	}

	// if use proxy, edit value of 'ProxyServer' and 'ProxyOverride'
	err = k.SetStringValue("ProxyServer", ce.WlanProfile.ProxyServer)
	if err != nil {
		fmt.Printf("Error: failed to set ProxyServer value with Key.SetStringValue(...)\n\t%s\n\n", err)
		return err
	}
	err = k.SetStringValue("ProxyOverride", ce.WlanProfile.ProxyOverride)
	if err != nil {
		fmt.Printf("Error: failed to set ProxyOverride value with Key.SetStringValue(...)\n\t%s\n\n", err)
		return err
	}

	return nil
}



type clickHandler struct {
	eventCh chan string
	eventList map[string]*clickEvent
}

func NewClickHandler() *clickHandler {
	ch := &clickHandler{eventCh: make(chan string, 1), eventList: map[string]*clickEvent{}}
	return ch
}

func (ch *clickHandler) AddEvent(wp wlanProfile) {
	mi := systray.AddMenuItemCheckbox(wp.Ssid, "", false)
	ce := &clickEvent{MenuItem: *mi, WlanProfile: wp, eventCh: ch.eventCh, CloseCh: make(chan struct{})}
	ch.eventList[wp.Ssid] = ce
}


func (ch *clickHandler) HandleClick() {
	//current network must be checked
	cssid := GetCurrentSsid()
	fmt.Printf("Current SSID: '%s'\n", cssid)
	if ec, ok := ch.eventList[cssid]; ok {
		ec.Check()
		ch.setTooltop(ec.WlanProfile.ProxyEnable)
	} else {
		ch.setTooltop(false)
	}

	// exec goroutine each Buttom
	for _, e := range ch.eventList {
		fmt.Printf("Goroutine for %s\n", e.WlanProfile.Ssid)
		go e.WaitClick()
	}

	// wait a channel input from a buttom
	fmt.Printf("Start proch!\n")
	for {
		select {
		case ssid := <-ch.eventCh:
			fmt.Printf("\tclicked %s!\n", ssid)
			// check clicked SSID
			e, ok := ch.eventList[ssid]
			if !ok {
				fmt.Printf("Error: not found such event represented as '%s'", ssid)
				continue
			}
			if !e.Checked() {
				// try connect
				if err := e.Connect(); err != nil {
					fmt.Printf("Error: failed to connect\n\t%s\n", err)
					continue
				}
				e.Check()

				// uncheck other menu
				for _, e := range ch.eventList {
					if e.WlanProfile.Ssid == ssid {
						continue
					}
					e.Uncheck()
				}

				// change Tooltip
				ch.setTooltop(e.WlanProfile.ProxyEnable)
			}
		}
	}
}

func (ch *clickHandler) CloseAllCh() {
	fmt.Printf("Close All Channel\n")
	for _, e := range ch.eventList {
		e.CloseCh <- struct{}{}
	}
}

func (ch *clickHandler) setTooltop(proxyEnable bool) {
	var str string
	if proxyEnable {
		str = "Enable"
	} else {
		str = "Disable"
	}
	systray.SetTooltip(fmt.Sprintf("Proxy Changer\nProxy: %s", str))
}
