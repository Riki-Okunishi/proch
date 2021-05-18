package proch

import (
	"fmt"
	"sync"

	"github.com/getlantern/systray"
)

type proxyChanger struct {
	ssidCh chan string
	current *ssidMenuItem
	ssidList map[string]*ssidMenuItem
	refresh *systray.MenuItem
	quit *systray.MenuItem
	netsh netshRunnable
	registry registryEditable
	json jsonLoadable
}


/** New version **/

func New(n netshRunnable, r registryEditable, j jsonLoadable) *proxyChanger {
	p := proxyChanger{ssidCh: make(chan string, 1), current: nil, ssidList: map[string]*ssidMenuItem{}, refresh: nil, quit: nil, netsh: n, registry: r, json: j}
	return &p
}

func (pc *proxyChanger) Run() {
	systray.Run(pc.onReady, pc.onExit)
}

func (pc *proxyChanger) onExit() {
	// now := time.Now()
	// ioutil.WriteFile(fmt.Sprintf(`on_exit_%d.txt`, now.UnixNano()), []byte(now.String()), 0644)
}

func (pc *proxyChanger) onReady() {
	// Here process must be done before adding MenuItem
	systray.SetTemplateIcon(iconData, iconData)
	systray.SetTitle("Proxy Changer")

	filepath := pc.registry.GetSettingJsonPath()
	profiles, err := pc.json.Load(filepath)
	if err != nil {
		fmt.Printf("error: failed to open file %s\n\t%s\n", filepath, err)
		return
	}
	var proxy_ssid []wlanProfile
	var non_proxy_ssid []wlanProfile
	for _, wp := range profiles {
		if wp.ProxyEnable {
			proxy_ssid = append(proxy_ssid, wp)
		} else {
			non_proxy_ssid = append(non_proxy_ssid, wp)
		}
	}

	for _, wp := range non_proxy_ssid {
		smi:= newSsidMenuItem(nil, pc.ssidCh, wp.Ssid, wp.ProxyEnable, wp.ProxyServer, wp.ProxyOverride)
		if smi != nil {
			mi := systray.AddMenuItemCheckbox(wp.Ssid, fmt.Sprintf("Connect to %s", wp.Ssid), false)
			smi.MenuItem = mi
			pc.ssidList[smi.ssid] = smi
		}
	}
	systray.AddSeparator()
	for _, wp := range proxy_ssid {
		smi:= newSsidMenuItem(nil, pc.ssidCh, wp.Ssid, wp.ProxyEnable, wp.ProxyServer, wp.ProxyOverride)
		if smi != nil {
			mi := systray.AddMenuItemCheckbox(wp.Ssid, fmt.Sprintf("Connect to %s", wp.Ssid), false)
			smi.MenuItem = mi
			pc.ssidList[smi.ssid] = smi
		}
	}
	systray.AddSeparator()
	pc.refresh = systray.AddMenuItem("Refresh", "Refresh the SSID list")
	systray.AddSeparator()
	pc.quit = systray.AddMenuItem("Quit", "Quit the whole app")

	go pc.HandleClick()
}

func (pc *proxyChanger) HandleClick() {
	pc.refreshCurrentSsid()

	wg := &sync.WaitGroup{}
	for _, smi := range pc.ssidList {
		wg.Add(1)
		fmt.Printf("Goroutine for %s\n", smi.ssid)
		go func(smi *ssidMenuItem) {
			smi.waitClick()
			wg.Done()
		}(smi)
	}

	fmt.Printf("Start proch!\n")
	for {
		select {
		case ssid := <-pc.ssidCh:
			fmt.Printf("clicked %s\n", ssid)
			// check clicked SSID
			smi, ok := pc.ssidList[ssid]
			if !ok {
				fmt.Printf("error: not found such event represented as '%s'\n", ssid)
				continue
			}

			// Check this ssid exists around here
			net := pc.netsh.ShowNetworks()
			find := false
			for _, s := range net {
				if s == smi.ssid {
					find = true
					break
				}
			}
			if !find {
				fmt.Printf("error: not found the network '%s' around here\n", smi.ssid)
				continue
			}

			if !smi.Checked() {

				//disconnect from previous network WARN: Why does this process required?
				if pc.current != nil {
					if err := pc.netsh.Disconnect(); err != nil {
						fmt.Printf("Error: failed to disconnect from %s\n\t%s\n", pc.current.ssid, err)
						continue
					}
					// uncheck previous menu
					pc.current.Uncheck()
					setTooltip(false)
					pc.current = nil
				}

				// try connect
				if err := pc.netsh.Connect(smi.ssid); err != nil {
					fmt.Printf("error: failed to connect\n\t%s\n", err)
					continue
				}
				if err := pc.registry.EditProxySettings(smi.proxyEnable, smi.proxyServer, smi.proxyOverride); err != nil {
					fmt.Printf("error: failed to edit the registries\n\t%s\n", err)
					continue
				}
				smi.Check()
				setTooltip(smi.proxyEnable)
				pc.current = smi
			} else {

				// Dissconnect current SSID
				if pc.current == nil {
					fmt.Printf("Error: expected ch.current is not nil when checked menu item was clicked\n")
					continue
				}

				if err := pc.netsh.Disconnect(); err != nil {
					fmt.Printf("Error: failed to disconnect\n\t%s\n", err)
					continue
				}

				pc.current.Uncheck()
				setTooltip(false)
				pc.current = nil
			}
		case <-pc.refresh.ClickedCh:
			fmt.Printf("Clicked Refresh\n")
			pc.refreshCurrentSsid()
		case <-pc.quit.ClickedCh:		
			fmt.Println("Requesting quit")
			pc.CloseAllCh()
			wg.Wait()
			systray.Quit()
			fmt.Println("Finished quitting")
			return
		}
	}
}

func (pc *proxyChanger) refreshCurrentSsid() {
	cssid := pc.netsh.GetCurrentSsid()
	fmt.Printf("Current SSID: '%s'\n", cssid)
	for _, smi := range pc.ssidList {
		smi.Uncheck()
	}
	if smi, ok := pc.ssidList[cssid]; ok {
		err := pc.registry.EditProxySettings(smi.proxyEnable, smi.proxyServer, smi.proxyOverride)
		if err != nil {
			fmt.Printf("error: failed to refresh proxy setting in this network '%s'\n\t%s\n", cssid, err)
		} else {
			fmt.Printf("refreshed proxy settings: %t\n", smi.proxyEnable)
		}
		smi.Check()
		setTooltip(smi.proxyEnable)
		pc.current = smi
	} else {
		setTooltip(false)
		pc.current = nil
	}
}

func (pc *proxyChanger) CloseAllCh() {
	fmt.Printf("Close All Channel\n")
	for _, s := range pc.ssidList {
		s.closeCh <- struct{}{}
	}
}

// setTooltip will set the Tooltip text whether the proxy settings is enable or disable
func setTooltip(proxyEnable bool) {
	var str string
	if proxyEnable {
		str = "Enable"
	} else {
		str = "Disable"
	}
	systray.SetTooltip(fmt.Sprintf("Proxy Changer\nProxy: %s", str))
}

