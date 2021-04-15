package proch

import (
	"fmt"
	"os"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

type proxyChanger struct {

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

func (pc *proxyChanger) Run() {
	systray.Run(pc.onReady, pc.onExit)
}

func (pc *proxyChanger) onExit() {
	// now := time.Now()
	// ioutil.WriteFile(fmt.Sprintf(`on_exit_%d.txt`, now.UnixNano()), []byte(now.String()), 0644)
}

func (pc *proxyChanger) onReady() {
	// Load from setting.json to create SSID list
	ps, err := ImportJson("./setting.json")
	if err != nil {
		fmt.Printf("Error: failed to open setting.json\n")
		os.Exit(1)
	}
	
	// application setting
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("Proxy Changer")
	systray.SetTooltip("Select Wlan Prifile")

	/* Add MenuItem representing wlan profiles */
	// Registry to clickHandler
	ch := NewClickHandler()
	for _, v := range ps.Profiles { // v1: addMenu before goroutine
		ch.AddEvent(v)
	}


	// Exec systray malutipulating in two goroutine
	// clickHandler goroutine
	go ch.HandleClick() // v2: AddMenu after goroutine
	

	// Quit Menu	
	systray.AddSeparator()
	mQuitOrig := systray.AddMenuItem("終了", "Quit the whole app")
	go func() {
		<-mQuitOrig.ClickedCh
		fmt.Println("Requesting quit")
		ch.CloseAllCh()
		systray.Quit()
		fmt.Println("Finished quitting")
	}()



}