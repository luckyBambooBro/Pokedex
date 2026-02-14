package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


func commandMap(config *configStruct) error {
	//set starting URL
	url := ""
	if config.next == "" {
		url = config.startURL
	} else {
		url = config.next 
	}

	//get http.Response from pokeAPI
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error obtaining http.get response from %s: %w", url, err)
	}
	//convert resp to io.Reader type so we can Unmarshal
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unable to obtain body of http.response: %w", err)
	}
	if resp.StatusCode > 299 {
		return fmt.Errorf("bad status code: %d", resp.StatusCode)
	}
	//safeguard close. can put another close call later if wanted
	defer resp.Body.Close()

	var apiResponse apiStruct
	if err = json.Unmarshal(data, &apiResponse); err != nil {
		return fmt.Errorf("Error unmarshaling http response body: %w", err)
	}

	//print api response to terminal
	for _, item := range apiResponse.Results {
		fmt.Println(item.Name)
	}
	
	//update url to increment map page for next function call to call the next page of 20 results
	config.previous = apiResponse.Previous
	config.next = apiResponse.Next
	return nil
}

func commandMapB(config *configStruct) error {
	//set starting URL
	url := ""
	if config.previous == "" {
		fmt.Println("you're on the first page")
		return nil		
	} else {
		url = config.previous 
	}

	//get http.Response from pokeAPI
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error obtaining http.get response from %s: %w", url, err)
	}
	//convert resp to io.Reader type so we can Unmarshal
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unable to obtain body of http.response: %w", err)
	}
	if resp.StatusCode > 299 {
		return fmt.Errorf("bad status code: %d", resp.StatusCode)
	}
	//safeguard close. can put another close call later if wanted
	defer resp.Body.Close()

	var apiResponse apiStruct
	if err = json.Unmarshal(data, &apiResponse); err != nil {
		return fmt.Errorf("Error unmarshaling http response body: %w", err)
	}

	//print api response to terminal
	for _, item := range apiResponse.Results {
		fmt.Println(item.Name)
	}
	
	//update url to increment map page for next function call to call the next page of 20 results
	config.previous = apiResponse.Previous
	config.next = apiResponse.Next
	return nil
}
// struct to capture apiStruct via http.Get()
type apiStruct struct {
	Count    int          `json:"count"`
	Next     string       `json:"next"`
	Previous string       `json:"previous"`
	Results  []NameAndUrl `json:"results"`
}

// struct to capture json data. fields may have nested structs to capture structured data, there is one anonymous struct
// at the bottom due to single use and to keep it clean
type LocationArea struct {
	Id                   int                   `json:"id"`
	Name                 string                `json:"name"`
	GameIndex            int                   `json:"game_index"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	Location             NameAndUrl            `json:"location"`
	Names                []NameStruct          `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

type EncounterMethodRate struct {
	EncounterMethod EncounterMethod           `json:"encounter_method"`
	VersionDetails  []EncounterVersionDetails `json:"version_details"`
}

type EncounterMethod struct {
	NameAndUrl
}

type EncounterVersionDetails struct {
	Rate    int        `json:"rate"`
	Version NameAndUrl `json:"version"`
}

type NameAndUrl struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type NameStruct struct {
	Name     string     `json:"name"`
	Language NameAndUrl `json:"language"`
}

type PokemonEncounter struct {
	Pokemon        NameAndUrl                        `json:"pokemon"`
	VersionDetails []PokemonEncountersVersionDetails `json:"version_details"`
}

type PokemonEncountersVersionDetails struct {
	Version          NameAndUrl `json:"version"`
	MaxChance        int        `json:"max_chance"`
	EncounterDetails []struct { //this is the only one i did an anonymous struct for. the rest are nested
		MinLevel        int        `json:"min_level"`
		MaxLevel        int        `json:"max_level"`
		ConditionValues []any      `json:"condition_values"` //GEMINI said this should be []NameAndUrl apparently. it found it in another part of the docs
		Chance          int        `json:"chance"`
		Method          NameAndUrl `json:"method"`
	} `json:"encounter_details"`
}
