package proch

import (
	"fmt"
)

type registryEditorMock struct {
	filepath string
	proxyEnable bool
	proxyServer string
	proxyOverride string
}

var _ registryEditable = &registryEditorMock{}

func (rem *registryEditorMock) GetSettingJsonPath() string {
	return rem.filepath
}

func (rem *registryEditorMock) EditProxySettings(proxyEnable bool, proxyServer string, proxyOverride string) error {
	if proxyEnable {
		if proxyServer != rem.proxyServer || proxyOverride != rem.proxyOverride {
			return fmt.Errorf("error [registryEditorMock]: cannot edit registry")
		}
	}
	return nil
}