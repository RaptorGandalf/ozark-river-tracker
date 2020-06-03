package usgs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	baseUrl = "https://waterservices.usgs.gov/nwis/iv/?format=JSON&sites=07064533&parameterCd=00065"
)

// TODO: For simplicity of parsing data, this should probably only get data for a single gauge
func GetData(sites, parameters []string) (TimeSeriesResponse, error) {
	var timeSeries TimeSeriesResponse

	url := buildUrl(sites, parameters)

	bytes, err := get(url)
	if err != nil {
		return timeSeries, err
	}

	err = json.Unmarshal(bytes, &timeSeries)

	return timeSeries, err
}

func buildUrl(sites, parameters []string) string {
	sitesString := strings.Join(sites, ",")
	parameterString := strings.Join(parameters, ",")

	return fmt.Sprintf("https://waterservices.usgs.gov/nwis/iv/?format=JSON&sites=%s&parameterCd=%s", sitesString, parameterString)
}

func get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
