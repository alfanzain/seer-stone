package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/alfanzain/seer-stone/scraper"
)

func printUsage() {
	fmt.Println("Usage: go run main.go [command]")
	fmt.Println("Commands:")
	fmt.Println("  jinada        - Scrape all data (items, heroes, abilities)")
	fmt.Println("  jinada items  - Scrape only items data")
	fmt.Println("  jinada heroes - Scrape only heroes data")
	fmt.Println("  jinada abilities - Scrape only abilities data")
	fmt.Println("  help         - Show this help message")
}

func main() {
	// Parse command-line arguments
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := strings.ToLower(os.Args[1])

	// Handle different commands
	switch command {
	case "jinada":
		if len(os.Args) > 2 {
			subCommand := strings.ToLower(os.Args[2])
			switch subCommand {
			case "items":
				fmt.Println("Jinada-ing items data...")
				if err := scraper.ScrapeItems(); err != nil {
					fmt.Printf("Error jinada items: %v\n", err)
					os.Exit(1)
				}
				fmt.Println("Items data jinada-ed successfully!")

			case "heroes":
				fmt.Println("Jinada-ing heroes data...")
				if err := scraper.ScrapeHeroes(); err != nil {
					fmt.Printf("Error jinada heroes: %v\n", err)
					os.Exit(1)
				}
				fmt.Println("Heroes data jinada-ed successfully!")

			case "abilities":
				fmt.Println("Jinada-ing abilities data...")
				if err := scraper.ScrapeAbilities(); err != nil {
					fmt.Printf("Error jinada abilities: %v\n", err)
					os.Exit(1)
				}
				fmt.Println("Abilities data jinada-ed successfully!")

			default:
				fmt.Printf("Unknown subcommand: %s\n", subCommand)
				printUsage()
				os.Exit(1)
			}
		} else {
			// Scrape all data
			fmt.Println("Jinada-ing all data...")

			if err := scraper.ScrapeItems(); err != nil {
				fmt.Printf("Error jinada items: %v\n", err)
			}

			if err := scraper.ScrapeHeroes(); err != nil {
				fmt.Printf("Error jinada heroes: %v\n", err)
			}

			if err := scraper.ScrapeAbilities(); err != nil {
				fmt.Printf("Error jinada abilities: %v\n", err)
			}

			fmt.Println("Jinada-ing completed successfully!")
		}

	case "help":
		printUsage()

	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}
