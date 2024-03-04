package pokeapi

type areaList struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) ListAreas(pageURL *string) (areaList, error) {
	var resp areaList
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}
	dat, err := c.get(fullURL)
	if err != nil {
		return resp, err
	}
	resp, err = unmarshalTo[areaList](dat)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
