package internal

import (
	"net/http"
	"time"
	"encoding/json"
	"io"
)

const(
	apiUrl = "https://pokeapi.co/api/v2"
)

type Client struct {
	httpClient http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

type LocationList struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client)GetLocations(reqUrl *string) (LocationList, error){	
	var locUrl string
	if reqUrl == nil{
		locUrl = apiUrl+"/location-area"
	}else{
		locUrl = *reqUrl
	}
	req,err := http.NewRequest("GET",locUrl,nil)
	if err != nil {
		return LocationList{},err
	}

	res,err := c.httpClient.Do(req)
	if err != nil {
		return LocationList{},err
	}
	defer res.Body.Close()

	data,err := io.ReadAll(res.Body)
	if err != nil {
		return LocationList{},err
	}

	var locslist LocationList
	if err := json.Unmarshal(data,&locslist); err != nil {
		return LocationList{},err
	}

	return locslist,nil
}