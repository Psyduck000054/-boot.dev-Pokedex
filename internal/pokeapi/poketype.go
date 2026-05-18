package pokeapi

type Pokemon struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PokeBall struct {
	Container Pokemon `json:"pokemon"`
}

type AreaEncounterInfo struct {
	Results []PokeBall `json:"pokemon_encounters"`
}
