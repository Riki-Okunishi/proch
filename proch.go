package proch

import (
	"fmt"
	"os"

	"golang.org/x/sys/windows/registry"

	"github.com/getlantern/systray"
)

const (
	registryKey = `SOFTWARE\Proch`
	defaultPath = "./setting.json"
)

type proxyChanger struct {
	json string
}

func NewProxyChanger() *proxyChanger {
	pc := &proxyChanger{}

	// Initialize if not exist setting.json
	jsonexist := true
	if !jsonexist {
		pc.createSettingJson()
	}

	return pc
}

func (pc *proxyChanger) createSettingJson() {
	// 1. Exec netsh to get profiles

	// 2. Extract SSID and Proxy info

	// 3. Export settings as JSON
}

// getSettingJsonPath will return "setting.json" file path from registry. 
// If not exist registry key "HKLM\SOFTWARE\Proch\SettingJson" or it's blank, will return defaultPath "./setting.json".
func getSettingJsonPath() string {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, registryKey, registry.QUERY_VALUE)
	if err != nil {
		return defaultPath
	}
	defer k.Close()

	s, _, err := k.GetStringValue("SettingJson")
	if err != nil || s == "" {
		return defaultPath
	}

	return s
}

func (pc *proxyChanger) Run() {
	systray.Run(pc.onReady, pc.onExit)
}

func (pc *proxyChanger) onExit() {
	// now := time.Now()
	// ioutil.WriteFile(fmt.Sprintf(`on_exit_%d.txt`, now.UnixNano()), []byte(now.String()), 0644)
}

func (pc *proxyChanger) onReady() {
	// Load from setting.json to create SSID list
	pc.json = getSettingJsonPath()

	ps, err := ImportJson(pc.json)
	if err != nil {
		fmt.Printf("Error: failed to open setting.json\n")
		os.Exit(1)
	}
	
	// application setting
	systray.SetTemplateIcon(iconData, iconData)
	systray.SetTitle("Proxy Changer")
	//systray.SetTooltip("Proxy Changer")

	/* Add MenuItem representing wlan profiles */
	// Registry to clickHandler
	ch := NewClickHandler()

	// separate proxy and non-proxy profiles
	var proxy []wlanProfile
	var non_proxy []wlanProfile
	for _, v := range ps.Profiles {
		if v.ProxyEnable {
			proxy = append(proxy, v)
		}else{
			non_proxy = append(non_proxy, v)
		}
	}

	// add non-Proxy network
	// systray.AddMenuItem("non-Proxy Network", "network not required proxy")
	for _, v := range non_proxy {
		ch.AddEvent(v)
	}

	systray.AddSeparator()
	
	// add Porxy network
	// systray.AddMenuItem("Proxy Network", "network required proxy")
	for _, v := range proxy {
		ch.AddEvent(v)
	}


	// Exec systray malutipulating in two goroutine
	// clickHandler goroutine
	go ch.HandleClick() // v2: AddMenu after goroutine
	
	// Refresh Menu
	systray.AddSeparator()
	ch.refresh = systray.AddMenuItem("Refresh", "Refresh the SSID list")

	// Quit Menu	
	systray.AddSeparator()
	mQuitOrig := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-mQuitOrig.ClickedCh
		fmt.Println("Requesting quit")
		ch.CloseAllCh()
		systray.Quit()
		fmt.Println("Finished quitting")
	}()



}