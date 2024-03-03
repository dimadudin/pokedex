package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type locationAreasResponse struct {
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
	} `json:"results"`
}

func (c *Client) ListLocationAreas(pageURL *string) (locationAreasResponse, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	dat, ok := c.pokeCache.Get(fullURL)
	if ok {
		locAreasResponse := locationAreasResponse{}
		err := json.Unmarshal(dat, &locAreasResponse)
		if err != nil {
			return locationAreasResponse{}, err
		}
		return locAreasResponse, nil
	}

	resp, err := http.Get(fullURL)
	if err != nil {
		return locationAreasResponse{}, err
	}
	if resp.StatusCode > 399 {
		return locationAreasResponse{}, fmt.Errorf("response failed with status code %d", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return locationAreasResponse{}, err
	}

	locAreasResponse := locationAreasResponse{}
	err = json.Unmarshal(dat, &locAreasResponse)
	if err != nil {
		return locationAreasResponse{}, err
	}

	c.pokeCache.Add(fullURL, dat)

	return locAreasResponse, nil
}
