package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (client *Client) GetLocation(name string) (Location, error) {
	fullURL := baseURL + "/location-area/" + name

	cached_data, ok := client.cache.Get(fullURL)
	if ok {
		response := Location{}
		err := json.Unmarshal(cached_data, &response)
		if err != nil {
			return Location{}, err
		}

		return response, nil
	}

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Location{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	response := Location{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return Location{}, err
	}

	client.cache.Add(fullURL, data)

	return response, nil
}
