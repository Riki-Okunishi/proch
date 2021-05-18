package proch

import (
	"fmt"
	"os"
	"os/exec"
	"time"
	"syscall"
	"runtime"
	"testing"

)


func TestMain(m *testing.M) {
	switch runtime.GOOS {
	case "windows":
		fmt.Printf("running on Windows. continue executing proch.\n")
	default:
		fmt.Printf("running on OS not supported. exit proch.")
		os.Exit(1)
	}

	// change encoding from Shift-JIS to UTF-8
	chcp := exec.Command("cmd", "/C", "chcp", "65001")
	chcp.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := chcp.Run()
	if err != nil {
		fmt.Printf("Error: Failed to change encoding to UTF-8 by executing 'chcp 65001'\n\t%s\n\n", err)
		os.Exit(1)
	}

	m.Run()
}

func TestProchRun(t *testing.T) {

	nrm := &netshRunnerMock{connectSsid: "SSID2", disconnect: true, profiles: []string{"SSID1", "SSID2"}, networks: []string{"SSID1", "SSID2"}, cssid: "SSID1"}
	rem := &registryEditorMock{filepath: "./test/test.json", proxyEnable: false, proxyServer: "proxy.com:80", proxyOverride: "192.168.0.*;<local>"}
	profiles := []wlanProfile{
		{
			Ssid: "SSID1",
			ProxyEnable: true,
			ProxyServer: "proxy.com:80",
			ProxyOverride: "192.168.0.*;<local>",
		},
		{
			Ssid: "SSID2",
			ProxyEnable: false,
		},
	}
	jlm := &jsonLoaderMock{filepath: "./test/test.json", profiles: profiles}

	pc := New(nrm, rem, jlm)

	go func() {
		time.Sleep(2*time.Second) // Wait proch

		pc.ssidCh <- "SSID2"

		time.Sleep(1*time.Second) // Wait proch

		if pc.current.ssid != "SSID2" {
			t.Errorf("clicked \"SSID2\": expected=SSID2, result=%s", pc.current.ssid)
		}

		t.Logf("finish proch")
		pc.quit.ClickedCh <- struct{}{}
	}()

	pc.Run()

}