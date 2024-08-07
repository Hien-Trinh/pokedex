package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (client *Client) GetPokemon(name string) (Pokemon, error) {
	fullURL := baseURL + "/pokemon/" + name

	cached_data, ok := client.cache.Get(fullURL)
	if ok {
		response := Pokemon{}
		err := json.Unmarshal(cached_data, &response)
		if err != nil {
			return Pokemon{}, err
		}

		return response, nil
	}

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Pokemon{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	response := Pokemon{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return Pokemon{}, err
	}

	client.cache.Add(fullURL, data)

	return response, nil
}
