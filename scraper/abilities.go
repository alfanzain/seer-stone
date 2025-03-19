package scraper

import (
	"encoding/json"
	"fmt"
)

type AbilityResponse struct {
	Abilities json.RawMessage `json:"abilitydata"`
}

func ScrapeAbilities() error {
	data, err := FetchData(AbilityURL)
	if err != nil {
		return fmt.Errorf("failed to fetch ability data: %v", err)
	}

	if err := SaveToJSON(data, "abilities_minified.json"); err != nil {
		return fmt.Errorf("failed to save ability data: %v", err)
	}

	fmt.Println("Abilities Jinada-ed: Abilities data saved to abilities_minified.json")

	return nil
}
