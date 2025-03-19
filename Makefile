.PHONY: all jinada jinada-all jinada-items jinada-heroes jinada-abilities help clean

all: jinada-all

jinada:
	go run main.go jinada $(CMD)

jinada-all:
	go run main.go jinada

jinada-items:
	go run main.go jinada items

jinada-heroes:
	go run main.go jinada heroes

jinada-abilities:
	go run main.go jinada abilities

help:
	go run main.go help

clean:
	rm -rf data/scrapped/*.json
	@echo "All JSON files in data/scrapped/ have been removed."