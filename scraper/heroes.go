package scraper

import (
	"encoding/json"
	"fmt"
)

type HeroResponse struct {
	Heroes json.RawMessage `json:"heroes"`
}

func ScrapeHeroes() error {
	data, err := FetchData(HeroURL)
	if err != nil {
		return fmt.Errorf("failed to fetch hero data: %v", err)
	}

	if err := SaveToJSON(data, "heroes_minified.json"); err != nil {
		return fmt.Errorf("failed to save hero data: %v", err)
	}

	fmt.Println("Heroes Jinada-ed: Heroes data saved to heroes_minified.json")

	return nil
}
