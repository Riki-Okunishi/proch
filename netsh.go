package proch

import (
	"os/exec"
	"fmt"
	"os"
	"strings"
	"syscall"

)

type netshRunnable interface {
	Connect(ssid string) error
	Disconnect() error
	ShowProfiles() []string
	ShowNetworks() []string
	GetCurrentSsid() string
}

type netshRunner struct {
	netshRunnable
}

var _ netshRunnable = &netshRunner{}

func NewNetshRunner() netshRunnable {
	return &netshRunner{}
}

func (nr *netshRunner) Connect(ssid string) error {
	netsh_connect := exec.Command("cmd", "/C", "netsh", "wlan", "connect", fmt.Sprintf("name=%s", ssid))
	netsh_connect.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := netsh_connect.Output()
	if err != nil {
		fmt.Printf("error: failed to connect (wlan=%s)\n\t%s\n", ssid, err)
		return err
	}
	fmt.Printf("Connect wlan %s\n\t%s\n\n", ssid, string(out))
	return nil
}

func (nr *netshRunner) Disconnect() error {
	netsh_disconnect := exec.Command("cmd", "/C", "netsh", "wlan", "disconnect")
	netsh_disconnect.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := netsh_disconnect.Output()
	if err != nil {
		fmt.Printf("error: faild to disconnect\n\t%s\n", err)
		return err
	}
	fmt.Printf("Disconnect wlan:\n\t%s\n\n", string(out))
	return nil
}

func (nr *netshRunner) ShowProfiles() []string {
	netsh_profiles := exec.Command("cmd", "/C", "netsh", "wlan", "show", "profiles")
	netsh_profiles.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := netsh_profiles.Output()
	if err != nil {
		fmt.Printf("error: failed to execute 'cmd /C netsh wlan show profile'\n\t%s\n\n", err)
		return []string{}
	}

	// Split profiles with ":" to get Wlan SSID
	ssid_list := strings.Split(string(out), ":")
	ssid_list = ssid_list[2:] // index 0 and 1 is not used

	// More split from each elements with "\n" because splitted element whose index is '0' describes SSID
	for i := 0; i < len(ssid_list); i++ {
		sp := strings.Split(ssid_list[i], "\n")
		ssid_list[i] = strings.Trim(sp[0], " ") // trim sapce " " 
	}

	return ssid_list
}

func (nr *netshRunner) ShowNetworks() []string {
	cmd := exec.Command("cmd", "/C", "netsh", "wlan", "show", "networks")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("error: falied to execute 'cmd /C netsh wlan show networks'\n\t%s\n", err)
		return []string{}
	}
	tmp := strings.Split(string(out), "\n")
	net := []string{}
	for _, l := range tmp {
		if strings.Index(l, "SSID") == 0 {
			ssid_line := strings.Split(l, ":")
			trimed_ssid := strings.Trim(ssid_line[1], " ")
			net = append(net, trimed_ssid[:len(trimed_ssid)-1])
		}
	}

	return net
}

func (nr *netshRunner) GetCurrentSsid() string {
	cmd := exec.Command("cmd", "/C", "netsh", "wlan", "show", "interface")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("error: failed to execute 'cmd /C netsh wlan show interface'\n\t%s\n\n", err)
		return ""
	}

	if_info := strings.Split(string(out), ":")
	if_info = strings.Split(if_info[12], "\n")
	cssid := strings.Trim(if_info[0], " ")
	cssid = cssid[:len(cssid)-1]
	
	return cssid
}

/**** Old version ****/

func execNetsh() string {
	cmd := exec.Command("cmd", "/C", "netsh", "wlan", "show", "profiles")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("error: failed to execute 'cmd /C netsh wlan show profile'\n\t%s\n\n", err)
		os.Exit(1)
	}
	return string(out)
}


func GetWlanProfiles() []string {
	// execute command 'netsh wlan show profiles'
	netsh_profiles := execNetsh()

	// Split profiles with ":" to get Wlan SSID
	ssid_list := strings.Split(netsh_profiles, ":")
	ssid_list = ssid_list[2:] // index 0 and 1 is not used

	// More split from each elements with "\n" because splitted element whose index is '0' describes SSID
	for i := 0; i < len(ssid_list); i++ {
		sp := strings.Split(ssid_list[i], "\n")
		ssid_list[i] = strings.Trim(sp[0], " ") // trim sapce " " 
	}

	return ssid_list
}

func GetCurrentSsid() string {
	cmd := exec.Command("cmd", "/C", "netsh", "wlan", "show", "interface")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("error: failed to execute 'cmd /C netsh wlan show interface'\n\t%s\n\n", err)
		return ""
	}

	if_info := strings.Split(string(out), ":")
	if_info = strings.Split(if_info[12], "\n")
	cssid := strings.Trim(if_info[0], " ")
	cssid = cssid[:len(cssid)-1]
	
	return cssid
}

func GetWlanNetworks() []string {
	tmp := []string{}
	cmd := exec.Command("cmd", "/C", "netsh", "wlan", "show", "networks")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("error: falied to execute 'cmd /C netsh wlan show networks'\n\t%s\n", err)
		return tmp
	}
	tmp = strings.Split(string(out), "\n")
	net := []string{}
	for _, l := range tmp {
		if strings.Index(l, "SSID") == 0 {
			ssid_line := strings.Split(l, ":")
			trimed_ssid := strings.Trim(ssid_line[1], " ")
			net = append(net, trimed_ssid[:len(trimed_ssid)-1])
		}
	}

	return net
}