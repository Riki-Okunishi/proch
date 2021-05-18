package proch

import (
	"testing"

	"github.com/getlantern/systray"
)

type testCase struct {
	mi *systray.MenuItem
	ssidCh chan string
	ssid string
	proxyEnable bool
	proxyServer string
	proxyOverride string
	expected interface{}
}

func TestNewSsidMenuItemWithSsid(t *testing.T) {
	
	mi := &systray.MenuItem{}
	testCh := make(chan string, 1)

	ssidTest := []testCase{
		{
			mi: mi,
			ssidCh: testCh,
			ssid: "SSID1",
			proxyEnable: true,
			proxyServer: "proxy.com:80",
			proxyOverride: "192.168.0.*;<local>",
			expected: "SSID1",
		},
	}


	for i, tc := range ssidTest {
		smi := newSsidMenuItem(tc.mi, tc.ssidCh, tc.ssid, tc.proxyEnable, tc.proxyServer, tc.proxyOverride)

		if smi.ssid != tc.expected {
			t.Errorf("in ssidTest[%d]: expected=%s but result=%s", i, tc.expected, smi.ssid)
		}
	}
}

func TestNewSsidMenuItemWithProxyEnable(t *testing.T) {
	
	testCh := make(chan string, 1)

	proxyEnableTest := []testCase{
		{
			mi: &systray.MenuItem{},
			ssidCh: testCh,
			ssid: "SSID1",
			proxyEnable: true,
			proxyServer: "proxy.com:80",
			proxyOverride: "192.168.0.*;<local>",
			expected: true,
		},
		{
			mi: &systray.MenuItem{},
			ssidCh: testCh,
			ssid: "SSID2",
			proxyEnable: false,
			proxyServer: "",
			proxyOverride: "",
			expected: false,
		},
	}


	for i, tc := range proxyEnableTest {
		smi := newSsidMenuItem(tc.mi, testCh, tc.ssid, tc.proxyEnable, tc.proxyServer, tc.proxyOverride)

		if smi.proxyEnable != tc.expected {
			t.Errorf("in ssidTest[%d]: expected=%v but result=%v", i, tc.expected, smi.proxyEnable)
		}
	}
}

func TestNewSsidMenuItemWithProxyServer(t *testing.T) {
	
	testCh := make(chan string, 1)

	proxyServerTest := []testCase{
		{
			mi: &systray.MenuItem{},
			ssidCh: testCh,
			ssid: "SSID1",
			proxyEnable: true,
			proxyServer: "proxy.com:80",
			proxyOverride: "192.168.0.*;<local>",
			expected: "proxy.com:80",
		},
		{
			mi: &systray.MenuItem{},
			ssidCh: testCh,
			ssid: "SSID2",
			proxyEnable: false,
			proxyServer: "",
			proxyOverride: "",
			expected: "",
		},
		{
			mi: &systray.MenuItem{},
			ssidCh: testCh,
			ssid: "SSID3",
			proxyEnable: true,
			proxyServer: "proxy.com",
			proxyOverride: "192.168.0.*;<local>",
			expected: nil,
		},
		{
			mi: &systray.MenuItem{},
			ssidCh: testCh,
			ssid: "SSID4",
			proxyEnable: true,
			proxyServer: ":80",
			proxyOverride: "192.168.0.*;<local>",
			expected: nil,
		},
	}


	for i, tc := range proxyServerTest {
		smi := newSsidMenuItem(tc.mi, testCh, tc.ssid, tc.proxyEnable, tc.proxyServer, tc.proxyOverride)

		if smi.proxyServer != tc.expected {
			t.Errorf("in ssidTest[%d]: expected=%v but result=%v", i, tc.expected, smi.proxyServer)
		}
	}
}

func TestNewSsidMenuItemWithProxyOverride(t *testing.T) {
	
	testCh := make(chan string, 1)

	proxyOverrideTest := []testCase{
		{
			mi: &systray.MenuItem{},
			ssidCh: testCh,
			ssid: "SSID1",
			proxyEnable: true,
			proxyServer: "proxy.com:80",
			proxyOverride: "192.168.0.*;<local>",
			expected: "192.168.0.*;<local>",
		},
		{
			mi: &systray.MenuItem{},
			ssidCh: testCh,
			ssid: "SSID2",
			proxyEnable: true,
			proxyServer: "proxy.com:80",
			proxyOverride: "192.168.0.1",
			expected: "192.168.0.1",
		},
		{
			mi: &systray.MenuItem{},
			ssidCh: testCh,
			ssid: "SSID3",
			proxyEnable: true,
			proxyServer: "proxy.com:80",
			proxyOverride: "<local>",
			expected: "<local>",
		},
		{
			mi: &systray.MenuItem{},
			ssidCh: testCh,
			ssid: "SSID4",
			proxyEnable: true,
			proxyServer: "proxy.com:80",
			proxyOverride: "192.168.0.256;<local>",
			expected: nil,
		},
		{
			mi: &systray.MenuItem{},
			ssidCh: testCh,
			ssid: "SSID5",
			proxyEnable: true,
			proxyServer: "proxy.com:80",
			proxyOverride: "192.168.*;<local>",
			expected: nil,
		},
	}


	for i, tc := range proxyOverrideTest {
		smi := newSsidMenuItem(tc.mi, testCh, tc.ssid, tc.proxyEnable, tc.proxyServer, tc.proxyOverride)

		if smi.proxyOverride != tc.expected {
			t.Errorf("in ssidTest[%d]: expected=%v but result=%v", i, tc.expected, smi.proxyOverride)
		}
	}
}