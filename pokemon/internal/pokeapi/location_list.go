package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseUrl + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// using the cache here
	if val, ok := c.cache.Get(url); ok {
		locationResp := RespShallowLocations{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	c.cache.Add(url, dat)
	return locationsResp, nil
}

func (c *Client) GetLocationArea(name string) (LocationAreaResp, error) {
	// 1. pass the url
	url := baseUrl + "/location-area/" + name

	// 2. build the request
	res, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResp{}, err
	}

	//3. send it, wait for response
	resp, err := c.httpClient.Do(res)
	if err != nil {
		return LocationAreaResp{}, err
	}
	defer resp.Body.Close()
	// 4. read

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return LocationAreaResp{}, err
	}

	locationData := LocationAreaResp{}

	//5 unmarshall it
	err = json.Unmarshal(data, &locationData)

	if err != nil {
		return LocationAreaResp{}, err
	}

	// 6. return the response
	return locationData, nil
}

func (c *Client) CatchPokemonByName(name string) (Pokemon, bool, error) {
	// 1.url
	url := baseUrl + "/pokemon/" + name

	res, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, false, fmt.Errorf("error preparing request for pokemon:%w", err)
	}

	resp, err := c.httpClient.Do(res)
	if err != nil {
		return Pokemon{}, false, fmt.Errorf("error sending request %w", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, false, fmt.Errorf("error reading data: %w", err)
	}

	pokemon := Pokemon{}

	if err := json.Unmarshal(data, &pokemon); err != nil {
		return Pokemon{}, false, fmt.Errorf("Error unmarshaling response %w", err)
	}
	// 2.math random
	catchThreshold := 40
	roll := rand.Intn(pokemon.BaseExperience)

	caught := roll >= catchThreshold

	if !caught {
		return Pokemon{}, false, nil
	}
	return pokemon, true, nil
}
