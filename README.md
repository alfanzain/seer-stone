# Seer Stone

A Dota 2 data scraper and database tool written in Go.

## Overview

This project provides tools to scrape and store data from the Dota 2 API, including:
- Items
- Heroes
- Abilities

## Getting Started

### Prerequisites

- Go 1.24.1 or higher

### Installation

```bash
# Clone the repository
git clone https://github.com/alfanzain/seer-stone.git
cd seer-stone

# Install dependencies
go mod tidy
```

## Usage

### Using Go directly

The scraper can be run with different commands:

```bash
# Scrape all data (items, heroes, abilities)
go run main.go jinada

# Scrape only specific data
go run main.go jinada items
go run main.go jinada heroes
go run main.go jinada abilities

# Show help
go run main.go help
```

### Using Makefile

For convenience, a Makefile is provided:

```bash
# Scrape all data
make jinada-all
# or simply
make

# Scrape specific data
make jinada-items
make jinada-heroes
make jinada-abilities

# Show help
make help

# Clean all scraped JSON files
make clean
```

This will generate JSON files in the `data/scrapped` directory:
- `items_minified.json`: Minified version of items data
- `heroes_minified.json`: Minified version of heroes data
- `abilities_minified.json`: Minified version of abilities data

## Project Structure

```
.
├── main.go         # Entry point with command-line interface
├── scraper/        # Scraper package
│   ├── abilities.go # Abilities scraping functionality
│   ├── heroes.go   # Heroes scraping functionality
│   ├── items.go    # Items scraping functionality
│   └── utils.go    # Shared utilities
├── data/           # Data directory
│   └── scrapped/   # Directory for scraped data
├── Makefile        # Convenience commands
└── README.md       # This file
```

## Data Sources

All data is sourced from the official Dota 2 API endpoints:
- Items: https://www.dota2.com/datafeed/itemlist?language=English
- Heroes: https://www.dota2.com/datafeed/herolist?language=English
- Abilities: https://www.dota2.com/datafeed/abilitylist?language=English

## License

MIT 