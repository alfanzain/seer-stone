package scraper

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// FetchData retrieves data from a URL
func FetchData(url string) ([]byte, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-OK HTTP status: %s", resp.Status)
	}

	return io.ReadAll(resp.Body)
}

// SaveToJSON saves data to a JSON file in the data/scrapped directory
func SaveToJSON(data []byte, filename string) error {
	// Ensure data/scrapped directory exists
	scrapDir := filepath.Join("data", "scrapped")
	if err := os.MkdirAll(scrapDir, 0755); err != nil {
		return fmt.Errorf("failed to create %s directory: %v", scrapDir, err)
	}

	// Save to the data/scrapped directory
	filePath := filepath.Join(scrapDir, filename)
	return os.WriteFile(filePath, data, 0644)
}
