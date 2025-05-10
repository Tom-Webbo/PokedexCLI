package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreLocations(location string) (RespLocationsExplore, error) {
	url := baseURL + "/location-area/" + location

	if val, ok := c.cache.Get(url); ok {
		locationExplore := RespLocationsExplore{}
		err := json.Unmarshal(val, &locationExplore)
		if err != nil {
			return RespLocationsExplore{}, err
		}
		return locationExplore, nil

	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocationsExplore{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocationsExplore{}, err
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationsExplore{}, err
	}
	locationExplore := RespLocationsExplore{}
	err = json.Unmarshal(data, &locationExplore)
	if err != nil {
		return RespLocationsExplore{}, err
	}
	c.cache.Add(url, data)
	return locationExplore, nil
}
