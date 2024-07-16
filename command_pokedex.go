package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		return errors.New("go catch some pokemon first")
	}

	fmt.Println("Caught Pokemon:")
	for name := range cfg.caughtPokemon {
		fmt.Printf("  -%v\n", name)
	}

	return nil
}
