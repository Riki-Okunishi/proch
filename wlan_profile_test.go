package proch_test

import (
	"testing"

	"github.com/Riki-Okunishi/proch"
)

var (
	filepath = "./test/test.json"
)

func TestImportJson(t *testing.T) {
	wp, err := proch.ImportJson(filepath)
	if err != nil {
		t.Errorf("Failed to import JSON file '%s'\n", filepath)
	}

	correct_ssid := "KE101-Proxy(2.4GHz)"
	if wp.Profiles[0].Ssid != correct_ssid {
		t.Errorf("Failed to convert from json to struct: first profile's SSID must be '%s' but %s\n", correct_ssid, wp.Profiles[0].Ssid)
	}


	t.Logf("JSON:\n%v\n", wp)
}