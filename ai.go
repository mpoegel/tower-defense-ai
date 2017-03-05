package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mpoegel/tower-defense-ai/strategy"
	"github.com/mpoegel/tower-defense-ai/tdef"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <credential_file> <strategy>\n", args[0])
		os.Exit(1)
	}
	player := tdef.NewPlayer(&args[1])
	var strat tdef.Strategy
	if args[2] == "attack" {
		strat = strategy.NewAttackStrategy(1)
	} else if args[2] == "null" {
		strat = strategy.NewNullStrategy()
	} else {
		fmt.Fprintf(os.Stderr, "Unrecognized strategy: %s", args[2])
		os.Exit(1)
	}
	log.Printf("Loaded credentials for: %s\n", player.Username)
	log.Printf("Using strategy: %s\n", strat.Name())

	tdef.StartGame(player, strat)
}
