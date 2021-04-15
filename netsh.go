package proch

import (
	"os/exec"
	"fmt"
	"os"
	"strings"

)

func execNetsh() string {
	out, err := exec.Command("cmd", "/C", "netsh", "wlan", "show", "profiles").Output()
	if err != nil {
		fmt.Printf("Error: Failed to execute 'cmd /C netsh wlan show profile'\n\t%s\n\n", err)
		os.Exit(1)
	}
	return string(out)
}


func GetWlanProfiles() []string {
	// change encoding from Shift-JIS to UTF-8
	err := exec.Command("cmd", "/C", "chcp", "65001").Run()
	if err != nil {
		fmt.Printf("Error: Failed to change encoding to UTF-8 by executing 'chcp 65001'\n\t%s\n\n", err)
		os.Exit(1)
	}

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
	out, err := exec.Command("cmd", "/C", "netsh", "wlan", "show", "interface").Output()
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