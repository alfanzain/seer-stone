package scraper

import (
	"encoding/json"
	"fmt"
)

type ItemResponse struct {
	ItemList json.RawMessage `json:"itemdata"`
}

func ScrapeItems() error {
	data, err := FetchData(ItemURL)
	if err != nil {
		return fmt.Errorf("failed to fetch item data: %v", err)
	}

	if err := SaveToJSON(data, "items_minified.json"); err != nil {
		return fmt.Errorf("failed to save item data: %v", err)
	}

	fmt.Println("Items Jinada-ed: Items data saved to items_minified.json")

	return nil
}
