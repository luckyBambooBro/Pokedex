package main

func main() {
	config := configStruct{
		startURL: "https://pokeapi.co/api/v2/location-area/",
		next: "",
		previous: "",
	}
	startRepl(&config)
}
