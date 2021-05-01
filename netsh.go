package proch

import (
	"os/exec"
	"fmt"
	"os"
	"strings"
	"syscall"

)

func execNetsh() string {
	cmd := exec.Command("cmd", "/C", "netsh", "wlan", "show", "profiles")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error: Failed to execute 'cmd /C netsh wlan show profile'\n\t%s\n\n", err)
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
		fmt.Printf("Error: Failed to execute 'cmd /C netsh wlan show interface'\n\t%s\n\n", err)
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
		fmt.Printf("Error: falied to execute 'cmd /C netsh wlan show networks'\n\t%s\n", err)
		return tmp
	}
	tmp = strings.Split(string(out), " : ")
	tmp = tmp[2:]
	net := []string{}
	for i := 0; i < len(tmp); i+=4 {
		splts := strings.Split(tmp[i], "\n")
		trimed := strings.Trim(splts[0], " ")
		net = append(net, trimed[:len(trimed)-1])
	}

	return net
}