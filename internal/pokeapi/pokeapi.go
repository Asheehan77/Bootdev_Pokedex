package pokeapi

import (
	"net/http"
	"time"
	"encoding/json"
	"io"
	"github.com/Asheehan77/Bootdev_Pokedex/internal/pokecache"
	"errors"
)

const(
	apiUrl = "https://pokeapi.co/api/v2"
)

type Client struct {
	pcache		pokecache.Cache
	httpClient 	http.Client
	collection	[]Pokemon
}

func NewClient(timeout time.Duration,reaptime time.Duration) Client {
	return Client{
		pcache:	pokecache.NewCache(reaptime),
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
	data, exists := c.pcache.Get(locUrl)
	if exists == false{
		req,err := http.NewRequest("GET",locUrl,nil)
		if err != nil {
			return LocationList{},err
		}

		res,err := c.httpClient.Do(req)
		if err != nil {
			return LocationList{},err
		}
		defer res.Body.Close()

		data,err = io.ReadAll(res.Body)
		if err != nil {
			return LocationList{},err
		}
		c.pcache.Add(locUrl,data)
	}
	
	var locslist LocationList
	err := json.Unmarshal(data,&locslist)
	if err != nil {
		return LocationList{},err
	}

	return locslist,nil
}

type Location struct {
	ID        int `json:"id"`
	Name  string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c *Client)GetLocationInfo(locName *string) (Location, error){	
	locUrl := apiUrl+"/location-area/"+*locName
	
	data, exists := c.pcache.Get(locUrl)
	if exists == false{
		req,err := http.NewRequest("GET",locUrl,nil)
		if err != nil {
			return Location{},err
		}

		res,err := c.httpClient.Do(req)
		if err != nil {
			return Location{},err
		}
		defer res.Body.Close()

		data,err = io.ReadAll(res.Body)
		if err != nil {
			return Location{},err
		}
		c.pcache.Add(locUrl,data)
	}
	
	var loclist Location
	err := json.Unmarshal(data,&loclist)
	if err != nil {
		return Location{},err
	}

	return loclist,nil
}

type Pokemon struct {
	BaseExperience int `json:"base_experience"`
	Height                 int    `json:"height"`
	ID                     int    `json:"id"`
	Name          string `json:"name"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

func (c *Client)GetPokemonInfo(pokName *string) (Pokemon, error){	
	locUrl := apiUrl+"/pokemon/"+*pokName
	
	data, exists := c.pcache.Get(locUrl)
	if exists == false{
		req,err := http.NewRequest("GET",locUrl,nil)
		if err != nil {
			return Pokemon{},err
		}

		res,err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{},err
		}
		defer res.Body.Close()

		data,err = io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{},err
		}
		c.pcache.Add(locUrl,data)
	}
	
	var pokeinfo Pokemon
	err := json.Unmarshal(data,&pokeinfo)
	if err != nil {
		return Pokemon{},err
	}

	return pokeinfo,nil
}

func (c *Client) AddPokemon(pok Pokemon){
	c.collection = append(c.collection,pok)
}

func (c * Client)GetPokemon(pokname string)(Pokemon,error){
	for _,pok := range c.collection {
		if pok.Name == pokname{
			return pok,nil
		}
	}
	return Pokemon{},errors.New("You havent caught that pokemon")
}

func (c * Client)GetPokemonList()([]Pokemon){

	return c.collection
}