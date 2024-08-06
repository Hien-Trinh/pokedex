package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (client *Client) ListLocations(pageURL *string) (LocationAreasResponse, error) {
	fullURL := baseURL + "/location-area?offset=0&limit=20"
	if pageURL != nil {
		fullURL = *pageURL
	}

	cached_data, ok := client.cache.Get(fullURL)
	if ok {
		response := LocationAreasResponse{}
		err := json.Unmarshal(cached_data, &response)
		if err != nil {
			return LocationAreasResponse{}, err
		}

		return response, nil
	}

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return LocationAreasResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	response := LocationAreasResponse{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	client.cache.Add(fullURL, data)

	return response, nil
}
