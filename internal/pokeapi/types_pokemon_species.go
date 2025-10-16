package pokeapi

// PokemonSpecies -
type PokemonSpecies struct {
    CaptureRate       int               `json:"capture_rate"`
    GenderRate        int               `json:"gender_rate"`
    GrowthRate        struct {
        Name string `json:"name"`
    } `json:"growth_rate"`
    FlavorTextEntries []struct {
        FlavorText string `json:"flavor_text"`
        Language   struct {
            Name string `json:"name"`
        } `json:"language"`
        Version struct {
            Name string `json:"name"`
        } `json:"version"`
    } `json:"flavor_text_entries"`
}