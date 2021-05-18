package proch

import (
	"fmt"
	"golang.org/x/sys/windows/registry"

)

const (
	registryKey = `SOFTWARE\Proch`
	defaultPath = "./setting.json"
)

type registryEditable interface {
	GetSettingJsonPath() string
	EditProxySettings(proxyEnable bool, proxyServer string, proxyOverride string) error
}

type registryEditor struct {
	registryEditable
}

var _ registryEditable = &registryEditor{}

func NewRegistryEditor() registryEditable {
	return &registryEditor{}
}

func (re *registryEditor) GetSettingJsonPath() string {
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

func (re *registryEditor) EditProxySettings(proxyEnable bool, proxyServer string, proxyOverride string) error {
	// Setting Proxy by Editing Regidtry
	var dword uint32
	if proxyEnable {
		dword = 1
	} else {
		dword = 0
	}

	k, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Internet Settings`, registry.WRITE)
	if err != nil {
		fmt.Printf("error: registry.OpenKey(...)\n\t%s\n\n", err)
		return err
	}
	defer k.Close()

	// Edit value of ProxyEnable
	err = k.SetDWordValue("ProxyEnable", dword)
	if err != nil {
		fmt.Printf("error: Key.SetDWordValue(...)\n\t%s\n\n", err)
		return err
	}

	if dword == 0 {
		return nil
	}

	// if use proxy, edit value of 'ProxyServer' and 'ProxyOverride'
	err = k.SetStringValue("ProxyServer", proxyServer)
	if err != nil {
		fmt.Printf("error: failed to set ProxyServer value with Key.SetStringValue(...)\n\t%s\n\n", err)
		return err
	}
	err = k.SetStringValue("ProxyOverride", proxyOverride)
	if err != nil {
		fmt.Printf("error: failed to set ProxyOverride value with Key.SetStringValue(...)\n\t%s\n\n", err)
		return err
	}

	return nil
}