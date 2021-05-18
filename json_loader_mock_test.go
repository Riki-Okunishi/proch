package proch

import (
	"fmt"
)

type jsonLoaderMock struct {
	filepath string
	profiles []wlanProfile
}

var _ jsonLoadable = &jsonLoaderMock{}

func (jlm *jsonLoaderMock) Load(filepath string) ([]wlanProfile, error) {
	if filepath != jlm.filepath {
		return []wlanProfile{}, fmt.Errorf("error [jsonLoaderMock]: cannot load the json file")
	}
	return jlm.profiles, nil
}