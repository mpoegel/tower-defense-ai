package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/mpoegel/tower-defense-ai/strategy"
	"github.com/mpoegel/tower-defense-ai/tdef"
)

func parseStrategy(input *string) (*tdef.Strategy, error) {
	var strat tdef.Strategy
	if *input == "attack" {
		strat = strategy.NewAttackStrategy(1)
	} else if *input == "null" {
		strat = strategy.NewNullStrategy()
	} else if *input == "random" {
		strat = strategy.NewRandomStrategy(time.Now().Unix())
	} else {
		err := errors.New("Unrecognized strategy: " + *input)
		return nil, err
	}
	return &strat, nil
}

func main() {
	args := os.Args
	if len(args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <credential_file> <strategy> [<opp_credential_file>"+
			"<opp_strategy>]\n", args[0])
		os.Exit(1)
	}
	if len(args) >= 5 {
		oppPlayer := tdef.NewPlayer(&args[3])
		oppStrat, err := parseStrategy(&args[4])
		if err != nil {
			log.Fatalln(err)
			os.Exit(1)
		}
		go func() {
			log.Printf("[Opponent] Loaded credentials for: %s\n", oppPlayer.Username)
			log.Printf("[Opponent] Using strategy: %s\n", (*oppStrat).Name())
			tdef.StartGame(oppPlayer, oppStrat)
		}()
	}
	player := tdef.NewPlayer(&args[1])
	strat, err := parseStrategy(&args[2])
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	log.Printf("Loaded credentials for: %s\n", player.Username)
	log.Printf("Using strategy: %s\n", (*strat).Name())

	tdef.StartGame(player, strat)
}
