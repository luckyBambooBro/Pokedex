package main

func main() {
	cfg := config{
		startURL: "https://pokeapi.co/api/v2/location-area/",
		next: "",
		previous: "",
	}
	startRepl(&cfg)
}
