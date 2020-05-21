package usgs

import (
	"io/ioutil"
	"net/http"
)

func GetSites() ([]byte, error) {

	resp, err := http.Get("https://waterservices.usgs.gov/nwis/iv/?format=JSON&sites=07064533&parameterCd=00065")
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}
