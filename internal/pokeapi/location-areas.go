package pokeapi

type locationAreasResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) ListLocationAreas(pageURL *string) (locationAreasResponse, error) {
	var resp locationAreasResponse
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}
	dat, err := c.get(fullURL)
	if err != nil {
		return resp, err
	}
	resp, err = unmarshalTo[locationAreasResponse](dat)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
