package main

import (
	"net/http"
	"fmt"
	"io"
)

const pokeLocationAreaURL = "https://pokeapi.co/api/v2/location-area/"
var mapPage int

func commandMap() error {
	mapPage++
	//get http.Response
	resp, err := http.Get(pokeLocationAreaURL)
	if err != nil {
		return fmt.Errorf("error obtaining http.get response from %s: %w", pokeLocationAreaURL, err)
	}
	//convert resp to io.Reader type so we can Unmarshal
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unable to obtain body of http.response: %w", err)
	}
	if resp.StatusCode > 299 {
		return fmt.Errorf("bad status code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	return nil
	}

}
//struct to capture json data. fields may have nested structs to capture structured data, there is one anonymous struct 
//at the bottom due to single use and to keep it clean
type LocationArea struct {
	Id int `json:"id"`
	Name string `json:"name"`
	GameIndex int `json:"game_index"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	Location NameAndUrl `json:"location"`
	Names []NameStruct `json:"names"`
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type EncounterMethodRate struct {
	EncounterMethod EncounterMethod `json:"encounter_method"`
	VersionDetails []EncounterVersionDetails `json:"version_details"`
}


type EncounterMethod struct {
	NameAndUrl
}

type EncounterVersionDetails struct {
	Rate int `json:"rate"`
	Version NameAndUrl `json:"version"`
}

type NameAndUrl struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

type NameStruct struct {
	Name string `json:"name"`
	Language NameAndUrl `json:"language"`
}

type PokemonEncounter struct {
	Pokemon NameAndUrl `json:"pokemon"`
	VersionDetails [] PokemonEncountersVersionDetails `json:"version_details"`
}

type PokemonEncountersVersionDetails struct {
	Version NameAndUrl `json:"version"`
	MaxChance int `json:"max_chance"`
	EncounterDetails []struct{ //this is the only one i did an anonymous struct for. the rest are nested
		MinLevel int `json:"min_level"`
		MaxLevel int `json:"max_level"`
		ConditionValues []any `json:"condition_values"` //GEMINI said this should be []NameAndUrl apparently. it found it in another part of the docs
		Chance int `json:"chance"`
		Method NameAndUrl `json:"method"`
	} `json:"encounter_details"`
}