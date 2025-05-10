package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	if resp.StatusCode > 299 {
		err := fmt.Errorf("failure with retrieval : %v", resp.Status)
		return Pokemon{}, err
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	retrievedPokemon := Pokemon{}
	err = json.Unmarshal(data, &retrievedPokemon)
	if err != nil {
		fmt.Println("error at unmarshal")
		return Pokemon{}, err
	}
	return retrievedPokemon, nil
}
